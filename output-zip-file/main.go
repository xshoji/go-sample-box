package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"encoding/binary"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	UsageRequiredPrefix   = "\x1b[33m(REQ)\x1b[0m "
	TimeFormat            = "2006-01-02 15:04:05.0000 [MST]"
	ContentsFileNameInZip = "dummy.tsv"
)

var (
	//go:embed main.go
	srcBytes []byte

	// Command options
	commandDescription      = "Output zip file tool."
	commandOptionFieldWidth = 12
	optionOutputPath        = flag.String("o" /*  */, "" /*    */, UsageRequiredPrefix+"Output path of dummy zip file")
	optionByteSize          = flag.Int("b" /*     */, 1024 /*  */, "Byte size for zip file")
	optionHelp              = flag.Bool("h" /*    */, false /* */, "Help")
)

func init() {
	formatUsage(commandDescription, commandOptionFieldWidth)
}

// # Build: APP="/tmp/tool"; MAIN="main.go"; GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o "${APP}" "${MAIN}"; chmod +x "${APP}"
func main() {

	flag.Parse()
	if *optionHelp || *optionOutputPath == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-30s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
	fmt.Printf("\n\n")

	zipFileDummy := CreateZipFileOnMemory(*optionByteSize)
	fmt.Printf("CreateZipFileOnMemory:\n")
	fmt.Printf("%v\n", zipFileDummy.Bytes())

	path := *optionOutputPath + "_CreateZipFileConcrete.zip"
	CreateZipFileConcrete(path, *optionByteSize)
	fmt.Printf("\nCreateZipFileConcrete:\n")
	fmt.Printf("Path: %s\n", path)

	path = *optionOutputPath + "_CreateZipFileConcreteStatic.zip"
	CreateZipFileConcreteStatic(path, *optionByteSize)
	fmt.Printf("\nCreateZipFileConcreteStatic:\n")
	fmt.Printf("Path: %s\n", path)
}

type ConstantDataUnbufferedReader struct {
	chunkSize          int
	repetitionsCurrent int
	repetitionsMax     int
	remainingByteSize  int
}

func NewConstantDataUnbufferedReader(byteSize int) *ConstantDataUnbufferedReader {
	chunkByteSize := 1024
	return &ConstantDataUnbufferedReader{
		chunkSize:          chunkByteSize,
		repetitionsCurrent: 0,
		repetitionsMax:     byteSize/chunkByteSize + 1,
		remainingByteSize:  byteSize % chunkByteSize,
	}
}

func (r *ConstantDataUnbufferedReader) Read(p []byte) (n int, err error) {
	if r.repetitionsCurrent >= r.repetitionsMax {
		return 0, io.EOF
	}
	chunkSize := r.chunkSize
	if r.repetitionsCurrent == r.repetitionsMax-1 {
		chunkSize = r.remainingByteSize
	}
	if chunkSize != 0 {
		copy(p, bytes.Repeat([]byte("0"), chunkSize))
	}
	r.repetitionsCurrent++
	return chunkSize, nil
}
func CreateZipFileOnMemory(optionByteSize int) *bytes.Buffer {
	var dummyFileBytes []byte
	dummyZipFile := bytes.NewBuffer(dummyFileBytes)
	createZipFile(dummyZipFile, optionByteSize)
	return dummyZipFile

}

func CreateZipFileConcrete(optionOutputPath string, optionByteSize int) {
	dummyZipFile, err := os.Create(optionOutputPath)
	handleError(err, `os.Create(*optionOutputPath)`)
	defer dummyZipFile.Close()
	createZipFile(dummyZipFile, optionByteSize)
}

func CreateZipFileConcreteStatic(optionOutputPath string, optionByteSize int) {
	// 固定値の設定
	filename := "2_" + ContentsFileNameInZip
	content := make([]byte, optionByteSize) // 0埋めされた1024バイト
	crc := crc32.ChecksumIEEE(content)      // 正しいCRC32を計算
	fileSize := uint32(len(content))

	// ファイル作成
	f, err := os.Create(optionOutputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// **ローカルファイルヘッダの書き込み**
	localHeaderOffset, _ := f.Seek(0, io.SeekCurrent)           // ファイルの最初から開始
	f.Write([]byte("PK\x03\x04"))                               // シグネチャ
	f.Write([]byte{0x14, 0x00})                                 // バージョン
	f.Write([]byte{0x00, 0x00})                                 // フラグ
	f.Write([]byte{0x00, 0x00})                                 // 圧縮方式 (0=ストア)
	f.Write([]byte{0x00, 0x00, 0x00, 0x00})                     // タイムスタンプ (0に固定)
	binary.Write(f, binary.LittleEndian, crc)                   // CRC32
	binary.Write(f, binary.LittleEndian, fileSize)              // 圧縮サイズ
	binary.Write(f, binary.LittleEndian, fileSize)              // 元サイズ
	binary.Write(f, binary.LittleEndian, uint16(len(filename))) // ファイル名長
	binary.Write(f, binary.LittleEndian, uint16(0))             // 拡張フィールド長
	f.Write([]byte(filename))                                   // ファイル名
	f.Write(content)                                            // ファイルデータ

	// **セントラルディレクトリの書き込み**
	centralDirectoryOffset, _ := f.Seek(0, io.SeekCurrent)      // 現在の位置
	f.Write([]byte("PK\x01\x02"))                               // シグネチャ
	f.Write([]byte{0x14, 0x00})                                 // バージョン
	f.Write([]byte{0x14, 0x00})                                 // 必要バージョン
	f.Write([]byte{0x00, 0x00})                                 // フラグ
	f.Write([]byte{0x00, 0x00})                                 // 圧縮方式
	f.Write([]byte{0x00, 0x00, 0x00, 0x00})                     // タイムスタンプ
	binary.Write(f, binary.LittleEndian, crc)                   // CRC32
	binary.Write(f, binary.LittleEndian, fileSize)              // 圧縮サイズ
	binary.Write(f, binary.LittleEndian, fileSize)              // 元サイズ
	binary.Write(f, binary.LittleEndian, uint16(len(filename))) // ファイル名長
	binary.Write(f, binary.LittleEndian, uint16(0))             // 拡張フィールド長
	binary.Write(f, binary.LittleEndian, uint16(0))             // ファイルコメント長
	binary.Write(f, binary.LittleEndian, uint16(0))             // ディスク番号
	binary.Write(f, binary.LittleEndian, uint16(0))             // 内部ファイル属性
	binary.Write(f, binary.LittleEndian, uint32(0))             // 外部ファイル属性
	binary.Write(f, binary.LittleEndian, localHeaderOffset)     // ローカルヘッダのオフセット
	f.Write([]byte(filename))                                   // ファイル名

	// **終端ヘッダ (EOCD) の書き込み**
	f.Write([]byte("PK\x05\x06"))                                        // シグネチャ
	f.Write([]byte{0x00, 0x00})                                          // ディスク番号
	f.Write([]byte{0x00, 0x00})                                          // セントラルディレクトリ開始ディスク
	f.Write([]byte{0x01, 0x00})                                          // セントラルディレクトリのエントリ数 (1)
	f.Write([]byte{0x01, 0x00})                                          // セントラルディレクトリの総エントリ数 (1)
	centralDirectorySize, _ := f.Seek(0, io.SeekCurrent)                 // aaa
	centralDirectorySize -= centralDirectoryOffset                       //
	binary.Write(f, binary.LittleEndian, uint32(centralDirectorySize))   // セントラルディレクトリサイズ
	binary.Write(f, binary.LittleEndian, uint32(centralDirectoryOffset)) // セントラルディレクトリのオフセット
	binary.Write(f, binary.LittleEndian, uint16(0))                      // コメント長

	f.Sync()

}

func createZipFile(dummyZipFile io.Writer, optionByteSize int) {
	zipWriter := zip.NewWriter(dummyZipFile)
	filename := "1_" + ContentsFileNameInZip
	w1, err := zipWriter.Create(filename)
	handleError(err, `zipWriter.Create(ContentsFileNameInZip)`)
	constantDataUnbufferedReader := NewConstantDataUnbufferedReader(optionByteSize)
	_, err = io.Copy(w1, constantDataUnbufferedReader)
	handleError(err, `io.Copy(w1, constantDataUnbufferedReader)`)
	err = zipWriter.Close()
	handleError(err, `zipWriter.Close()`)
}

func handleError(err error, prefixErrMessage string) {
	if err != nil {
		fmt.Printf("%s [ERROR %s]: %v\n", time.Now().Format(TimeFormat), prefixErrMessage, err)
	}
}

// formatUsage optionFieldWidth [ recommended width = general: 12, bool only: 5 ]
func formatUsage(description string, optionFieldWidth int) {
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	usageLines := strings.Split(b.String(), "\n")
	usage := strings.Replace(strings.Replace(usageLines[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + description + "\n\nOptions:\n"
	re := regexp.MustCompile(` +(-\S+)(?: (\S+))?\n*(\s+)(.*)\n`)
	usage += re.ReplaceAllStringFunc(strings.Join(usageLines[1:], "\n"), func(m string) string {
		parts := re.FindStringSubmatch(m)
		return fmt.Sprintf("  %-"+strconv.Itoa(optionFieldWidth)+"s %s\n", parts[1]+" "+strings.TrimSpace(parts[2]), parts[4])
	})
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}

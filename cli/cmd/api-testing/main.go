package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"maps"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	HttpContentTypeHeader = "content-type"
	TimeFormat            = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	// Command options ( the -h, --help option is defined by default in the flag package )
	commandDescription               = "Web API testing tool."
	commandOptionFieldWidth          = 7
	optionByteSizePostData           = flag.Int("b" /*  */, 1024 /*  */, "Byte size for post data")
	optionUseChunkedTransferEncoding = flag.Bool("c" /* */, false /* */, "Use \"Transfer-Encoding: chunked\" ( only for HTTP/1.1 ).")
	optionTrimDownHttpMessages       = flag.Bool("t" /* */, false /* */, "Trim down HTTP messages in stdout.")
	optionSkipTlsVerification        = flag.Bool("s" /* */, false /* */, "Skip TLS verification.")
	optionDisableHttp2               = flag.Bool("d" /* */, false /* */, "Disable HTTP/2.")

	// HTTP Header templates
	createHttpHeaderEmpty = func() map[string]string {
		return maps.Clone(make(map[string]string))
	}
	createHttpHeaderContentTypeForm = func() map[string]string {
		return maps.Clone(map[string]string{HttpContentTypeHeader: "application/x-www-form-urlencoded;charset=utf-8"})
	}
	createHttpHeaderContentTypeJson = func() map[string]string {
		return maps.Clone(map[string]string{HttpContentTypeHeader: "application/json"})
	}
	createHttpHeaderContentTypeOctetStream = func() map[string]string {
		return maps.Clone(map[string]string{HttpContentTypeHeader: "application/octet-stream"})
	}

	// Masking console log
	muskingRegex = regexp.MustCompile(`(Accept-Encoding:|Etag:|"key1":)(.*)`)
)

func init() {
	formatUsage(commandDescription, commandOptionFieldWidth)
}

// # Build: GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /tmp/main main.go
func main() {

	flag.Parse()
	client := http.Client{
		Transport: CreateCustomTransport(
			&tls.Config{InsecureSkipVerify: *optionSkipTlsVerification},
			*optionDisableHttp2,
			"tcp4",
		),
	}

	// Print all options
	fmt.Printf("[ Command options ]\n")
	flag.VisitAll(func(a *flag.Flag) {
		fmt.Printf("  -%-10s %s\n", fmt.Sprintf("%s %v", a.Name, a.Value), strings.Trim(a.Usage, "\n"))
	})
	fmt.Printf("\n\n")

	headers := createHttpHeaderContentTypeJson()

	//
	//
	//
	//
	//
	//
	//
	//
	targetUrl := "https://httpbin.org/post"
	body := `{
  "title": "title1",
  "obj": {
    "list": [
      {
        "key1": "value"
      }
    ]
  }
}
`
	response := HttpRequest(client, "POST", targetUrl, headers, strings.NewReader(body))
	fmt.Println()
	fmt.Printf("\"obj.list.0.key1\": %s\n", response.(map[string]any)["json"].(map[string]any)["obj"].(map[string]any)["list"].([]any)[0].(map[string]any)["key1"].(string))
	fmt.Printf("\"obj.list.0.key1\": %s\n\n\n\n", Get(response, "json.obj.list.0.key1").(string))

	//
	//
	//
	//
	//
	//
	//
	//
	csvFileContents := `title,length,memo
movie1,120,Horror movie
movie2,240,Action movie
movie3,5,Short movie
`
	f, err := os.CreateTemp("", "*_api_testing.csv")
	handleError(err, `os.CreateTemp()`)

	// Successful handling
	defer os.Remove(f.Name())
	// Signal handling
	go func() {
		// 1. Create channel for os.Signal and Notify interrupt signal to channel.
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM) // syscall.SIGINT, syscall.SIGTERM ( [!] Cannot handle os.Kill )

		// 2. Wait signal
		s := <-signalChannel
		fmt.Printf("Signal is received: %v\n", s)
		os.Remove(f.Name())
		signalValue := int(s.(syscall.Signal))
		os.Exit(signalValue)
	}()

	err = os.WriteFile(f.Name(), []byte(csvFileContents), 0655)
	handleError(err, "os.WriteFile(f.Name(), []byte(csvFileContents), 0655)")
	fmt.Printf("temp file: %s\n\n\n", f.Name())

	response = HttpRequestMultipartFormData(client, "POST", targetUrl, headers, map[string]map[string]io.Reader{
		"request": {
			"application/json": strings.NewReader(`{"title":"movie_title"}`),
		},
		"file": {
			//"text/csv; charset=utf-8": func() (f *os.File) { f, _ = os.Open("/tmp/aaa.csv"); return }(),
			"text/csv;charset=utf-8": f,
		},
	})
	fmt.Println(response)
	fmt.Printf("\n\n\n\n")

	//
	//
	//
	//
	//
	//
	//
	//
	csvFileContents2 := `aaa,bbb,ccc
aaa1,bbb2,ccc3
aaa4,bbb5,ccc6
aaa7,bbb8,ccc9
`

	response = HttpRequestMultipartFormData(client, "POST", targetUrl, headers, map[string]map[string]io.Reader{
		"request": {
			"application/json": strings.NewReader(`{"title":"movie_title"}`),
		},
		"file": {
			"text/csv;charset=utf-8": bytes.NewReader([]byte(csvFileContents2)),
		},
	})
	fmt.Println(response)
	fmt.Printf("\n\n\n\n")

	//
	//
	//
	//
	//
	//
	//
	//
	response = HttpRequestMultipartFormData(client, "POST", targetUrl, headers, map[string]map[string]io.Reader{
		"request": {
			"application/json": strings.NewReader(`{"title":"movie_title"}`),
		},
		"file": {
			"text/csv;charset=utf-8": NewConstantDataUnbufferedReader(*optionByteSizePostData),
		},
	})
	fmt.Println(response)

	//
	//
	//
	//
	//
	//
	//
	//
	headers = createHttpHeaderContentTypeOctetStream()
	csvFileContents3 := `ddd,eee,fff
aaa8,bbb11,ccc14
aaa9,bbb12,ccc15
aaa10,bbb13,ccc16
`
	targetUrl = "https://httpbin.org/put"
	response = HttpRequest(client, "PUT", targetUrl, headers, bytes.NewBufferString(csvFileContents3))
	fmt.Println(response)
}

// =======================================
// io.Reader implementation
// =======================================

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

// =======================================
// HTTP Utils
// =======================================

// CustomTransport Debugging HTTP Client requests with Go: https://www.jvt.me/posts/2023/03/11/go-debug-http/
type CustomTransport struct {
	// Embed default transport
	*http.Transport
}

func (s *CustomTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	httpMessageBytes, err := httputil.DumpRequestOut(r, true)
	handleError(err, "httputil.DumpRequestOut(r, true)")

	adjustMessage := func(message string) string {
		if *optionTrimDownHttpMessages {
			message = strings.Replace(message, "\r\n", ", ", -1)
			message = strings.Replace(message, "\n", " ", -1)
		}
		return message
	}

	// Print remote IP
	r = r.WithContext(httptrace.WithClientTrace(r.Context(), &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("%s", adjustMessage(fmt.Sprintf("[RemoteAddr=%s]\n", connInfo.Conn.RemoteAddr())))
		},
	}))

	// mask request header and body
	httpMessageString := muskingRegex.ReplaceAllString(string(httpMessageBytes), "$1 ****")
	fmt.Printf("Req. %s%s", time.Now().Format(TimeFormat), adjustMessage("\n"+httpMessageString+"\n"))

	resp, err := s.Transport.RoundTrip(r)
	handleError(err, "s.Transport.RoundTrip(r)")
	// Goのnet/httpのkeep-aliveで気をつけること - Carpe Diem: https://christina04.hatenablog.com/entry/go-keep-alive
	respBytes, err := httputil.DumpResponse(resp, false)
	handleError(err, "httputil.DumpResponse(resp, true)")

	// mask response header
	respString := muskingRegex.ReplaceAllString(string(respBytes), "$1 ****")
	fmt.Printf("Res. %s%s", time.Now().Format(TimeFormat), adjustMessage("\n"+respString))

	return resp, err
}

// CreateCustomTransport
// [golang custom http client] #go #golang #http #client #timeouts #dns #resolver
// https://gist.github.com/Integralist/8a9cb8924f75ae42487fd877b03360e2?permalink_comment_id=4863513
func CreateCustomTransport(tlsConfig *tls.Config, disableHttp2 bool, networkType string) *CustomTransport {
	customTr := &CustomTransport{Transport: http.DefaultTransport.(*http.Transport).Clone()}
	if tlsConfig != nil {
		customTr.TLSClientConfig = tlsConfig
	}
	if disableHttp2 {
		// hdr-HTTP_2 - http package - net/http - Go Packages: https://pkg.go.dev/net/http#hdr-HTTP_2
		// disable HTTP/2 can do so by setting [Transport.TLSNextProto] (for clients) or [Server.TLSNextProto] (for servers) to a non-nil, empty map.
		customTr.TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
	}
	// Go http get force to use ipv4 - Stack Overflow : https://stackoverflow.com/questions/77718022/go-http-get-force-to-use-ipv4
	customTr.DialContext = func(ctx context.Context, network string, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, networkType, addr)
	}
	return customTr
}

func HttpRequestMultipartFormData(client http.Client, method string, url string, headers map[string]string, multipartValues map[string]map[string]io.Reader) any {
	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)
	for fieldName, contentTypeAndIoReader := range multipartValues {

		// Create MIME encoded form files that auto-detect the content type. by danny-cheung · Pull Request #170 · go-openapi/runtime
		// https://github.com/go-openapi/runtime/pull/170/files

		// Create the MIME headers for the new part
		var contentType string
		var ioReader io.Reader
		var fw io.Writer
		var err error
		for c, i := range contentTypeAndIoReader {
			contentType = c
			ioReader = i
		}
		h := make(textproto.MIMEHeader)
		h.Set("Content-Type", contentType)
		if file, ok := ioReader.(*os.File); ok {
			// Create the MIME headers for the new part
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, filepath.Base(file.Name())))
		} else if _, ok := ioReader.(*ConstantDataUnbufferedReader); ok {
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, "dummyFileName"))
		} else if _, ok := ioReader.(*bytes.Reader); ok {
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, "dummyFileName"))
		} else {
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`, fieldName))
		}
		fw, err = multipartWriter.CreatePart(h)
		handleError(err, "multipartWriter.CreatePart(h)")
		_, err = io.Copy(fw, ioReader)
		handleError(err, "io.Copy(fw, fileContent)")

	}
	multipartWriter.Close()
	headers["content-type"] = multipartWriter.FormDataContentType()
	return HttpRequest(client, method, url, headers, body)
}

func HttpRequestFormUrlencoded(client http.Client, method string, url string, headers map[string]string, values url.Values) any {
	return HttpRequest(client, method, url, headers, strings.NewReader(values.Encode()))
}

func HttpRequest(client http.Client, method string, url string, headers map[string]string, body io.Reader) any {
	req, err := http.NewRequest(method, url, body)
	handleError(err, "http.NewRequest(method, url, body)")
	if *optionUseChunkedTransferEncoding {
		req.TransferEncoding = []string{"chunked"}
		req.ContentLength = -1
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)
	handleError(err, "client.Do(req)")
	responseBody, err := io.ReadAll(res.Body)
	handleError(err, "io.ReadAll(res.Body)")
	responseBodyJsonObject := ToJsonObject(responseBody)

	var jsonString string
	var jsonBytes []byte

	if *optionTrimDownHttpMessages {
		jsonBytes, err = json.Marshal(responseBodyJsonObject)
		handleError(err, `json.Marshal(responseBodyJsonObject)`)
		jsonString = strings.Replace(string(jsonBytes), "\r\n", ", ", -1)
		jsonString = strings.Replace(jsonString, "\n", " ", -1)
	} else {
		jsonBytes, err = json.MarshalIndent(responseBodyJsonObject, "", "  ")
		handleError(err, `json.MarshalIndent(responseBodyJsonObject, "", "  ")`)
		jsonString = string(jsonBytes)
	}

	// mask response body
	prettyJsonString := muskingRegex.ReplaceAllString(jsonString, "$1 ****")
	fmt.Printf("%s\n", prettyJsonString)
	return responseBodyJsonObject
}

// =======================================
// Json Utils
// =======================================

// ToJsonObject json bytes to any object
func ToJsonObject(body []byte) any {
	var jsonObject any
	err := json.Unmarshal(body, &jsonObject)
	handleError(err, "json.Unmarshal")
	return jsonObject
}

// Get get value in any object [ example : object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb" ]
func Get(object any, keyChain string) any {
	var result any
	var exists bool
	for _, key := range strings.Split(keyChain, ".") {
		exists = false
		if _, ok := object.(map[string]any); ok {
			exists = true
			object = object.(map[string]any)[key]
			result = object
			continue
		}
		if values, ok := object.([]any); ok {
			for i, v := range values {
				if strconv.FormatInt(int64(i), 10) == key {
					exists = true
					object = v
					result = object
					continue
				}
			}
		}
	}
	if exists {
		return result
	}
	return nil
}

// ToMap to map
func ToMap(v any, keys []string) map[string]any {
	resultMap := make(map[string]any, len(keys))
	for _, key := range keys {
		resultMap[key] = Get(v, key)
	}
	return resultMap
}

// ToJsonString to json string
func ToJsonString(v any) string {
	result, _ := json.Marshal(v)
	return string(result)
}

// =======================================
// Common Utils
// =======================================

// Create uuid
func createUuid() string {
	seed := strconv.FormatInt(time.Now().UnixNano(), 10)
	shaBytes := sha256.Sum256([]byte(seed))
	return hex.EncodeToString(shaBytes[:16])
}

// Handle error
func handleError(err error, prefixErrMessage string) {
	if err != nil {
		fmt.Printf("%s [ERROR %s]: %v\n", time.Now().Format(TimeFormat), prefixErrMessage, err)
	}
}

// GetEnvOrDefault environment value ( with default value )
func GetEnvOrDefault(key string, defaultValue string) string {
	value := defaultValue
	v := os.Getenv(key)
	if v != "" {
		value = v
	}
	return value
}

// formatUsage optionFieldWidth [ general: 12, bool only: 5 ]
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
	flag.Usage = func() { _, _ = fmt.Fprint(flag.CommandLine.Output(), usage) }
}

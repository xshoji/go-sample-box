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
	"math"
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
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	UsageRequiredPrefix   = "\u001B[33m[required]\u001B[0m "
	UsageDummy            = "########"
	CommandDescription    = "Web API testing tool."
	HttpContentTypeHeader = "content-type"
	TimeFormat            = "2006-01-02 15:04:05.9999 [MST]"
)

var (
	// Define options
	optionEnableCompressHttpMessage = defineBoolFlag("c", "compress-http-message", "compress http message")
	optionSkipTlsVerification       = defineBoolFlag("s", "skip-tls-verification", "skip tls verification")
	optionDisableHttp2              = defineBoolFlag("d", "disable-http2", "disable HTTP/2")
	optionHelp                      = defineBoolFlag("h", "help", "show help")

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

	// Masking console log
	muskingRegex = regexp.MustCompile(`(Accept-Encoding:|Etag:|"key1":)(.*)`)
)

func init() {
	formatUsage()
}

func main() {

	flag.Parse()
	if *optionHelp {
		flag.Usage()
		os.Exit(0)
	}

	client := http.Client{
		Transport: CreateCustomTransport(
			&tls.Config{InsecureSkipVerify: *optionSkipTlsVerification},
			*optionDisableHttp2,
			"tcp4",
		),
	}

	fmt.Println("#--------------------")
	fmt.Println("# Command options")
	fmt.Println("#--------------------")
	fmt.Printf("compress http message : %t\n", *optionEnableCompressHttpMessage)
	fmt.Printf("skip tls Verification : %t\n", *optionSkipTlsVerification)
	fmt.Printf("disable HTTP/2        : %t\n\n\n", *optionDisableHttp2)

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
	response := DoHttpRequest(client, "POST", targetUrl, headers, body)
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

	response = DoHttpRequestMultipartFormData(client, "POST", targetUrl, headers, map[string]map[string]io.Reader{
		"request": {
			"application/json": bytes.NewReader([]byte(`{"title":"movie_title"}`)),
		},
		"file": {
			//"text/csv; charset=utf-8": func() (f *os.File) { f, _ = os.Open("/tmp/aaa.csv"); return }(),
			"text/csv;charset=utf-8": f,
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
	csvFileContents2 := `aaa,bbb,ccc
aaa1,bbb2,ccc3
aaa4,bbb5,ccc6
aaa7,bbb8,ccc9
`

	response = DoHttpRequestMultipartFormData(client, "POST", targetUrl, headers, map[string]map[string]io.Reader{
		"request": {
			"application/json": bytes.NewReader([]byte(`{"title":"movie_title"}`)),
		},
		"file": {
			"text/csv;charset=utf-8": &AnyContentNoBufferedReader{
				content:        []byte(csvFileContents2),
				byteArrayIndex: 0,
			},
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
	response = DoHttpRequestMultipartFormData(client, "POST", targetUrl, headers, map[string]map[string]io.Reader{
		"request": {
			"application/json": bytes.NewReader([]byte(`{"title":"movie_title"}`)),
		},
		"file": {
			"text/csv;charset=utf-8": &ConstantContentNoBufferedReader{
				kbSize:      5,
				repetitions: 0,
			},
		},
	})
	fmt.Println(response)
}

// =======================================
// io.Reader implementation
// =======================================

type ConstantContentNoBufferedReader struct {
	kbSize      int
	repetitions int
}

func (r *ConstantContentNoBufferedReader) Read(p []byte) (n int, err error) {
	chunkSize := 1024 // 1kb
	if r.repetitions >= r.kbSize {
		return 0, io.EOF
	}
	copy(p, bytes.Repeat([]byte("0"), chunkSize))
	r.repetitions++
	return chunkSize, nil
}

type AnyContentNoBufferedReader struct {
	content        []byte
	byteArrayIndex int
}

func (r *AnyContentNoBufferedReader) Read(p []byte) (n int, err error) {
	chunkSize := 1024 // 1kb
	if r.byteArrayIndex == len(r.content) {
		return 0, io.EOF
	}
	if diff := r.byteArrayIndex + chunkSize - len(r.content); diff > 0 {
		chunkSize = chunkSize - diff
	}
	copy(p, r.content[r.byteArrayIndex:chunkSize])
	r.byteArrayIndex = r.byteArrayIndex + chunkSize
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
		if *optionEnableCompressHttpMessage {
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
	httpMessageString := string(httpMessageBytes)
	httpMessageString = muskingRegex.ReplaceAllString(httpMessageString, "$1 ****")
	fmt.Printf("Req. %s%s", time.Now().Format(TimeFormat), adjustMessage("\n"+httpMessageString+"\n"))

	resp, err := s.Transport.RoundTrip(r)
	handleError(err, "s.Transport.RoundTrip(r)")
	// Goのnet/httpのkeep-aliveで気をつけること - Carpe Diem: https://christina04.hatenablog.com/entry/go-keep-alive
	respBytes, err := httputil.DumpResponse(resp, false)
	handleError(err, "httputil.DumpResponse(resp, true)")

	// mask response header
	respString := string(respBytes)
	respString = muskingRegex.ReplaceAllString(respString, "$1 ****")
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

func DoHttpRequestMultipartFormData(client http.Client, method string, url string, headers map[string]string, multipartValues map[string]map[string]io.Reader) interface{} {
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
		} else if _, ok := ioReader.(*ConstantContentNoBufferedReader); ok {
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, "dummyFileName"))
		} else if _, ok := ioReader.(*AnyContentNoBufferedReader); ok {
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
	return internalDoHttpRequest(client, method, url, headers, body)
}

func DoHttpRequestFormUrlencoded(client http.Client, method string, url string, headers map[string]string, values url.Values) interface{} {
	return internalDoHttpRequest(client, method, url, headers, strings.NewReader(values.Encode()))
}

func DoHttpRequest(client http.Client, method string, url string, headers map[string]string, body string) interface{} {
	return internalDoHttpRequest(client, method, url, headers, strings.NewReader(body))
}

func internalDoHttpRequest(client http.Client, method string, url string, headers map[string]string, body io.Reader) interface{} {
	req, err := http.NewRequest(method, url, body)
	handleError(err, "http.NewRequest(method, url, body)")
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

	if *optionEnableCompressHttpMessage {
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

// ToJsonObject json bytes to interface{} object
func ToJsonObject(body []byte) interface{} {
	var jsonObject interface{}
	err := json.Unmarshal(body, &jsonObject)
	handleError(err, "json.Unmarshal")
	return jsonObject
}

// Get get value in interface{} object [ example : object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb" ]
func Get(object interface{}, keyChain string) interface{} {
	var result interface{}
	var exists bool
	for _, key := range strings.Split(keyChain, ".") {
		exists = false
		if _, ok := object.(map[string]interface{}); ok {
			exists = true
			object = object.(map[string]interface{})[key]
			result = object
			continue
		}
		if values, ok := object.([]interface{}); ok {
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
func ToMap(v interface{}, keys []string) map[string]interface{} {
	resultMap := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		resultMap[key] = Get(v, key)
	}
	return resultMap
}

// ToJsonString to json string
func ToJsonString(v interface{}) string {
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

func defineBoolFlag(short, long, description string) (v *bool) {
	v = flag.Bool(short, false, UsageDummy)
	flag.BoolVar(v, long, false, description)
	return
}

func formatUsage() {
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + ".*\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageOptions := re.FindAllString(b.String(), -1)
	maxLength := 0.0
	sort.Slice(usageOptions, func(i, j int) bool {
		maxLength = math.Max(maxLength, math.Max(float64(len(re.ReplaceAllString(usageOptions[i], "$1, -$3$4"))), float64(len(re.ReplaceAllString(usageOptions[j], "$1, -$3$4")))))
		if len(strings.Split(usageOptions[i]+usageOptions[j], UsageRequiredPrefix))%2 == 1 {
			return strings.Compare(usageOptions[i], usageOptions[j]) == -1
		} else {
			return strings.Index(usageOptions[i], UsageRequiredPrefix) >= 0
		}
	})
	usage := strings.Replace(strings.Replace(strings.Split(b.String(), "\n")[0], ":", " [OPTIONS]", -1), " of ", ": ", -1) + "\n\nDescription:\n  " + CommandDescription + "\n\nOptions:\n"
	for _, v := range usageOptions {
		usage += fmt.Sprintf("%-6s%-"+strconv.Itoa(int(maxLength))+"s", re.ReplaceAllString(v, "  $1,"), re.ReplaceAllString(v, "-$3$4")) + re.ReplaceAllString(v, "$5\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}

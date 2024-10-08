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
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	UsageRequiredPrefix     = "\u001B[33m[required]\u001B[0m "
	UsageDummy              = "########"
	HttpContentTypeHeader   = "Content-Type"
	ContextKeyPrettyHttpLog = "ContextKeyLoggingPrettyHttpLog"
	TimeFormat              = "2006-01-02 15:04:05.9999 [MST]"
)

var (
	// Define short parameters ( this default value will be not used ).
	paramsPrettyHttpMessage   = flag.Bool("p", true, UsageDummy)
	paramsSkipTlsVerification = flag.Bool("s", false, UsageDummy)
	paramsDisableHttp2        = flag.Bool("d", false, UsageDummy)
	paramsHelp                = flag.Bool("h", false, UsageDummy)

	// HTTP Header templates
	httpHeaderEmptyMap        = make(map[string]string)
	httpHeaderContentTypeForm = map[string]string{HttpContentTypeHeader: "application/x-www-form-urlencoded;charset=utf-8"}
	httpHeaderContentTypeJson = map[string]string{HttpContentTypeHeader: "application/json;charset=utf-8"}

	// Masking console log
	muskingRegex = regexp.MustCompile(`(Accept-Encoding:|Etag:|"key1":)(.*)`)
)

func init() {
	// Define long parameters
	flag.BoolVar(paramsPrettyHttpMessage /*   */, "pretty-http-message" /*    */, true /*   */, "print pretty http message")
	flag.BoolVar(paramsSkipTlsVerification /* */, "skip-tls-verification" /*  */, false /*   */, "skip tls verification")
	flag.BoolVar(paramsDisableHttp2 /*        */, "disable-http2" /*          */, false /*   */, "disable HTTP/2")
	flag.BoolVar(paramsHelp /*                */, "help" /*                   */, false /*   */, "show help")

	adjustUsage()
}

func main() {

	flag.Parse()
	if *paramsHelp {
		flag.Usage()
		os.Exit(0)
	}

	client := http.Client{
		Transport: CreateCustomTransport(
			&tls.Config{
				InsecureSkipVerify: *paramsSkipTlsVerification,
			},
			*paramsDisableHttp2,
			"tcp4",
		),
	}

	fmt.Println("#--------------------")
	fmt.Println("# Command information")
	fmt.Println("#--------------------")
	fmt.Printf("pretty http message   : %t\n", *paramsPrettyHttpMessage)
	fmt.Printf("skip tls Verification : %t\n", *paramsSkipTlsVerification)
	fmt.Printf("disable HTTP/2        : %t\n", *paramsDisableHttp2)

	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKeyPrettyHttpLog, *paramsPrettyHttpMessage)

	headers := maps.Clone(httpHeaderContentTypeJson)

	//
	//
	//
	//
	//
	//
	//
	//
	targetUrl := "https://my-json-server.typicode.com/typicode/demo/posts/"
	body := `{
  "title": "hoge2",
  "obj": {
    "list": [
      {
        "key1": "value"
      }
    ]
  }
}
`

	response := DoHttpRequest(ctx, client, "POST", targetUrl, headers, body)
	fmt.Println("\"obj.list.0.key1\": " + Get(response, "obj.list.0.key1").(string))

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
		if !r.Context().Value(ContextKeyPrettyHttpLog).(bool) {
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
	respBodyString := string(respBytes)
	respBodyString = muskingRegex.ReplaceAllString(respBodyString, "$1 ****")
	fmt.Printf("Res. %s%s\n", time.Now().Format(TimeFormat), adjustMessage("\n"+respBodyString))

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

func DoHttpRequestMultipartFormData(ctx context.Context, client http.Client, method string, url string, headers map[string]string, multipartValues map[string]io.Reader) interface{} {
	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)
	for fieldName, ioReader := range multipartValues {
		var fw io.Writer
		var err error
		if fileContent, ok := ioReader.(*os.File); ok {
			fw, err = multipartWriter.CreateFormFile(fieldName, fileContent.Name())
			handleError(err, "multipartWriter.CreateFormFile(fieldName, fileContent.Name())")
		} else {
			fw, err = multipartWriter.CreateFormField(fieldName)
			handleError(err, "multipartWriter.CreateFormField(fieldName)")
		}
		_, err = io.Copy(fw, body)

	}
	defer multipartWriter.Close()
	return internalDoHttpRequest(ctx, client, method, url, headers, body)
}

func DoHttpRequestFormUrlencoded(ctx context.Context, client http.Client, method string, url string, headers map[string]string, values url.Values) interface{} {
	return internalDoHttpRequest(ctx, client, method, url, headers, strings.NewReader(values.Encode()))
}

func DoHttpRequest(ctx context.Context, client http.Client, method string, url string, headers map[string]string, body string) interface{} {
	return internalDoHttpRequest(ctx, client, method, url, headers, strings.NewReader(body))
}

func internalDoHttpRequest(ctx context.Context, client http.Client, method string, url string, headers map[string]string, body io.Reader) interface{} {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	handleError(err, "http.NewRequestWithContext")
	res, err := client.Do(req)
	handleError(err, "client.Do(req)")
	responseBody, err := io.ReadAll(res.Body)
	handleError(err, "io.ReadAll(res.Body)")
	responseBodyJsonObject := ToJsonObject(responseBody)
	prettyJsonBytes, err := json.MarshalIndent(responseBodyJsonObject, "", "  ")
	handleError(err, "json.MarshalIndent(ToJsonObject(responseBody), \"\", \"  \")")
	// mask response body
	prettyJsonString := muskingRegex.ReplaceAllString(string(prettyJsonBytes), "$1 ****")
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

// Get environment value ( with default value )
func getEnv(key string, defaultValue string) string {
	value := defaultValue
	v := os.Getenv(key)
	if v != "" {
		value = v
	}
	return value
}

func adjustUsage() {
	// Get default flags usage
	b := new(bytes.Buffer)
	func() { flag.CommandLine.SetOutput(b); flag.Usage(); flag.CommandLine.SetOutput(os.Stderr) }()
	// Sort params and description ( order by UsageRequiredPrefix )
	re := regexp.MustCompile("(-\\S+)( *\\S*)+\n*\\s+" + UsageDummy + "\n*\\s+(-\\S+)( *\\S*)+\n\\s+(.+)")
	usageParams := re.FindAllString(b.String(), -1)
	maxLengthParam := 0.0
	sort.Slice(usageParams, func(i, j int) bool {
		maxLengthParam = math.Max(maxLengthParam, float64(len(re.ReplaceAllString(usageParams[i], "$1, -$3$4"))))
		maxLengthParam = math.Max(maxLengthParam, float64(len(re.ReplaceAllString(usageParams[j], "$1, -$3$4"))))
		isRequired1 := strings.Index(usageParams[i], UsageRequiredPrefix) >= 0
		isRequired2 := strings.Index(usageParams[j], UsageRequiredPrefix) >= 0
		if isRequired1 == isRequired2 {
			return strings.Compare(usageParams[i], usageParams[j]) == -1
		} else {
			return isRequired1
		}
	})
	// Adjust usage
	usage := strings.Split(b.String(), "\n")[0] + "\n\n"
	usage = usage + "Description:\n  HTTP request/response testing tool.\n\n"
	usage = usage + "Options:\n"
	for _, v := range usageParams {
		usage = usage + fmt.Sprintf("  %-3s"+"%-"+strconv.Itoa(int(maxLengthParam))+"s", re.ReplaceAllString(v, "$1"), re.ReplaceAllString(v, ", -$3$4")) + re.ReplaceAllString(v, "$5\n")
	}
	flag.Usage = func() { _, _ = fmt.Fprintf(flag.CommandLine.Output(), usage) }
}

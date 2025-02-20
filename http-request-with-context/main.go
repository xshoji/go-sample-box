package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	ContextKeyCompressHttpLog = "ContextKeyLoggingCompressHttpLog"
	TimeFormat                = "2006-01-02 15:04:05.0000 [MST]"
)

var (
	// HTTP Header templates
	createHttpHeaderContentTypeJson = func() map[string]string {
		return maps.Clone(map[string]string{"content-type": "application/json"})
	}

	// Masking console log
	muskingRegex = regexp.MustCompile(`(Accept-Encoding:|Etag:|"key1":)(.*)`)
)

func main() {

	client := http.Client{
		Transport: CreateCustomTransport(
			&tls.Config{InsecureSkipVerify: false},
			false,
			"tcp4",
		),
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKeyCompressHttpLog, false)

	headers := createHttpHeaderContentTypeJson()

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
	response := DoHttpRequest(ctx, client, "POST", targetUrl, headers, body)
	fmt.Printf("\"obj.list.0.key1\": %s\n\n\n\n", Get(response, "json.obj.list.0.key1").(string))

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
		if r.Context().Value(ContextKeyCompressHttpLog).(bool) {
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

func DoHttpRequest(ctx context.Context, client http.Client, method string, url string, headers map[string]string, body string) interface{} {
	return internalDoHttpRequest(ctx, client, method, url, headers, strings.NewReader(body))
}

func internalDoHttpRequest(ctx context.Context, client http.Client, method string, url string, headers map[string]string, body io.Reader) interface{} {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	handleError(err, "http.NewRequestWithContext")
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
	if ctx.Value(ContextKeyCompressHttpLog).(bool) {
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

// =======================================
// Common Utils
// =======================================

// Handle error
func handleError(err error, prefixErrMessage string) {
	if err != nil {
		fmt.Printf("%s [ERROR %s]: %v\n", time.Now().Format(TimeFormat), prefixErrMessage, err)
	}
}

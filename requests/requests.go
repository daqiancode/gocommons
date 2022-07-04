package requests

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"

	"github.com/daqiancode/jsoniter"
	"github.com/ddliu/go-httpclient"
)

type HttpRequest struct {
	*httpclient.HttpClient
	json jsoniter.API
}

func NewHttpRequest(headers, cookies map[string]string, options map[int]interface{}) *HttpRequest {
	client := httpclient.NewHttpClient()
	if headers != nil {
		client = client.WithHeaders(headers)
	}
	if cookies != nil {
		httpCookies := make([]*http.Cookie, len(cookies))
		i := 0
		for k, v := range cookies {
			httpCookies[i] = &http.Cookie{Name: k, Value: v}
			i++
		}
		client = client.WithCookie(httpCookies...)
	}
	if options != nil {
		for k, v := range options {
			client.WithOption(k, v)
		}
	}
	return &HttpRequest{
		HttpClient: client,
		json:       jsoniter.Decapitalized,
	}
}

func (s *HttpRequest) Call(method string, url string, value interface{}, headers map[string]string) (*httpclient.Response, error) {
	bs, err := s.json.Marshal(value)
	if err != nil {
		return nil, err
	}
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Content-Type"] = "application/json"

	return s.Do(method, url, headers, bytes.NewBuffer(bs))
}

type PostFile struct {
	Field    string
	FileName string
	Content  io.Reader
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

//PostFile PostFile.Content will not close after PostFile
func (s *HttpRequest) PostFile(url string, form map[string]string, files []PostFile) (*httpclient.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range form {
		writer.WriteField(k, v)
	}

	for _, v := range files {
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(v.Field), escapeQuotes(v.FileName)))
		w, err := writer.CreatePart(h)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(w, v.Content)
		if err != nil {
			return nil, err
		}
	}

	headers := make(map[string]string)

	headers["Content-Type"] = writer.FormDataContentType()
	err := writer.Close()
	if err != nil {
		return nil, err
	}
	return s.Do("POST", url, headers, body)
}

func (s *HttpRequest) GetJsoniter() jsoniter.API {
	return s.json
}

type Caller struct {
	client  *HttpRequest
	baseUrl string
	headers map[string]string
	cookies map[string]string
}

func NewCaller(baseUrl string, headers map[string]string, cookies map[string]string) *Caller {
	return &Caller{
		client:  NewHttpRequest(headers, cookies, nil),
		baseUrl: baseUrl,
		headers: headers,
		cookies: cookies,
	}
}

func (s *Caller) SetBearer(bearer string) {
	if s.headers == nil {
		s.headers = make(map[string]string)
	}
	s.headers["Authorization"] = "Bearer " + bearer
}

func (s *Caller) makeUrl(path string) string {
	return strings.TrimRight(s.baseUrl, "/") + filepath.Join("/", path)
}
func (s *Caller) handleResponse(resp *httpclient.Response, outputRef interface{}) ([]byte, *httpclient.Response, error) {
	bs, err := resp.ReadAll()
	if err != nil {
		return bs, resp, err
	}
	s.client.GetJsoniter().Unmarshal(bs, outputRef)
	return bs, resp, err
}

func (s *Caller) Post(path string, inputValue interface{}, outputRef interface{}) ([]byte, *httpclient.Response, error) {
	resp, err := s.client.Call("POST", s.makeUrl(path), inputValue, s.headers)
	if err != nil {
		return nil, nil, err
	}
	return s.handleResponse(resp, outputRef)
}

func (s *Caller) PostForm(path string, inputValue interface{}, outputRef interface{}) ([]byte, *httpclient.Response, error) {
	resp, err := s.client.Post(s.makeUrl(path), inputValue)
	if err != nil {
		return nil, nil, err
	}
	return s.handleResponse(resp, outputRef)
}

func (s *Caller) Get(path string, inputValue interface{}, outputRef interface{}) ([]byte, *httpclient.Response, error) {
	resp, err := s.client.Call("GET", s.makeUrl(path), inputValue, s.headers)
	if err != nil {
		return nil, nil, err
	}
	return s.handleResponse(resp, outputRef)
}

func (s *Caller) Put(path string, inputValue interface{}, outputRef interface{}) ([]byte, *httpclient.Response, error) {
	resp, err := s.client.Call("PUT", s.makeUrl(path), inputValue, s.headers)
	if err != nil {
		return nil, nil, err
	}
	return s.handleResponse(resp, outputRef)
}
func (s *Caller) Patch(path string, inputValue interface{}, outputRef interface{}) ([]byte, *httpclient.Response, error) {
	resp, err := s.client.Call("PATCH", s.makeUrl(path), inputValue, s.headers)
	if err != nil {
		return nil, nil, err
	}
	return s.handleResponse(resp, outputRef)
}

func (s *Caller) Delete(path string, inputValue interface{}, outputRef interface{}) ([]byte, *httpclient.Response, error) {
	resp, err := s.client.Call("DELETE", s.makeUrl(path), inputValue, s.headers)
	if err != nil {
		return nil, nil, err
	}
	return s.handleResponse(resp, outputRef)
}

func (s *Caller) Client() *HttpRequest {
	return s.client
}

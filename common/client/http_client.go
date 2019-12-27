package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type HttpClient interface {
	Post(url string, in, out interface{}, header map[string]string) error
	Get(url string, in, out interface{}, header map[string]string) error
	Forward(url, method string, in io.Reader, header map[string]string, ctx context.Context) (*http.Response, error)
	Do(url, method string, in, out interface{}, header map[string]string) error

	SetTransportForTest(t http.RoundTripper)
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient() HttpClient {
	c := new(httpClient)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	c.client = &http.Client{Transport: tr}
	return c
}

func (c *httpClient) SetTransportForTest(t http.RoundTripper) {
	c.client.Transport = t
}

func (c *httpClient) Post(url string, in, out interface{}, header map[string]string) error {
	return c.Do(url, http.MethodPost, in, out, header)
}

func (c *httpClient) Get(url string, in, out interface{}, header map[string]string) error {
	return c.Do(url, http.MethodGet, in, out, header)
}

func (c *httpClient) Forward(url, method string, in io.Reader, header map[string]string, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest(method, url, in)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if val, ok := header["Content-Length"]; ok {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		req.ContentLength = contentLength
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	return c.client.Do(req)
}

func (c *httpClient) Do(url, method string, in, out interface{}, header map[string]string) error {
	postData, err := json.Marshal(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(postData))
	if err != nil {
		return err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %v, body %s", resp.StatusCode, data)
	}

	err = json.Unmarshal(data, out)
	return err
}

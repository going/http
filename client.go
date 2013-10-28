/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          client.go
 * Description:   httplib client
 */

package httplib

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type status int

type Client struct {
	conn *http.Client
}

func NewClient() *Client {
	return &Client{
		conn: new(http.Client),
	}
}

func NewProxyClient(proxy string) *Client {
	proxyURL, _ := url.Parse(proxy)
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	return &Client{
		conn: &http.Client{
			Transport: transport,
		},
	}
}

func NewSession() *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		conn: &http.Client{
			Jar: jar,
		},
	}
}

func NewProxySession(proxy string) *Client {
	proxyURL, _ := url.Parse(proxy)
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	jar, _ := cookiejar.New(nil)
	return &Client{
		conn: &http.Client{
			Jar:       jar,
			Transport: transport,
		},
	}
}

func (c *Client) Do(method, url string, headers map[string][]string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)
	if checkError(err) {
		return nil, err
	}

	if headers != nil {
		for key, v := range headers {
			for _, val := range v {
				if key == "Content-Type" {
					req.Header.Set(key, val)
					continue
				}
				req.Header.Add(key, val)
			}
		}
	}

	resp, err := c.conn.Do(req)
	if checkError(err) {
		return nil, err
	}

	return &Response{resp}, nil
}

func (c *Client) GET(url string, headers map[string][]string) (*Response, error) {
	return c.Do("GET", url, headers, nil)
}

func (c *Client) POST(url string, headers map[string][]string, body io.Reader) (*Response, error) {
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = []string{"application/x-www-form-urlencoded; param=value"}
	}
	return c.Do("POST", url, headers, body)
}

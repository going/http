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

func NewSession() *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		conn: &http.Client{
			Jar: jar,
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
	return c.Do("POST", url, headers, body)
}

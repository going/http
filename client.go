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
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"
)

type status int

type Client struct {
	sync.Mutex
	Conn    *http.Client
	session http.Header
}

func NewClient() *Client {
	return &Client{
		Conn: &http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					deadline := time.Now().Add(5 * time.Second)
					conn, err := net.DialTimeout(network, addr, 5*time.Second)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(deadline)
					return conn, nil
				},
				ResponseHeaderTimeout: 5 * time.Second,
			},
		},
	}
}

func NewProxyClient(proxy string) *Client {
	proxyURL, _ := url.Parse(proxy)
	return &Client{
		Conn: &http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					deadline := time.Now().Add(5 * time.Second)
					conn, err := net.DialTimeout(network, addr, 5*time.Second)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(deadline)
					return conn, nil
				},
				Proxy: http.ProxyURL(proxyURL),
				ResponseHeaderTimeout: 5 * time.Second,
			},
		},
	}
}

func NewSession() *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		Conn: &http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					deadline := time.Now().Add(5 * time.Second)
					conn, err := net.DialTimeout(network, addr, 5*time.Second)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(deadline)
					return conn, nil
				},
				ResponseHeaderTimeout: 5 * time.Second,
			},
			Jar: jar,
		},
		session: make(http.Header),
	}
}

func NewProxySession(proxy string) *Client {
	proxyURL, _ := url.Parse(proxy)
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	jar, _ := cookiejar.New(nil)
	return &Client{
		Conn: &http.Client{
			Jar: jar,
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					deadline := time.Now().Add(5 * time.Second)
					conn, err := net.DialTimeout(network, addr, 5*time.Second)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(deadline)
					return conn, nil
				},
				Proxy: http.ProxyURL(proxyURL),
				ResponseHeaderTimeout: 5 * time.Second,
			},
		},
		session: make(http.Header),
	}
}

func (c *Client) Do(method, url string, headers map[string][]string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)
	if checkError(err) {
		return nil, err
	}
	c.Lock()
	defer c.Unlock()

	if c.session != nil {
		if headers != nil {
			for key, v := range headers {
				for _, val := range v {
					c.session.Set(key, val)
				}
			}
		}

		for key, v := range c.session {
			for _, val := range v {
				req.Header.Set(key, val)
			}
		}

	}

	if headers != nil {
		for key, v := range headers {
			for _, val := range v {
				req.Header.Set(key, val)
			}
		}
	}

	resp, err := c.Conn.Do(req)
	if checkError(err) {
		return nil, err
	}

	return &Response{resp}, nil
}

func (c *Client) GET(url string, headers map[string][]string) (*Response, error) {
	return c.Do("GET", url, headers, nil)
}

func (c *Client) POST(url string, headers map[string][]string, body io.Reader) (*Response, error) {
	if headers == nil {
		headers = make(map[string][]string)
	}
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	}
	return c.Do("POST", url, headers, body)
}

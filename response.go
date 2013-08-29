/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          response.go
 * Description:   response
 */

package httplib

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

func (r *Response) String() (string, error) {
	b, err := r.Bytes()
	if checkError(err) {
		return "", err
	}
	return string(b), nil
}

func (r *Response) Bytes() ([]byte, error) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if checkError(err) {
		return nil, err
	}
	return b, nil
}

/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          httplib.go
 * Description:   httplib
 */

package httplib

import (
	"io"
)

var DefaultClient = NewClient()

func GET(w io.Writer, url string) (int64, error) {
	resp, err := DefaultClient.GET(url, nil)
	if checkError(err) {
		return 0, err
	}
	defer resp.Body.Close()

	return io.Copy(w, resp.Body)
}

func POST(url string, r io.Reader) error {
	resp, err := DefaultClient.POST(url, nil, r)
	if checkError(err) {
		return err
	}
	defer resp.Body.Close()

	return nil
}

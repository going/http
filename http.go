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

// DefaultClient is the default http Client used by this package.
// It's defaults are expected to represent the best practice
// at the time, but may change over time. If you need more
// control or reproducibility, you should construct your own client.
var DefaultClient = NewClient()

// Get issues a GET request using the DefaultClient and writes the result to
// to w if successful. If the status code of the response is not a success (see
// Success.IsSuccess()) no data will be written and the status code will be
// returned as an error.
func GET(w io.Writer, url string) (int64, error) {
	resp, err := DefaultClient.GET(url, nil)
	if checkError(err) {
		return 0, err
	}
	defer resp.Body.Close()

	return io.Copy(w, resp.Body)
}

// Post issues a POST request using the DefaultClient using r as the body.
// If the status code was not a success code, it will be returned as an error.
func POST(url string, r io.Reader) error {
	resp, err := DefaultClient.POST(url, nil, r)
	if checkError(err) {
		return err
	}
	defer resp.Body.Close()

	return nil
}

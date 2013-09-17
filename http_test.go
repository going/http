/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          http_test.go
 * Description:   test
 */

package httplib

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHTTPGet(t *testing.T) {
	var buf bytes.Buffer
	_, err := GET(&buf, "http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

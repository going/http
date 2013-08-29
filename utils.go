/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          utils.go
 * Description:   utils
 */

package httplib

import (
	"github.com/going/toolkit/log"
)

func checkError(err error) bool {
	if err != nil {
		log.Error(err)
		return true
	}
	return false
}

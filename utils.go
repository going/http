/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          utils.go
 * Description:   utils
 */

package httplib

func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

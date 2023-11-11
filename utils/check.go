package utils

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

func EmailCheck(email string) bool {
	r, _ := regexp.MatchString("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$", email)
	return r
}
func PasswordCheck(password string) bool {
	r, _ := regexp2.Compile("^(?=.*[a-zA-Z])(?=.*\\d)[a-zA-Z\\d!@#$%^&*()-_=+{};:,<.>]{6,20}$", 0)
	flag, _ := r.MatchString(password)
	return flag
}
func UsernameCheck(username string) bool {
	r, _ := regexp2.Compile("^[a-zA-Z0-9]{6,20}$", 0)
	flag, _ := r.MatchString(username)
	return flag
}
func NickNameCheck(nickname string) bool {
	r, _ := regexp2.Compile("^[a-zA-Z0-9!@#%^&*()-=_+{}\\[\\]|\\\\;:'\",.<>/?]{2,20}", 0)
	flag, _ := r.MatchString(nickname)
	return flag
}

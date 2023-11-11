package utilsTest

import (
	"Blog/utils"
	"fmt"
	"testing"
)

func Test_Email(t *testing.T) {
	fmt.Println(utils.EmailCheck("123456"))
	fmt.Println(utils.EmailCheck("123456ghgkj@163.com"))
	fmt.Println(utils.EmailCheck("123456@asd.sc"))
	fmt.Println(utils.EmailCheck("1710718510@qq.com"))
}
func Test_Password(t *testing.T) {
	//fmt.Println(utils.PasswordCheck("123456"))
	//fmt.Println(utils.PasswordCheck("123456q"))
	//fmt.Println(utils.PasswordCheck("123456Qa@"))
	//fmt.Println(utils.PasswordCheck("12345Q6#q"))
	//fmt.Println(utils.PasswordCheck("q132@"))
	//fmt.Println(utils.PasswordCheck("133545465464"))
	//_, err := regexp2.Compile("^(?=.*[a-zA-Z])(?=.*\\d)[a-zA-Z\\d!@#$%^&*()-_=+{};:,<.>]{6,20}$", 0)
	//if err != nil {
	//	fmt.Println(err)
	//}
	str := "   "
	fmt.Println(len(str))
}

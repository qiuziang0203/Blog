package utilsTest

import (
	"Blog/utils"
	"fmt"
	"testing"
)

func Test_Token(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjIsImV4cCI6MTY5ODgyNDkzMiwiaXNzIjoicXphIiwibmJmIjoxNjk4ODE0MTMyfQ.tZxAZ-0MOnb8vfz_Yf-9DC2ERYeupQJTVomq8q7o10w"
	id, err := utils.ParseToken(token)
	fmt.Println(id, err)
}

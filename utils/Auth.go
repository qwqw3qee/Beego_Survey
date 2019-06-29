package utils

import (
	"strconv"

	"github.com/astaxie/beego"
)

func AuthSign(ID int64, Password string) string {
	s := strconv.FormatInt(ID, 10) + beego.Substr(Password, 0, 13)
	return SHA256Encode(s)
}

func AuthSignCheck(ID int64, Password string, sign string) bool {
	s := AuthSign(ID, Password)
	return s == sign
}

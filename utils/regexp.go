package utils

import "regexp"

func CheckUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	space := regexp.MustCompile(" ")
	special := regexp.MustCompile("[-_+=^a-zA-Z0-9]")
	if space.MatchString(username) {
		return false
	}
	r := special.ReplaceAllString(username, "")
	return r == ""
}

func MatchVerifyCode(vCode string) bool {
	reg := regexp.MustCompile(`^[0-9]{4,6}$`)
	return reg.MatchString(vCode)
}

func MatchEmailFormat(email string) bool {
	reg := regexp.MustCompile(`^[0-9a-z][0-9a-z-_.]{0,35}@([0-9a-z][0-9a-z-]{0,35}[0-9a-z]\.){1,5}[a-z]{2,4}$`)
	return reg.MatchString(email)
}

// MatchPhoneNumber 目前只能匹配国内电话 +86xxxxx
func MatchPhoneNumber(phone string) bool {
	reg := regexp.MustCompile(`^\+861[3-9][0-9]{8}[0-9]$`)
	return reg.MatchString(phone)
}

// CheckPasswordStrength 检测密码强度
func CheckPasswordStrength(password string) bool {
	if len(password) < 6 {
		return false // 长度大于 6
	}
	A := regexp.MustCompile(`[A-Z]`)
	a := regexp.MustCompile(`[a-z]`)
	figure := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[!@#$%^&*()\-_+=\\|\[\]{}:'",<.>/?]`) // 除掉 ; sql注入常用符号
	if !A.MatchString(password) {
		return false // 必须要有大写字母
	}
	if !a.MatchString(password) {
		return false // 必须要有小写字母
	}
	if !figure.MatchString(password) {
		return false // 必须要有数字
	}
	if !special.MatchString(password) {
		return false // 必须要有特殊字符
	}
	if len(special.ReplaceAll(figure.ReplaceAll(a.ReplaceAll(A.ReplaceAll([]byte(password), []byte("")), []byte("")), []byte("")), []byte(""))) != 0 {
		return false // 不能有其他字符
	}
	return true
}

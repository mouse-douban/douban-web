package test

import (
	"douban-webend/model"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestVerifyPassword(t *testing.T) {
	user := model.User{PlaintPassword: "114514"}
	hash := user.EncryptPassword()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte("114514"))
	if err != nil {
		t.Error(err)
	}
}

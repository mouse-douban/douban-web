package utils

import (
	"douban-webend/utils"
	"testing"
)

func TestJWTSignAndAuthorization(t *testing.T) {
	var inputs = []int64{
		114514,
		998833,
		21,
		312,
	}
	for _, input := range inputs {
		accessToken, refreshToken, err := utils.GenerateTokenPair(input)
		if err != nil {
			t.Error("Error in signing jwt, uid = ", input)
		}
		err, uid, _ := utils.AuthorizeJWT(accessToken)
		if err != nil {
			t.Error("Error in authorizing access jwt")
		}
		if uid != input {
			t.Error("Error in authorizing access jwt, uid = ", input)
		}
		err, uid, _ = utils.AuthorizeJWT(refreshToken)
		if err != nil {
			t.Error("Error in authorizing refresh jwt")
		}
		if uid != input {
			t.Error("Error in authorizing refresh jwt, uid = ", input)
		}
	}

}

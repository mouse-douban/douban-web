package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"douban-webend/config"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// 又到了经典的造轮子环节
// 签名算法使用的是 HMAC SHA256

const (
	HeaderPlain = "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
)

// Claims Payload
type Claims struct {
	InterArrivalTime int64 `json:"iat"` // 到达时间
	ExpirationDate   int64 `json:"exp"` // 认证时间
	Uid              int64 `json:"uid"` // 用户 id
}

func GenerateTokenPair(uid int64) (accessToken, refreshToken string, err error) {
	now := time.Now().Unix()
	tokenClaims := Claims{
		InterArrivalTime: now,
		ExpirationDate:   config.Config.JwtTimeOut,
		Uid:              uid,
	}

	// refreshToken时间会比token要长
	refreshTokenClaim := Claims{
		InterArrivalTime: now,
		ExpirationDate:   config.Config.JwtTimeOut * 10,
		Uid:              uid,
	}

	accessToken, err = generateByClaims(tokenClaims)
	if err != nil {
		return "", "", ServerInternalError
	}
	refreshToken, err = generateByClaims(refreshTokenClaim)
	if err != nil {
		return "", "", ServerInternalError
	}
	return
}

func generateByClaims(claims Claims) (string, error) {
	bytes, err := json.Marshal(claims)
	if err != nil {
		return "", ServerInternalError
	}

	header := base64.StdEncoding.EncodeToString([]byte(HeaderPlain))
	payload := base64.StdEncoding.EncodeToString(bytes)

	return signJWT(header, payload), nil
}

func signJWT(header, payload string) string {

	hash := hmac.New(sha256.New, []byte(config.Config.JwtKey))
	hash.Write([]byte(header + payload))

	signed := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(hash.Sum(nil))))

	return header + "." + payload + "." + signed
}

// AuthorizeJWT 验证 JWT
// 返回：
// - nil, uid 验证成功
// - ServerInternalError 服务器错误
// - ServerError
// - 	- 40002 JWT 认证错误
// -    - 40003 JWT 过期
func AuthorizeJWT(jwtStr string) (error, int64) {

	reg := regexp.MustCompile(`\.`)
	find := reg.FindAllString(jwtStr, -1)
	if len(find) != 2 {
		return ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40002,
			Info:       "invalid jwt",
			Detail:     "JWT 认证错误!",
		}, 0
	}

	claims := Claims{}

	parts := strings.Split(jwtStr, ".")

	payload, _ := base64.StdEncoding.DecodeString(parts[1])
	signed := parts[2]
	dSigned := strings.Split(signJWT(parts[0], parts[1]), ".")[2]
	if signed != dSigned {
		return ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40002,
			Info:       "invalid jwt",
			Detail:     "JWT 认证错误!",
		}, 0
	}
	err := json.Unmarshal(payload, &claims)
	if err != nil {
		return ServerInternalError, 0
	}
	now := time.Now().Unix()

	// 如果现在的时间减去上一次登录时间大于认证时间
	if claims.ExpirationDate < now-claims.InterArrivalTime {
		return ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40003,
			Info:       "invalid jwt",
			Detail:     "JWT 过期!",
		}, 0
	}
	return nil, claims.Uid
}

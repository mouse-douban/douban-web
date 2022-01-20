package config

import (
	"encoding/json"
	"log"
	"os"
)

type Template struct {
	DefaultDbName           string `json:"default_db_name"`
	DefaultIpAndPort        string `json:"default_ip_and_port"`
	DefaultRoot             string `json:"default_root"`
	DefaultPassword         string `json:"default_password"`
	DefaultCharset          string `json:"default_charset"`
	JwtKey                  string `json:"jwt_key"`
	JwtTimeOut              int64  `json:"jwt_time_out"`
	PasswordKey             string `json:"password_key"`
	GithubOauthClientId     string `json:"github_oauth_client_id"`
	GithubOauthClientSecret string `json:"github_oauth_client_secret"`
	GiteeOauthClientId      string `json:"gitee_oauth_client_id"`
	GiteeOauthClientSecret  string `json:"gitee_oauth_client_secret"`
	ServerIp                string `json:"server_ip"`
	TencentAppId            string `json:"tencent_app_id"`
	TencentSecretId         string `json:"tencent_secret_id"`
	TencentSecretKey        string `json:"tencent_secret_key"`
	TencentLogBucketUrl     string `json:"tencent_log_bucket_url"`
	TencentSmsSdkAppId      string `json:"tencent_sms_sdk_app_id"`
	TencentSignName         string `json:"tencent_sign_name"`
	TencentTemplateId       string `json:"tencent_template_id"`
	EmailAuthAccount        string `json:"email_auth_account"`
	EmailAuthSender         string `json:"email_auth_sender"`
	EmailAuthPassword       string `json:"email_auth_password"`
	RedisAddr               string `json:"redis_addr"`
	RedisAddrInner          string `json:"redis_addr_inner"`
	RedisPassword           string `json:"redis_password"`
}

var Config Template

func Init(configPath string) {
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		log.Fatalln(err)
	}
}

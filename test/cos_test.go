package test

import (
	"douban-webend/config"
	"douban-webend/utils"
	"strings"
	"testing"
)

func TestUploadFromReader(t *testing.T) {
	config.Init("../config/config.json") // 运行 test 的根目录不在项目的根目录那
	utils.UploadFile(config.Config.TencentLogBucketUrl, "test_1.txt", strings.NewReader("hi i am test1"))
}

func TestUploadFromLocal(t *testing.T) {
	config.Init("../config/config.json") // 运行 test 的根目录不在项目的根目录那
	utils.UploadFileFromLocal(config.Config.TencentLogBucketUrl, "test_2.txt", "./test_2.txt")
}

func TestDelete(t *testing.T) {
	config.Init("../config/config.json") // 运行 test 的根目录不在项目的根目录那
	utils.DeleteFile(config.Config.TencentLogBucketUrl, "test_1.txt")
	utils.DeleteFile(config.Config.TencentLogBucketUrl, "test_2.txt")
}

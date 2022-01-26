package test

import (
	"douban-webend/utils"
	"testing"
)

func TestMatchEmail(t *testing.T) {
	var inputs = []string{
		"1545766400@qq.com",
		"igxnon@gmail.com",
		"114514@163.com;DELETE * FROM user", // 防止 sql 注入
		"918845478@qq.com;",
		"918845478&qq.com",
	}
	var outputs = []bool{
		true,
		true,
		false,
		false,
		false,
	}
	for i, input := range inputs {
		if outputs[i] != utils.MatchEmailFormat(input) {
			t.Error("Error in utils.MatchEmailFormat(" + input + ")")
		}
	}
}

func TestMatchPhone(t *testing.T) {
	var inputs = []string{
		"+8617805691274",
		"+8618303759304",
		"+8619420128402;DELETE * FROM user;32", // 防止 sql 注入
		"+8611134712830",
		"+325134712830",
		"+8625134712830",
		"+8615683055233",
	}
	var outputs = []bool{
		true,
		true,
		false,
		false,
		false,
		false,
		true,
	}
	for i, input := range inputs {
		if outputs[i] != utils.MatchPhoneNumber(input) {
			t.Error("Error in utils.MatchPhoneNumber(" + input + ")")
		}
	}
}

func TestUserName(t *testing.T) {
	var inputs = []string{
		"iGxnon",
		"_iGxnon_^v^",
		"Hacker;DELETE * FROM user;32", // 防止 sql 注入
		"åede))",
		"0912ii",
		"+(+++)+",
	}
	var outputs = []bool{
		true,
		true,
		false,
		false,
		true,
		false,
	}
	for i, input := range inputs {
		if outputs[i] != utils.CheckUsername(input) {
			t.Error("Error in utils.CheckUsername(" + input + ")")
		}
	}
}

func TestCheckPassword(t *testing.T) {
	var inputs = []string{
		"iwHwEw8921tW+-^^",
		"fDenIAwVio(08*|+wwA",
		"cMewDah*io_2*|+wffA;DELETE * FROM user;32", // 不允许出现空格防止 sql 注入
		"åçßeiuwe8492TTfewio_)(8",                   // 有不允许字符
		"An1+",                                      // 太短
		"21831298342983",                            // 纯数字
		"aadewefauhowiai",                           // 纯小写
		"ADSUHDWOQWHDW",                             // 纯大写
	}
	var outputs = []bool{
		true,
		true,
		false,
		false,
		false,
		false,
		false,
		false,
	}
	for i, input := range inputs {
		if outputs[i] != utils.CheckPasswordStrength(input) {
			t.Error("Error in utils.CheckPasswordStrength(" + input + ")")
		}
	}
}

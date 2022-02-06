package utils

import (
	"douban-webend/utils"
	"fmt"
	"testing"
)

// func TestGetSubjectInfo(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Error(r)
// 		}
// 	}()
// 	utils.QuerySubjectInfo("3541415")
// }

func TestGetSubjects(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	subjects := utils.QuerySubjects("恐怖", 20, 0)

	fmt.Println(subjects)
}

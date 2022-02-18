package dao

import (
	"douban-webend/model"
	"douban-webend/utils"
	"strings"
)

func SelectCelebrity(id int64) (err error, celebrity model.Celebrity) {
	sqlStr := "SELECT * FROM celebrity WHERE id = ?"
	row := dB.QueryRow(sqlStr, id)
	err = row.Scan(
		&celebrity.Id,
		&celebrity.Name,
		&celebrity.NameEn,
		&celebrity.Gender,
		&celebrity.Sign,
		&celebrity.Birth,
		&celebrity.Hometown,
		&celebrity.Job,
		&celebrity.IMDb,
		&celebrity.Brief,
		&celebrity.Avatar,
	)
	return
}

func SelectCelebrityNameLike(name string) (err error, celebrities []model.Celebrity) {
	celebrities = make([]model.Celebrity, 0)
	sqlStr := "SELECT * FROM celebrity WHERE name LIKE '%{}%' OR name_en LIKE '%{}%'"
	rows, err := dB.Query(strings.Replace(sqlStr, "{}", name, -1))
	defer utils.LoggerError("rows 关闭失败", rows)
	for rows.Next() {
		var celebrity model.Celebrity
		err = rows.Scan(
			&celebrity.Id,
			&celebrity.Name,
			&celebrity.NameEn,
			&celebrity.Gender,
			&celebrity.Sign,
			&celebrity.Birth,
			&celebrity.Hometown,
			&celebrity.Job,
			&celebrity.IMDb,
			&celebrity.Brief,
			&celebrity.Avatar,
		)
		if err != nil {
			return
		}
		celebrities = append(celebrities, celebrity)
	}
	return
}

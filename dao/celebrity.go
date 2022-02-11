package dao

import "douban-webend/model"

func SelectCelebrity(id int64) (err error, celebrity model.Celebrity) {
	sqStr := "SELECT * FROM celebrity WHERE id = ?"
	row := dB.QueryRow(sqStr, id)
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
	)
	return
}

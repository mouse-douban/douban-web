package dao

import (
	"database/sql"
	"douban-webend/model"
	"encoding/json"
	"log"
	"strings"
)

func InsertSubject(tags string) error {
	panic("TODO")
}

func SelectSubjects(tags, sortBy string) (err error, subjects []model.Movie) {
	sqlStr := "SELECT mid, tags, date, stars, detail, name, score, plot FROM subject WHERE tags LIKE '%{tags}%'"
	sqlStr = strings.Replace(sqlStr, "{tags}", tags, -1)
	sqlStr = sqlStr + "ORDER BY " + sortBy
	rows, err := dB.Query(sqlStr)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	if err != nil {
		return
	}
	for rows.Next() {
		var subject model.Movie
		var detail string
		var score string

		err = rows.Scan(
			&subject.Mid,
			&subject.Tags,
			&subject.Date,
			&subject.Stars,
			&detail,
			&subject.Name,
			&score,
			&subject.Plot,
		)
		if err != nil {
			return
		}

		err = json.Unmarshal([]byte(detail), &subject.Detail)
		if err != nil {
			return
		}
		err = json.Unmarshal([]byte(score), &subject.Score)
		if err != nil {
			return
		}

		subjects = append(subjects, subject)
	}
	return
}

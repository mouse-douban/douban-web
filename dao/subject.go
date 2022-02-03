package dao

import (
	"database/sql"
	"douban-webend/model"
	"encoding/json"
	"log"
	"strings"
)

func InsertSubject(movie model.Movie) error {
	sqlStr := "INSERT INTO subject (tags, date, detail, name, score, plot) VALUES (?, ?, ?, ?, ?, ?)"
	detail, err := json.Marshal(movie.Detail)
	if err != nil {
		return err
	}
	score, err := json.Marshal(movie.Score)
	if err != nil {
		return err
	}
	_, err = dB.Exec(sqlStr, movie.Tags, movie.Date, string(detail), movie.Name, string(score), movie.Plot)
	if err != nil {
		return err
	}
	return nil
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

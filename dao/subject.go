package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"encoding/json"
	"strings"
)

func InsertSubject(movie model.Movie) error {
	sqlStr := "INSERT INTO subject (tags, date, detail, name, score, plot, avatar) VALUES (?, ?, ?, ?, ?, ?, ?)"
	detail, err := json.Marshal(movie.Detail)
	if err != nil {
		return err
	}
	score, err := json.Marshal(movie.Score)
	if err != nil {
		return err
	}
	_, err = dB.Exec(sqlStr,
		movie.Tags,
		movie.Date,
		string(detail),
		movie.Name,
		string(score),
		movie.Plot,
		movie.Avatar,
	)
	if err != nil {
		return err
	}
	return nil
}

func SelectSubjects(tag, sortBy string) (err error, subjects []model.Movie) {
	sqlStr := "SELECT mid, tags, date, stars, detail, name, score, plot, avatar FROM subject WHERE tags LIKE '%{tag}%'"
	sqlStr = strings.Replace(sqlStr, "{tag}", tag, -1)
	sqlStr = sqlStr + "ORDER BY " + sortBy
	rows, err := dB.Query(sqlStr)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerWarning("rows 关闭异常", err)
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
			&subject.Avatar,
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

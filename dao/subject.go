package dao

import (
	"database/sql"
	"douban-webend/model"
	"douban-webend/utils"
	"encoding/json"
	"strings"
)

func InsertSubject(movie model.Movie) error {
	sqlStr := "INSERT INTO subject (tags, date, detail, name, score, plot, avatar, celebrities) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	detail, err := json.Marshal(movie.Detail)
	if err != nil {
		return err
	}
	score, err := json.Marshal(movie.Score)
	if err != nil {
		return err
	}
	celebrities, err := json.Marshal(movie.Celebrities)
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
		string(celebrities),
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

func SelectSubjectBaseInfo(mid int64) (err error, movie model.Movie) {
	sqlStr := "SELECT mid, tags, date, stars, name, avatar, detail, score, plot, celebrities FROM subject WHERE mid = ?"
	row := dB.QueryRow(sqlStr, mid)
	var detail, score, celebrities string
	err = row.Scan(
		&movie.Mid,
		&movie.Tags,
		&movie.Date,
		&movie.Stars,
		&movie.Name,
		&movie.Avatar,
		&detail,
		&score,
		&movie.Plot,
		&celebrities,
	)
	err = json.Unmarshal([]byte(detail), &movie.Detail)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(score), &movie.Score)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(celebrities), &movie.Celebrities)
	return
}

func SelectSubjectComments(mid int64, orderBy, kind string, comments *[]interface{}) (err error) {
	sqlStr := "SELECT c.id, c.mid, c.uid, c.content, c.date, c.score, u.username, c.tag, c.type, c.stars  FROM comment c JOIN user u ON c.uid = u.uid AND c.mid = ? AND c.type = ?"
	sqlStr += " ORDER BY " + orderBy
	rows, err := dB.Query(sqlStr, mid, kind)
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.LoggerWarning("rows 关闭失败!")
		}
	}(rows)
	for rows.Next() {
		var comment model.Comment
		var tag string
		err = rows.Scan(
			&comment.Id,
			&comment.Mid,
			&comment.Uid,
			&comment.Content,
			&comment.Date,
			&comment.Score,
			&comment.Username,
			&tag,
			&comment.Type,
			&comment.Stars,
		)
		if err != nil {
			return
		}
		comment.Tag = strings.Split(tag, ",")
		*comments = append(*comments, comment)
	}
	return
}

func SelectSubjectReviews(mid int64, orderBy string, comments *[]interface{}) (err error) {
	return
}

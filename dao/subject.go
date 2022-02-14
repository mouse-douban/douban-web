package dao

import (
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

func SelectSubjects(tag, sortBy string, start, limit int) (err error, subjects []model.Movie) {
	sqlStr := "SELECT mid, tags, date, stars, detail, name, score, plot, avatar FROM subject WHERE tags LIKE '%{tag}%' "
	sqlStr = strings.Replace(sqlStr, "{tag}", tag, -1)
	sqlStr = sqlStr + " ORDER BY " + sortBy + " LIMIT ? OFFSET ?"
	rows, err := dB.Query(sqlStr, limit, start)
	defer utils.LoggerError("rows 关闭异常", rows)

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

func SelectSubjectComments(mid int64, orderBy, kind string, comments *[]interface{}, start, limit int) (err error) {
	sqlStr := "SELECT c.id, c.mid, c.uid, c.content, c.date, c.score, u.username, c.tag, c.type, c.stars  FROM comment c JOIN user u ON c.uid = u.uid AND c.mid = ? AND c.type = ?"
	sqlStr += " ORDER BY " + orderBy + " LIMIT ? OFFSET ?"
	rows, err := dB.Query(sqlStr, mid, kind, limit, start)
	if err != nil {
		return
	}
	defer utils.LoggerError("rows 关闭异常", rows)
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

func SelectSubjectReviews(mid int64, orderBy string, reviews *[]interface{}, start, limit int) (err error) {
	sqlStr := "SELECT r.id, r.mid, r.uid, r.name, u.username, u.avatar, r.score, r.date, r.stars, r.bads, r.reply_cnt, r.content FROM review r JOIN user u ON u.uid = r.uid AND r.mid = ?"
	sqlStr += " ORDER BY " + orderBy + " LIMIT ? OFFSET ?"
	rows, err := dB.Query(sqlStr, mid, limit, start)
	if err != nil {
		return
	}
	defer utils.LoggerError("rows 关闭异常", rows)
	for rows.Next() {
		var review model.ReviewSnapshot
		err = rows.Scan(
			&review.Id,
			&review.Mid,
			&review.Uid,
			&review.Name,
			&review.Username,
			&review.Avatar,
			&review.Score,
			&review.Date,
			&review.Stars,
			&review.Bads,
			&review.ReplyCnt,
			&review.Brief,
		)
		if err != nil {
			return
		}
		if len(review.Brief) > 165*3 {
			review.Brief = review.Brief[:165*3]
		}
		*reviews = append(*reviews, review)
	}
	return
}

func SelectSubjectDiscussions(mid int64, orderBy string, discussions *[]interface{}, start, limit int) (err error) {
	sqlStr := "SELECT d.id, d.uid, d.mid, d.name, u.username, d.reply_cnt, d.date, u.avatar, d.stars FROM discussion d JOIN user u ON d.uid = u.uid AND d.mid = ?"
	sqlStr += " ORDER BY " + orderBy + " LIMIT ? OFFSET ?"
	rows, err := dB.Query(sqlStr, mid, limit, start)
	if err != nil {
		return
	}
	defer utils.LoggerError("rows 关闭异常", rows)
	for rows.Next() {
		var discussion model.DiscussionSnapshot
		err = rows.Scan(
			&discussion.Id,
			&discussion.Uid,
			&discussion.Mid,
			&discussion.Name,
			&discussion.Username,
			&discussion.ReplyCnt,
			&discussion.Date,
			&discussion.Avatar,
			&discussion.Stars,
		)
		if err != nil {
			return
		}
		*discussions = append(*discussions, discussion)
	}
	return
}

func UpdateSubjectScore(mid int64, score model.MovieScore) (err error) {
	sqlStr := "UPDATE subject SET score = ? WHERE mid = ?"
	scoreB, err := json.Marshal(score)
	if err != nil {
		return
	}
	_, err = dB.Exec(sqlStr, string(scoreB), mid)
	return
}

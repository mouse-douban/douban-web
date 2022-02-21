package service

import (
	"douban-webend/dao"
	"douban-webend/model"
	"douban-webend/utils"
	"strconv"
	"strings"
	"sync"
)

func GetSubjects(start, limit int, sort string, tags string) (err error, subjects []model.Movie) {
	err, subjects = dao.SelectSubjects(tags, orderBys[sort], start, limit)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40015,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
	}
	return
}

func GetSubjectBaseInfo(mid int64) (err error, movie model.Movie) {
	err, movie = dao.SelectSubjectBaseInfo(mid)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 400,
			Status:     40015,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
	}
	return err, movie
}

func GetSubjectScopeInfo(mid int64, scopes []string, info *map[string][]interface{}) (err error) {
	for _, scope := range scopes {
		switch scope {
		case "comments":
			var in = make([]interface{}, 0)
			err = GetSubjectComments(mid, &in, 0, 6, "hotest", "after")
			if err != nil {
				return
			}
			(*info)["comments"] = in
		case "reviews":
			var in = make([]interface{}, 0)
			err = GetSubjectReviews(mid, &in, 0, 6, "hotest")
			if err != nil {
				return
			}
			(*info)["reviews"] = in
		case "discussions":
			var in = make([]interface{}, 0)
			err = GetSubjectDiscussions(mid, &in, 0, 6, "hotest")
			if err != nil {
				return
			}
			(*info)["discussions"] = in
		}
	}
	return
}

func GetSubjectComments(mid int64, comments *[]interface{}, start, limit int, sort, kind string) (err error) {
	err = dao.SelectSubjectComments(mid, orderBys[sort], kind, comments, start, limit)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 40015,
			Status:     0,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	return
}

func GetSubjectReviews(mid int64, reviews *[]interface{}, start, limit int, sort string) (err error) {
	err = dao.SelectSubjectReviews(mid, orderBys[sort], reviews, start, limit)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 40015,
			Status:     0,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	return
}

func GetSubjectDiscussions(mid int64, discussions *[]interface{}, start, limit int, sort string) (err error) {
	err = dao.SelectSubjectDiscussions(mid, orderBys[sort], discussions, start, limit)
	if err != nil {
		err = utils.ServerError{
			HttpStatus: 40015,
			Status:     0,
			Info:       "invalid request",
			Detail:     "影片不存在",
		}
		return
	}
	return
}

var mu = sync.Mutex{}

// UpdateSubjectScore 更新电影评分信息
func UpdateSubjectScore(mid int64, score int) (err error) {
	mu.Lock()
	err, movie := GetSubjectBaseInfo(mid)
	if err != nil {
		return err
	}
	instance, err := strconv.ParseFloat(movie.Score.Score, 64)
	if err != nil {
		return utils.ServerInternalError
	}
	cnt := float64(movie.Score.TotalCnt)
	to := (cnt*instance*0.5 + float64(score)) / (cnt + 1) * 2 // 更新后的评分 10 分制
	toInstance := strconv.FormatFloat(to, 'f', 2, 64)

	var toScore = model.MovieScore{
		Score:    toInstance,
		TotalCnt: int(cnt + 1),
		Five:     movie.Score.Five,
		Four:     movie.Score.Four,
		Three:    movie.Score.Three,
		Two:      movie.Score.Two,
		One:      movie.Score.One,
	}

	switch score {
	case 1:
		ret, err := parsePercentage(movie.Score.One)
		if err != nil {
			return utils.ServerInternalError
		}
		toScore.One = strings.TrimLeft(strconv.FormatFloat((ret*cnt+1)/(cnt+1), 'f', 2, 64), "0.") + "%"
	case 2:
		ret, err := parsePercentage(movie.Score.Two)
		if err != nil {
			return utils.ServerInternalError
		}
		toScore.Two = strings.TrimLeft(strconv.FormatFloat((ret*cnt+1)/(cnt+1), 'f', 2, 64), "0.") + "%"
	case 3:
		ret, err := parsePercentage(movie.Score.Three)
		if err != nil {
			return utils.ServerInternalError
		}
		toScore.Three = strings.TrimLeft(strconv.FormatFloat((ret*cnt+1)/(cnt+1), 'f', 2, 64), "0.") + "%"
	case 4:
		ret, err := parsePercentage(movie.Score.Four)
		if err != nil {
			return utils.ServerInternalError
		}
		toScore.Four = strings.TrimLeft(strconv.FormatFloat((ret*cnt+1)/(cnt+1), 'f', 2, 64), "0.") + "%"
	case 5:
		ret, err := parsePercentage(movie.Score.Five)
		if err != nil {
			return utils.ServerInternalError
		}
		toScore.Five = strings.TrimLeft(strconv.FormatFloat((ret*cnt+1)/(cnt+1), 'f', 2, 64), "0.") + "%"
	}
	// TODO 减少其他的比例

	err = dao.UpdateSubjectScore(mid, toScore)

	mu.Unlock()
	return
}

func parsePercentage(v string) (ret float64, err error) {
	v = strings.Replace(v, "%", "", -1)
	ret, err = strconv.ParseFloat(v, 64)
	ret = 0.01 * ret
	return
}

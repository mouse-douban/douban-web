package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type SubjectInfo struct {
	Items []struct {
		Topic struct {
			FocusedSubject struct {
				Rating struct {
					Count     int     `json:"count"`
					Max       int     `json:"max"`
					StarCount float64 `json:"star_count"`
					Value     float64 `json:"value"`
				} `json:"rating"`
				Title        string `json:"title"`
				Url          string `json:"url"`
				CoverUrl     string `json:"cover_url"`
				CardSubtitle string `json:"card_subtitle"`
				Id           string `json:"id"`
			} `json:"focused_subject"`
		} `json:"topic"`
	} `json:"items"`
}

type Subjects struct {
	Entry []struct {
		Rate  string `json:"rate"`
		Title string `json:"title"`
		Url   string `json:"url"`
		Cover string `json:"cover"`
		Id    string `json:"id"`
	} `json:"subjects"`
}

const (
	SubjectInfoGet = "https://movie.douban.com/subject/%v/"
	SubjectsGet    = "https://movie.douban.com/j/search_subjects?type=movie&tag=%v&sort=rank&page_limit=%v&page_start=%v"
)

var client = http.Client{}

func QuerySubjectInfo(mid string) {
	req, err := http.NewRequest("GET", fmt.Sprintf(SubjectInfoGet, mid), nil)
	req.Header.Set("Referer", "from Gallifrey") // 不加这个会400，貌似设置成什么都可以 :)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15")
	if err != nil {
		panic(err)
	}

	response, err := client.Do(req)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("无效 response")
	}

	all, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(all))
}

func QuerySubjects(tag string, limit, start int) Subjects {
	req, err := http.NewRequest("GET", fmt.Sprintf(SubjectsGet, tag, limit, start), nil)
	req.Header.Set("Referer", "from Gallifrey") // :D
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15")
	if err != nil {
		panic(err)
	}

	response, err := client.Do(req)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("无效 response")
	}

	all, err := ioutil.ReadAll(response.Body)

	result, err := strconv.Unquote(strings.Replace(strconv.Quote(string(all)), `\\u`, `\u`, -1))
	if err != nil {
		panic(err)
	}

	result = strings.Replace(result, `\/`, `/`, -1)

	var ret Subjects
	err = json.Unmarshal([]byte(result), &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

package utils

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetPOSTBytesWithEmptyBody(url string) chan []byte {
	return GetPOSTBytes(url, "application/json", nil)
}

// GetPOSTBytes 返回一个 channel 写入数据 async/await 模式
func GetPOSTBytes(url, contentType string, body io.Reader) (resp chan []byte) {

	resp = make(chan []byte, 1)

	go func(ch *chan []byte, url, contentType string, body io.Reader) {
		client := &http.Client{}
		res, err := client.Post(url, contentType, body)
		if err != nil {
			*ch <- make([]byte, 0)
			return
		}

		all, err := ioutil.ReadAll(res.Body)
		if err != nil {
			*ch <- make([]byte, 0)
			return
		}

		*ch <- all

		err = res.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(&resp, url, contentType, body)

	return
}

func GetGETBytes(url string, header map[string]string) (resp chan []byte) {

	resp = make(chan []byte, 1)

	go func(ch *chan []byte, url string, header map[string]string) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			*ch <- make([]byte, 0)
			return
		}

		if header != nil {
			for k, v := range header {
				req.Header.Set(k, v)
			}
		}

		res, err := client.Do(req)
		if err != nil {
			*ch <- make([]byte, 0)
			return
		}

		all, err := ioutil.ReadAll(res.Body)
		if err != nil {
			*ch <- make([]byte, 0)
			return
		}

		*ch <- all

		err = res.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(&resp, url, header)

	return
}

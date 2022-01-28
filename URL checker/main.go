package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) { // 여기 추가된 arrow는 send only라는 뜻이다.이렇게 명시적으로 해주는거다.
	resp, err := http.Get(url) //go에서 제공하는 표준
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status} // 채널로 우리의 struct를 보내자.
}

// 이 채널은 데이터를 받을 수만 있고 보낼 수는 없어! 이런 식으로 코드상에서 명시해주는게 더 좋다.
// 최적화가 끝났다. 이제 체크하는데 가장 오래 걸리는 url 하나의 시간이 프로그램의 시간이다.

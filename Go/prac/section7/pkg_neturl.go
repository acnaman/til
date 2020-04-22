package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func infohandlerfunc(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>インフォメーション</title>
</head>
<body>
<h1>Hello, World!</h1>
</body>
</html>
`)
}

func infohandlefunccall() {
	http.HandleFunc("/info", infohandlerfunc)
	http.ListenAndServe(":8080", nil)
}

func urlpostfunc() {
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())
	res, err := http.PostForm("https://example.com/commnets/post", vs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.StatusCode)
}

func urlgetfunc() {
	res, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])
	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func urlparsefunc() {
	u, err := url.Parse("https://www.example.com/search?a=1&b=2#top")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)

	fmt.Println(u.IsAbs())
	fmt.Println(u.Query())
}

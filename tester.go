package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func test() {

	url := "http://localhost:8000/api/v1/videos"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Basic Z2lyaXB1dHJhYWRoaXR0YW5hOjEyMw==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

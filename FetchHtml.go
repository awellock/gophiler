package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func getHtml(url string) string{
	res, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
	}
	return fmt.Sprintf("boom: %s \n", body)
}

func main(){
	html := getHtml("http://localhost")
	fmt.Println(html)
}

package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title  string	`json:'title'`
	Year   int	`json:'year'`
	Price  int	`json:'price'`
	Actors []string	`json:'actors'`
}

type StudentInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	College string `json:"college"`
	Major   string `json:"major"`
}

func main() {
	movie := Movie{Title:"喜剧之王", Year:2000, Price:10, Actors:[]string{"李傲飞", "zhang3"}}
	// 编码的过程：将结构体转化为json
	fmt.Println(movie)
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal err", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程：json转化为结构体
	// jsonStr = {"Title":"喜剧之王","Year":2000,"Price":10,"Actors":["李傲飞","zhang3"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
	fmt.Println(myMovie.Title)

	studentInfo := StudentInfo{
		Id: "1", Name: "张三", Age: "21", College: "机电", Major: "机械",
	}
	studentInfoAsBytes, err := json.Marshal(studentInfo)
	fmt.Printf("str=%s\n", studentInfoAsBytes)
}
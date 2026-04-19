package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[TIME]")

	timee := time.Now()
	formatedTime := time.Now().Format("2006-01-02 15:04:05 Monday") // bruh the way of formatting

	fmt.Println(timee)
	fmt.Println(formatedTime)

	createDate := time.Date(2030, time.April, 10, 20, 4, 50, 4, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("2006-01-02 15:04:05 Monday")) // i hate GO's time formatting why not "yyyy-mm-dd day"
}

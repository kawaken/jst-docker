package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Println(t.In(jst))
}

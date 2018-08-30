package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t.In(jst))
	}

	jst = time.FixedZone("JST", 9*60*60)
	fmt.Println(t.In(jst))

	time.Local = jst
	fmt.Println(time.Now())

}

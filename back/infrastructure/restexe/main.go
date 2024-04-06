package main

import (
	"fmt"
	timezone "photo-share/back/sharelib/timezone"
)

func main() {
	timezone.SetTimeZoneAsiaTokyo()
	fmt.Println("test")
}

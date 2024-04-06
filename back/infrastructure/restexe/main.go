package main

import (
	"fmt"
	"photo-share/back/infrastructure"
	timezone "photo-share/back/sharelib/timezone"
)

func main() {
	timezone.SetTimeZoneAsiaTokyo()

	config := infrastructure.Configs{}
	fmt.Println(config)
}

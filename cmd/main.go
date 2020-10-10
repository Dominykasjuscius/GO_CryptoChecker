package main

import (
	"api"
	"time"
)

func main() {
	var curr api.API = new(api.CurrencyObj)
	err := curr.CompareData()

	for true {
		if err != "Success" {
			println(err)
			break
		}
		timer1 := time.NewTimer(30 * time.Second)
		<-timer1.C
		err = curr.CompareData()
	}

}

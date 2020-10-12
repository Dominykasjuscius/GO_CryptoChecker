package main

import (
	"gocryptochecker/api"
	"time"
)

func main() {
	var curr api.API = new(api.CurrencyObj)

	// curr.CompareData() grazina string ne error, tai kintamaji turbut reiketu
	// pavadinit kazkaip kitaip. Galbut state ar pan.
	err := curr.CompareData()

	// kaip uzduotyje buvo mineta, kad gerai butu panaudoti os.Signal, tai cia
	// yra ta vieta. Cia tokio amzino ciklo neturetu but. Turetu but kazkas pan:
	/*
		for {
			select {
				case <-timer1.C:
					...
				case <-os.Sgnal:
					return
			}
		}
	*/
	for true {
		if err != "Success" {
			println(err)
			break
		}
		// kai galima, timeri visada reiketu sustabdyti su timer1.Stop().
		timer1 := time.NewTimer(30 * time.Second)
		<-timer1.C
		err = curr.CompareData()
	}

}

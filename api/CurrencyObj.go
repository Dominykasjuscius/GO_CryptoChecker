package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Labiau go'ish variantas yra rasyti komentarus pasinaudojant //. Pirmas zodis
// turetu but kintamojo pavadinimas arba funkcijos pavadinimas. Apie komentaru
// rasyma neblogai parasyta https://blog.golang.org/godoc
/*
	Struct for storing JSON data
*/
type CurrencyObj struct {
	Name  string `json:"name"`
	Price string `json:"price_usd"`
}

/*
	Converts the given JSON obtain from URL to a CurrencyObj Struct
*/
// Si funkcija turetu grazinti CurrencyObj, ne []CurrencyObj.
func (r CurrencyObj) ConvertJSONtoCurr(id int) []CurrencyObj {
	obj := []CurrencyObj{}
	link := "https://api.coinlore.net/api/ticker/?id=" + strconv.Itoa(id)
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	}

	// galima pernaudot ta pati kintamaji err, nebutina handlinti errora su
	// nauju kintamuoju. Jei funkcijoje yra handlinamas err kelis kartus,
	// daznai funkcijos pradzioje yra aprasomas `var err error` ir jis naudojamas
	// visoje funkcijoje
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		fmt.Println(err)
	}

	return obj
}

/*
	Compares the the rules obtained from a local JSON file to the prices
	obtained from URL
*/
// CompareData funkcija turetu grazinti error tipa ne string. Galimas klaidas
// galima aprasyti paketo lygmenyje e.g.:
// https://github.com/docker/compose-cli/blob/main/backend/backend.go#L35
func (r CurrencyObj) CompareData() string {

	var curr API = new(CurrencyObj)
	var rule RulesObj

	var rules = rule.ConvertJSONtoRules()
	if len(rules) == 0 {
		return "\nNo rules in rules.json"
	}
	var hasChanged bool = false

	// reiketu naudoti for _, v := range rules { .. }
	for i := 0; i < len(rules); i++ {
		var currency = curr.ConvertJSONtoCurr(rules[i].Id)
		num, _ := strconv.ParseFloat(currency[0].Price, 64)
		if rules[i].Rule == "gt" {
			if rules[i].Price < num {
				fmt.Printf("Cryptocurrency id:%v %v price is greater than\n%v\n", rules[i].Id, currency[0].Name, rules[i].Price)
				if _, err := rules[i].RemoveElement(i); err != nil {
					// .. handle error
				}
				hasChanged = true
			}
		} else {
			if rules[i].Price > num {
				fmt.Printf("Cryptocurrency id:%v %v price is less than\n%v\n", rules[i].Id, currency[0].Name, rules[i].Price)
				if _, err := rules[i].RemoveElement(i); err != nil {
					// .. handle error
				}
				hasChanged = true
			}
		}
	}
	if !hasChanged {
		println("No changes in price")
	}
	return "Success"
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CurrencyObj struct {
	Name  string `json:"name"`
	Price string `json:"price_usd"`
}

func (r CurrencyObj) ConvertJSONtoCurr(id int) []CurrencyObj {
	var obj []CurrencyObj
	var link = "https://api.coinlore.net/api/ticker/?id=" + strconv.Itoa(id)
	var resp, err = http.Get(link)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	err1 := json.NewDecoder(resp.Body).Decode(&obj)

	if err1 != nil {
		fmt.Println(err1)
	}

	return obj
}

func (r CurrencyObj) CompareData() string {

	var curr API = new(CurrencyObj)
	var rule RulesObj

	var rules = rule.ConvertJSONtoRules()
	if len(rules) == 0 {
		return "\nNo rules in rules.json"
	}
	var hasChanged bool = false

	for i := 0; i < len(rules); i++ {
		var currency = curr.ConvertJSONtoCurr(rules[i].Id)
		num, _ := strconv.ParseFloat(currency[0].Price, 64)
		if rules[i].Rule == "gt" {
			if rules[i].Price < num {
				fmt.Printf("Cryptocurrency id:%v %v price is greater than\n%v\n", rules[i].Id, currency[0].Name, rules[i].Price)
				rules[i].RemoveElement(i)
				hasChanged = true
			}
		} else {
			if rules[i].Price > num {
				fmt.Printf("Cryptocurrency id:%v %v price is less than\n%v\n", rules[i].Id, currency[0].Name, rules[i].Price)
				rules[i].RemoveElement(i)
				hasChanged = true
			}
		}
	}
	if !hasChanged {
		println("No changes in price")
	}
	return "Success"
}

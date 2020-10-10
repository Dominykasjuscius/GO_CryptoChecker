package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var rules []RulesObj

const path = "../configs/rules.json"

type RulesObj struct {
	Id    int     `json:"crypto_id"`
	Price float64 `json:"price"`
	Rule  string  `json:"rule"`
}

/*
	Convert local JSON file into RulesObj struct
*/
func (r RulesObj) ConvertJSONtoRules() []RulesObj {
	byteValue := openJson()

	err := json.Unmarshal([]byte(byteValue), &rules)

	if err != nil {
		fmt.Println(err)
	}
	return rules
}

/*
	Dynamically removes object with given index from slice
*/
func (r RulesObj) RemoveElement(index int) []RulesObj {
	length := len(rules)
	if length == 0 {
		return rules
	}

	byteValue := openJson()
	err := json.Unmarshal([]byte(byteValue), &rules)

	if length > 1 {
		rules = append(rules[:index], rules[index+1:]...)
	} else {
		rules = rules[:0]
	}

	result, err := json.Marshal(rules)
	err = ioutil.WriteFile(path, result, 0644)

	if err != nil {
		fmt.Println(err)
	}
	return rules
}

/*
	Opens and reads the bytes from a JSON file using the give path
*/
func openJson() []byte {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

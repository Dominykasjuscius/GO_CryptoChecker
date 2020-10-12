package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// define package level errors.
var (
	ErrNoRules      = errors.New("no rules")
	ErrRuleNotFount = errors.New("rule not found")
)

var rules []RulesObj

const path = "configs/rules.json"

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
// RemoveElement funkcija neturetu but prikabinta prie `r RulesObj`, kadangi
// funkcija neturi nieko bendro su RulesObj, o bando modifikuot globalu
// kintamaji. Si funkcija turetu but perrasyta labiau funkciniu stilium:
// RemoveElement(r []RulesObj, idx int) ([]RulesObj, error)
// tokia funkcija lengviau testuoti.
func (r RulesObj) RemoveElement(index int) ([]RulesObj, error) {
	var (
		err  error
		data []byte
	)

	if len(rules) == 0 {
		return nil, ErrNoRules
	}
	if data, err = ioutil.ReadFile(path); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &rules); err != nil {
		return nil, err
	}

	if len(rules) > 1 && len(rules) >= index {
		// cia yra blogai. reiketu sukurti nauja tuscia slice
		// ir ten appendinti rules subseta.
		rules = append(rules[:index], rules[index+1:]...)
	}

	// galima pernaudot data kitnamaji
	if data, err = json.Marshal(rules); err != nil {
		return nil, err
	}
	if err = ioutil.WriteFile(path, data, 0644); err != nil {
		return nil, err
	}

	return rules, nil
}

/*
	Opens and reads the bytes from a JSON file using the give path
*/
// si funkcija yra nereikalinga. Galima naudoti ioutils.ReadFile funkcija
// kuri darys ta pati.
func openJson() []byte {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

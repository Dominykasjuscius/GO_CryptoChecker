package api

type API interface {
	ConvertJSONtoCurr(id int) []CurrencyObj
	CompareData() string
}

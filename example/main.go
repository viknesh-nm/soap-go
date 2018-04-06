package main

import (
	"encoding/xml"
	"fmt"

	"github.com/viknesh-nm/soap-go"
)

// Result -
type Result struct {
	Celsius float64 `xml:"FahrenheitToCelsiusResult"`
}

func main() {
	res, err := soap.NewClient("https://www.w3schools.com/xml/tempconvert.asmx?WSDL")
	if err != nil {
		fmt.Printf("error not expected: %s", err)
		return
	}

	params := soap.Params{
		"Fahrenheit": "100",
	}

	v := &Result{}

	se, err := res.GetData("FahrenheitToCelsius", params)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = xml.Unmarshal([]byte(se), v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Celcius", v.Celsius)
}

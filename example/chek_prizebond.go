package main

import (
	"fmt"

	prizebond "github.com/mhshajib/bd-prizebond-checker"
)

func init() {
	//Setting Prizebond Interface
	prizebond.Prizebond = &prizebond.PrizebondConnection{}

	prizebondNumbers := []string{"XXXXXXX"} //Place your prizebond number here

	//Initializing Bangladesh band api as prizebond gateway
	prizebond.Prizebond.Init(prizebondNumbers)
}

func main() {
	var prizeBondData []map[string]string

	//Sendding sms
	prizeBondData, err := prizebond.Prizebond.Fetch()
	if err != nil {
		//If some error happens during fetching results
		fmt.Println(err)
	}

	//API called successfully. Printing API Response in map.
	fmt.Println(prizeBondData)
}

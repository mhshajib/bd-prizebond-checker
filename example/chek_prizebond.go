package main

import (
	"fmt"

	prizebond "github.com/mhshajib/bd-prizebond-checker"
)

func init() {
	//Setting Prizebond Interface
	prizebond.Prizebond = &prizebond.PrizebondConnection{}

	prizebondNumbers := []string{"0920249", "0030401"}

	//Initializing Bangladesh band api as prizebond gateway
	prizebond.Prizebond.Init(prizebondNumbers)
}

func main() {
	var prizeBondData []map[string]string

	//Sendding sms
	prizeBondData, err := prizebond.Prizebond.Fetch()
	if err != nil {
		//If something happen before sending sms
		fmt.Println("Something Went Wrong")
	}

	//API called successfully. Printing API Response in map.
	fmt.Println(prizeBondData)
}

# GO Package For Checking Bangladeshi Prizebond Number


### Installation

Install the package using
```go
$ go get github.com/mhshajib/bd-prizebond-checker
```

### Usage

To use the package import it in your `*.go` code
```go
import prizebond "github.com/mhshajib/bd-prizebond-checker"
```

### Example

***Note: The Example Numbers Are `dummy` Please Use Your Original Prizebond numbers to get proper results***

***You can add multiple bond numbers as array***

```go
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

```

You'll get response as map


### **License**
The **bd-prizebond-checker** is an open-source software licensed under the [MIT License](LICENSE)
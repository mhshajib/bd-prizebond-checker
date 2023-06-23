package prizebond

// Interface of Prizebond
type IPrizebond interface {
	Init(bondNumbers []string)
	Fetch() ([]map[string]string, error)
}

// use this variable to access the implementation of this interface
var Prizebond IPrizebond

package utils

const (
	previousVar = "$pr"
)

var (
	reservedVarNames = map[string]bool{
		previousVar: true,
	}

	vars = map[string]float64{}
)

func GetPreviousResult() float64 {
	return GetVar(previousVar)
}

func SetPreviousResult(value float64) {
	vars[previousVar] = value
}

func SetVar(name string, value float64) {
	if reservedVarNames[name] {
		panic("variable name is reserved: " + name)
	}
	vars[name] = value
}

func GetVar(name string) float64 {
	if value, ok := vars[name]; ok {
		return value
	} else {
		panic("variable is not defined: " + name)
	}
}

func GetVars() map[string]float64 {
	return vars
}

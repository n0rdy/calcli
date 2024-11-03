package utils

var (
	// KnownCalculatorTokens set manually for now, let's see if I can automate this to keep in sync with the grammar
	KnownCalculatorTokens = []string{
		"abs(",
		"acos(",
		"asin(",
		"atan(",
		"ceil(",
		"cos(",
		"e",
		"exp(",
		"exp2(",
		"floor(",
		"ln(",
		"log(",
		"log2(",
		"log10(",
		"mod(",
		"nrt(",
		"percent(",
		"pi",
		"pmem()",
		"round(",
		"$pr",
		"sin(",
		"sqrt(",
		"tan(",
	}
)

func AddToken(token string) {
	KnownCalculatorTokens = append(KnownCalculatorTokens, token)
}

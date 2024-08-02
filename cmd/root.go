package cmd

import (
	"bufio"
	"fmt"
	"math"
	"n0rdy.foo/calcli/calc"
	"n0rdy.foo/calcli/calc/parser"
	"n0rdy.foo/calcli/calc/utils"
	"os"

	"github.com/spf13/cobra"
)

const (
	exitCmd = "exit"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calcli",
	Short: "A CLI calculator app",
	Long: `A CLI calculator app.

Supports basic arithmetic operations: +, -, *, /, %, ^, !, as well as parentheses.
Contains a set of predefined constants: pi, e.
Supports the following math functions:
- abs(x) - the absolute value of x
- acos(x) - the arccosine of x
- asin(x) - the arcsine of x
- atan(x) - the arctangent of x
- ceil(x) - the smallest integer value greater than or equal to x
- cos(x) - the cosine of x
- exp(x) - the value of e^x
- exp2(x) - the value of 2^x
- floor(x) - the largest integer value less than or equal to x
- ln(x) - the natural logarithm of x
- log(x, base) - the logarithm of x to the specified base
- log2(x) - the base-2 logarithm of x
- log10(x) - the base-10 logarithm of x
- mod(x, y) - the remainder of the division of x by y
- nrt(x, degree) - the root of the specified degree
- percent(x, y) - the percentage of x from y
- round(x) - the value of x rounded to the nearest integer
- sin(x) - the sine of x
- sqrt(x) - the square root of x
- tan(x) - the tangent of x

The result of the previous calculation is stored in the variable '$pr'
and can be used in subsequent calculations.

Have fun =)`,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		c := calc.NewCalcProcessor()

		for {
			input := userInput()
			if input == "" {
				continue
			}
			if input == exitCmd {
				break
			}

			res, err := c.Process(input)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
				continue
			}
			processResult(*res)
			fmt.Println()
		}
	},
}

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func processResult(res parser.CalcResult) {
	if !res.HasValue {
		return
	}

	val := res.Value
	if math.IsNaN(val) {
		fmt.Println("NaN")
		return
	}
	if math.IsInf(val, -1) {
		fmt.Println("-∞")
		return
	}
	if math.IsInf(val, 1) {
		fmt.Println("+∞")
		return
	}

	utils.SetPreviousResult(val)
	fmt.Println(val)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

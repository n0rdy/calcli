package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"n0rdy.foo/calcli/calc"
	"n0rdy.foo/calcli/tui"
	"os"
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

Also, supports the following system functions:
- pmem() - prints all the variables and their values stored in memory

The result of the previous calculation is stored in the variable '$pr'
and can be used in subsequent calculations.

Have fun =)`,
	Version: "0.0.5",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(
			tui.InitialModel(calc.NewCalcProcessor()),
			tea.WithAltScreen(),
		)
		if _, err := p.Run(); err != nil {
			fmt.Println("could not run program:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"

	lab2 "github.com/Homa4/KPI3-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, error := lab2.CalculatePrefix("2 2 +")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(res)
	}
}

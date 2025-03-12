package main

import (
	"fmt"
)

/*
*
* a dirty improvement percentage calculation
*
* just for self-usage only, not for external use of distribution.
* buggy, unmaintained code
*
 */

func main() {
	var err error
	var usageDiff, diffDiv, initialPercentage, currentPercentage, improvementPercentage float64

	fmt.Println("Percentage Difference Calculator - v0.0.1-beta")
	fmt.Println("")

	fmt.Print("Enter the inital percentage: ")
	_, err = fmt.Scan(&initialPercentage)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	fmt.Print("Enter the current percentage: ")
	_, err = fmt.Scan(&currentPercentage)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	/*
		fmt.Printf("You entered: %f\n", initialPercentage)
		fmt.Printf("You entered: %f\n", currentPercentage)

		fmt.Printf("Rounded to 2 decimal places: %.2f\n", initialPercentage)
		fmt.Printf("Rounded to 2 decimal places: %.2f\n", currentPercentage)
	*/

	usageDiff = initialPercentage - currentPercentage
	diffDiv = usageDiff / initialPercentage
	improvementPercentage = diffDiv * 100

	fmt.Println("=========================================")

	fmt.Printf("The improvement in percentage is : %.2f\n", improvementPercentage)

}

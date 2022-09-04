package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// create files
	f, err := os.Create("two-step-equations.html")
	fs, err2 := os.Create("two-step-equation-solutions.html")

	// error checking
	if err != nil {
		fmt.Println(err)
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	// create variables and assign them to 0
	num1, num2, num3 := 0, 0, 0
	probStr, solStr := "", ""

	// create for loop for files
	for i := 1; i < 101; i++ {
		// create random numbers
		num1 = rand.Intn(20) + 1
		num2 = rand.Intn(20) + 1
		num3 = rand.Intn(20) + 1

		// create string for number in loop
		probNum := fmt.Sprint(i)

		// start writing to files
		_, errf := f.WriteString("<div><p>" + probNum + ") ")
		_, errfs := fs.WriteString("<div><p>" + probNum + ") ")

		// check errors
		if errf != nil {
			fmt.Println(errf)
		}
		if errfs != nil {
			fmt.Println(errfs)
		}

		// get problem string and solution string
		probStr = retProb(num1, num2, num3)
		solStr = solProb(num1, num2, num3)

		_, errf2 := f.WriteString(probStr)
		_, errfs2 := fs.WriteString(solStr)

		// check errors
		if errf2 != nil {
			fmt.Println(errf)
		}
		if errfs2 != nil {
			fmt.Println(errfs)
		}

		// finish problem and solution strings
		_, errf3 := f.WriteString("</p></div>\n")
		_, errfs3 := fs.WriteString("</p></div>\n")

		// check errors
		if errf3 != nil {
			fmt.Println(errf3)
		}
		if errfs3 != nil {
			fmt.Println(errfs3)
		}
	}

	// close files
	f.Close()
	fs.Close()
}

func retProb(num1 int, num2 int, num3 int) string {
	return fmt.Sprintf("%dx + %d = %d", num1, num2, num3)
}

func solProb(num1 int, num2 int, num3 int) string {
	numerator := num3 - num2
	if numerator != 0 {
		return fmt.Sprintf("<sup>%d</sup>/<sub>%d</sub>", numerator, num1)
	}
	for j := 1; j < 6; j++ {
		if numerator % j == 0 && num1 % j == 0 {
			return fmt.Sprintf("<sup>%d</sup>/<sub>%d</sub>", numerator / j, num1 / j)
		}
	}
	return "0"
}

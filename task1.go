package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string
	var numSubjects int

	fmt.Print("Enter student's name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&numSubjects)

	subjects := make([]string, numSubjects)
	grades := make([]float64, numSubjects)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter name of subject %d: ", i+1)
		fmt.Scanln(&subjects[i])

		var grade float64
		for {
			fmt.Printf("Enter grade for %s: ", subjects[i])
			var gradeStr string
			fmt.Scanln(&gradeStr)
			var err error
			grade, err = strconv.ParseFloat(gradeStr, 64)
			if err != nil || grade < 0 || grade > 100 {
				fmt.Println("Invalid grade. Please enter a number between 0 and 100.")
			} else {
				break
			}
		}
		grades[i] = grade
	}

	average := calculateAverage(grades)

	fmt.Printf("\nStudent Name: %s\n", name)
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("%s: %.2f\n", subjects[i], grades[i])
	}
	fmt.Printf("Average Grade: %.2f\n", average)
}

func calculateAverage(grades []float64) float64 {
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

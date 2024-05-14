package main

import (
	"fmt"
	"os"
	"path/filepath"

	mathskills "mathskills/programfiles"
)

func main() {
	// check length of arguments
	if len(os.Args) != 2 {
		fmt.Println("usage go run . data.txt")
		return
	}
	filePath := os.Args[1]
	fileinfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if fileinfo.Size() == 0 {
		fmt.Println("File is empty")
		return
	}
	if filepath.Ext(filePath) != ".txt" {
		fmt.Println("Wrong file extension use: go run . data.txt")
		return
	}
	if fileinfo.Name() != "data.txt" {
		fmt.Println("Wrong file name , use data.txt")
		return
	}

	mathskills.DataGenerate("data.txt", 100)
	
	data, err := mathskills.OpenFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	// average , kept as a float for accurate calculations
	average := mathskills.CalcAverage(data)

	// calculate median and round it safely checking for overflows
	median, err := mathskills.SafeRound(mathskills.CalcMedian(data))
	if err != nil {
		fmt.Printf("Error calculating median: %v\n", err)
		return
	}

	// calculate variance and keep it as float for accurate calculations
	variance := mathskills.CalcVariance(data, average)

	// claculate median and round it safely
	standarddeviation, err := mathskills.SafeRound(mathskills.CalcStdDev(variance))
	if err != nil {
		fmt.Printf("Error calculating standard deviation: %v\n", err)
		return
	}

	// round the average safely
	roundedAverage, err := mathskills.SafeRound(average)
	if err != nil {
		fmt.Printf("Error calculating average: %v\n", err)
		return
	}
	// round the variance safely
	roundedVariance, err := mathskills.SafeRound(variance)
	if err != nil {
		fmt.Printf("Error calculating variance: %v\n", err)
		return
	}

	// Print all values to the nearest integer
	fmt.Printf("Average: %d\n", roundedAverage)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", roundedVariance)
	fmt.Printf("Standard Deviation: %d\n", standarddeviation)
}

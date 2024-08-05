package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Read data from file and return slice of floats
func readData(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

// Calculate statistics for linear regression and Pearson correlation
func calculateStatistics(x []float64, y []float64) (float64, float64, float64) {
	n := float64(len(x))

	// Calculate sums
	var sumX, sumY, sumXx, sumYy, sumXY float64
	for i := range x {
		xi, yi := x[i], y[i]
		sumX += xi
		sumY += yi
		sumXx += xi * xi
		sumYy += yi * yi
		sumXY += xi * yi
	}

	// Linear Regression Line calculation
	m := (n*sumXY - sumX*sumY) / (n*sumXx - sumX*sumX)
	c := (sumY - m*sumX) / n

	// Pearson Correlation Coefficient calculation
	numerator := n*sumXY - sumX*sumY
	denominator := math.Sqrt((n*sumXx - sumX*sumX) * (n*sumYy - sumY*sumY))
	r := numerator / denominator

	return m, c, r
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run your-program.go <datafile>")
		return
	}

	filePath := os.Args[1]

	// Read data
	data, err := readData(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Generate x values
	x := make([]float64, len(data))
	for i := range x {
		x[i] = float64(i)
	}
	y := data

	// Calculate statistics
	m, c, r := calculateStatistics(x, y)

	// Print results
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, c)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

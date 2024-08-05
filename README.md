# click Ctrl + Shift + V to Preview
## Linear-Stats

This Go program calculates the Linear Regression Line and the Pearson Correlation Coefficient for a given dataset. The dataset is read from a file where each line represents a y-value, and x-values are the indices of these lines.

## Overview

- Linear Regression Line: Computes the best-fit line y = mx + c, where m is the slope and c is the y-intercept.
- Pearson Correlation Coefficient: Measures the strength and direction of the linear relationship between two variables, with values ranging from -1 to 1.

## Features

- Reads data from a file where each line is a y-value.
- Computes and prints the Linear Regression Line with a slope and intercept.
- Computes and prints the Pearson Correlation Coefficient.
- Outputs results with precise formatting.


## Installation

1. Clone the Repository:

``` bash
   git clone https://learn.reboot01.com/git/yrahma/linear-stats.git
   cd linear-regression-calculator 
```

2. Build the Program:

   Navigate to the directory containing main.go and run:
``` go
   go build -o linear-regression-calculator
```
   This will produce an executable named linear-regression-calculator.

## Usage

1. Prepare Your Data File:

   Create a file (e.g., data.txt) with numerical data, each number on a new line:

   189
   113
   121
   114
   145
   110

2. Run the Program:

   Use the following command to execute the program:

   ./linear-regression-calculator data.txt

   Replace data.txt with the path to your data file.

3. Example Output:

   Linear Regression Line: y = 0.543210x + 123.456789
   Pearson Correlation Coefficient: 0.9876543210

Code Explanation

- Data Reading: readData(filePath string) ([]float64, error) reads the y-values from the file and returns them as a slice of float64.
- Statistics Calculation: calculateStatistics(x []float64, y []float64) (float64, float64, float64) computes the slope, intercept, and Pearson correlation coefficient using the provided formulas.
- Main Function: Handles command-line arguments, reads data, calculates statistics, and prints results.



Contact

For any questions or feedback, please contact:

- Name:  Yusuf Rahma
- GitHub: https://github.com/Byqb

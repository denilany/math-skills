# Statistics Calculator

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

This program is written in Go to perform statistical calculations such as, (this data represents a statistical population):
- Mean
- Median
- Variance
- Standard Deviation

## Prerequisites
For this program to run, a Go environment need to be installed first.
This can be downloaded [here](https://go.dev/dl/).

## Instalation

Clone the repo in your local machine

```bash
 $ git clone https://github.com/denilany/math-skills.git
 ```

Change the directory to `math-skills` by running this:

```bash
 $ cd math-skills
 ```

## Usage

Run the program

```go
 $ go run . data.txt 
 /*
 data.txt is the file that contains data (numbers) for the program to compute the statistics
 */
```

## File data
Since the statistics is calculated from the data.txt file present, it it important to note that each individual line should only hold a single set of number/numbers with no spaces in between them. 

## Challenges
 
Some of the challenges faced was how to handle errors such as:
- If the file contains a mixture of int, int64, float32, and float64 values
- If the file contains non-numeric data
- If the values in the file are too big to be handled

## Credits

Created by:
- [Denil Anyonyi](https://github.com/denilany)


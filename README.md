# Statistics Calculator

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)

This program is written in Go to perform statistical calculations such as, (this data represents a statistical population):
- Mean
- Median
- Variance
- Standard Deviation

## Instalation

Clone the repo in your local machine

    $ git clone https://learn.zone01kisumu.ke/git/danyonyi/math-skills.git

cd into math-skills

    $ cd math-skills

## Usage

Run the program

    $ go run . data.txt 
    // data.txt is the file that contains data (numbers) for the program to compute the statistics

## Challenges
 
Some of the challenges faced was how to handle errors such as:
- If the file contains a mixture of int, int64, float32, and float64 values
- If the file contains non-numeric data
- If the values in the file are too big to be handled

## Credits

Created by:
- [Denil Anyonyi](https://github.com/denilany)


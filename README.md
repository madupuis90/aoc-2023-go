# AOC 2023 in Go

## Project Structure
The code for each day is contained in their own folders /day1, /day2. /day3...etc.

Each folder have the the code for the solutions, input and sample files and a `main_test.go` file that tests the solutions.

To run all tests for a specific day, use this command:

`go test ./dayX -v` 

To run a specific test, you can use: `go test ./dayX -run TestSamplePart1 -v`
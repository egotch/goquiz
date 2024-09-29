# Go Quiz Game

a simple quiz game that takes in a list of questions from a csv file
questions are read in then printed out one at a time
if user does not answer within the alotted time limit, the game exits

## Usage

* Create a CSV file that's 2 columns Question,Answer
* call `goquiz --csv file_name_with_full_path --time quiz-duration-in-seconds`
* program will iterate over each question in in the file and ask for input on each
* after the last question, the users's score is printed back to console

## Example

```text

╰⎯ ./goquiz -csv test/testData.csv --time 5
Question #1/13:    5+5=  10
Question #2/13:    7+3=  10
Question #3/13:    1+1=  2
Question #4/13:    8+3=  3
Question #5/13:    1+2=

TIMES UP!

Your score = 3 / 13

```

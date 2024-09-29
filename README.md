# Go Quiz Game

a simple quiz game that takes in a list of questions from a csv file
questions are read in then printed out one at a time
if user does not answer within the alotted time limit, the game exits

## Usage

* Create a CSV file that's 2 columns Question,Answer
* call `goquiz --csv file_name_with_full_path`
* program will iterate over each question in in the file and ask for input on each
* after the last question, the users's score is printed back to console

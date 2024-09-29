package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/egotch/goquiz/helpers"
  "github.com/egotch/goquiz/models"
)

func main()  {

  var score int = 0

  csvPtr := flag.String("csv", "problems.csv", "csv file name to read in format of 'question,answer'")

  flag.Parse()

  file, err := os.Open(*csvPtr)
  if err != nil {
    helpers.Exit(fmt.Sprintf("Error opening file: %s\n", *csvPtr))
  }

  _r := csv.NewReader(file)
  lines, err := _r.ReadAll()
  if err != nil {
    helpers.Exit(fmt.Sprintf("Failed to parse csv file: %s", *csvPtr))
  }


  problems := models.ParseLines(lines)

  for i, p := range problems {
    var answerIn string
    fmt.Printf("Question #%d/%d:    %s=  ",i+1, len(problems), p.Q)

    // ask for user input
    fmt.Scanf("%s\n", &answerIn)
    if err != nil {
      panic(err)
    }

    // compare answer
    if answerIn == p.A {
      score += 1
    }
  }

  fmt.Println()
  fmt.Printf("Your score = %d / %d\n", score, len(lines))
  fmt.Println()

}


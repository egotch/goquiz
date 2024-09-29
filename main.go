package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/egotch/goquiz/helpers"
	"github.com/egotch/goquiz/models"
)

func main()  {

  // get arg flags from cmd input
  timePtr := flag.Int("time", 30, "allowed duration of quiz in seconds")
  csvPtr := flag.String("csv", "problems.csv", "csv file name to read in format of 'question,answer'")
  flag.Parse()

  // declare vars
  var score int = 0
  lines := helpers.GetFile(*&csvPtr)
  timer := time.NewTimer(time.Duration(*timePtr) * time.Second)
  answerCh := make(chan string)
  

  problems := models.ParseLines(lines)

  for i, p := range problems {

    go models.GetAnswer(answerCh, problems, &p, i)

    select {

    case <- timer.C:
      helpers.Exit(fmt.Sprintf("\n\nTIME'S UP!\n\nYour score = %d / %d\n", score, len(lines)))

    case answerIn := <- answerCh:
      models.IncScore(&score, answerIn, &p)
    }
  }

  defer fmt.Printf("Your score = %d / %d\n", score, len(lines))

}


package models

import (
  "strings"
  "fmt"
)

type Problem struct {
  Q string
  A string
}

// ParseLines: parses a 2D slice of strings into
// a 2D slice of Problem stucts
func ParseLines(lines [][]string) []Problem {
  ret := make([]Problem, len(lines))

  for i, v := range lines {
    ret[i] = Problem {
      Q: v[0],
      A: strings.TrimSpace(v[1]),
    }

  }

  return ret
}

// IncScore: increments the user's score if the
// provided answer matches the expected answer
func IncScore(score *int, answerIn string, problem *Problem) {

  if answerIn == problem.A {
    *score += 1
  }
  
}

func GetAnswer(answerCh chan<- string, problems []Problem, p *Problem, i int){
  
  var answerIn string

  fmt.Printf("Question #%d/%d:    %s=  ",i+1, len(problems), p.Q)

  // ask for user input
  fmt.Scanf("%s\n", &answerIn)
  answerCh <- answerIn
}

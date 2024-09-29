package models

import "strings"

type Problem struct {
  Q string
  A string
}

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

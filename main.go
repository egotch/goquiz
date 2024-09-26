package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"github.com/egotch/goquiz/helpers"
)

func main()  {
 
  csvPtr := flag.String("csv", "default.csv", "csv file name to read in format of 'question,answer'")

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

  fmt.Println(lines)
}


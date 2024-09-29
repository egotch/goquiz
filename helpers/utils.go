package helpers

import (
  "fmt"
  "os"
  "encoding/csv"
)

// gracefully exits program with provided message
func Exit(msg string) {
  fmt.Println(msg)
  os.Exit(1)
}

// GetFile: takes in string pointer of file location
// opens the file and returns a 2D slice of the contents
func GetFile(filePtr *string) [][]string {

  file, err := os.Open(*filePtr)
  if err != nil {
    Exit(fmt.Sprintf("Error opening file: %s\n", *filePtr))
  }

  _r := csv.NewReader(file)
  lines, err := _r.ReadAll()
  if err != nil {
    Exit(fmt.Sprintf("Failed to parse csv file: %s", *filePtr))
  }

  return lines
}

package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strconv"
  "strings"
)

func main() {

  //Open file
  file, err := os.Open("input2.txt")    
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  defer file.Close()


  scanner := bufio.NewScanner(file)

  var sum int
  for scanner.Scan() {
    line := scanner.Text()
    replacedNumbersLine := ReplaceWordsToInts(line)
    sum += GetNumberPair(replacedNumbersLine)
  }

  fmt.Println(sum)
  
  // check for errors
  if err := scanner.Err(); err != nil {
    fmt.Println("Error: ", err)
  }
}


// extract first and last number in string
func GetNumberPair(input string) int {
  //regular expression to match all single digit integers
  re := regexp.MustCompile(`\d`)

  //find all matches
  matches := re.FindAllString(input, -1)

  if(len(matches) == 0) {
    return 0
  }

  firstNum := matches[0]
  lastNum := matches[len(matches) - 1]
  targetNum, err := strconv.Atoi(firstNum + lastNum)
  // fmt.Println(targetNum)
  if err != nil {
    return 0 
  }

  return targetNum
}


func ReplaceWordsToInts(input string) string {

  wordMap := map[string]string{
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
  } 

  // add last character to word ocurance 
  // this is done for the case where words overlap, like: oneight
  // which needs to be replaced with 1 and 8
  sanitizedInput := input
  for key, _ := range wordMap {
    what := key
    with := key + key[len(key) - 1:]
    sanitizedInput = strings.Replace(sanitizedInput, what, with, -1)
  }
  // fmt.Println(sanitizedInput) 
  replacedNumbersInput := sanitizedInput
  for key, val := range wordMap {
    replacedNumbersInput = strings.Replace(replacedNumbersInput, key, val, -1)
  }
 
  // fmt.Println(input, replacedNumbersInput)

  return replacedNumbersInput
}

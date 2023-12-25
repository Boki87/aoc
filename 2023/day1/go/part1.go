package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strconv"
)

func main() {


  //Open file
  file, err := os.Open("input.txt")    
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  defer file.Close()


  scanner := bufio.NewScanner(file)

  var sum int
  for scanner.Scan() {
    line := scanner.Text()
    sum += GetNumberPair(line)
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
  fmt.Println(targetNum)
  if err != nil {
    return 0 
  }

  return targetNum
}

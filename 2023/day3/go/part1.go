package main

import (
  "fmt"
  "os"
  "bufio"
  "unicode"
  "strconv"
)


func main() {

  lines, err := readLinesFromFile("input.txt")
  if err != nil {
    fmt.Println("Error reading file: ", err)
  }


  var sum int
  for lineIndex, line := range lines {
    numbers := returnAdjecentNumbers(lineIndex, line, lines)
    for _, number := range numbers {
        sum += number
    }
  }

  fmt.Println(sum)

}

func readLinesFromFile(pathToFile string) ([]string, error) {

  var lines []string
  //read the input file
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error: ", err)
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file) 

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    return nil, err 
  }

  return lines, nil
}

func returnAdjecentNumbers(lineIndex int, line string, lines []string) []int {
  var numbers []int
  for charIndex, char := range line {

    if !isNaN(string(char)) && char != '.' {
      adjecentNumbers := findAdjecentNumbers(lineIndex, charIndex, lines)
      numbers = append(numbers, adjecentNumbers...)
    }
  }

  return numbers
}

func isNaN(char string) bool {
  r := []rune(char)
  return unicode.IsDigit(r[0])
}

func findAdjecentNumbers(lineIndex int, symbolIndex int, lines []string) []int  {

  var numbers []int

  // numbers on the line obove the symbol     
  if lineIndex > 0 {
    foundNumbersAbove := findNumbersOnLine(lines[lineIndex - 1], symbolIndex)
    numbers = append(numbers, foundNumbersAbove...)
  }

  // numbers on the same line as symbol

  foundNumbersOnSameLine := extractNumbers(lines[lineIndex]) 
  for _, numOnLine := range foundNumbersOnSameLine {
    // check if symbol is not on the edge
    if symbolIndex > 0 && symbolIndex < len(lines[lineIndex]) - 1 {
      if(numOnLine.StartIndex == symbolIndex + 1 || numOnLine.EndIndex == symbolIndex - 1) {
        numOnLineValue, _ := strconv.Atoi(numOnLine.Value)
        numbers = append(numbers, numOnLineValue)
      }
    }

    // if symbol is first character on line
    if symbolIndex == 0 {
      if(numOnLine.StartIndex == symbolIndex + 1) {
        numOnLineValue, _ := strconv.Atoi(numOnLine.Value)
        numbers = append(numbers, numOnLineValue)
      }       
    }

    // if symbol is last character on line
    if symbolIndex == len(lines[lineIndex]) - 1 {
      if(numOnLine.EndIndex == symbolIndex - 1) {
        numOnLineValue, _ := strconv.Atoi(numOnLine.Value)
        numbers = append(numbers, numOnLineValue)
      }       
    }
  }


  // numbers on the line under the symbol

  if lineIndex < len(lines) - 1 {
    foundNumbersAbove := findNumbersOnLine(lines[lineIndex + 1], symbolIndex)
    numbers = append(numbers, foundNumbersAbove...)
  }



  return numbers 
}

func findNumbersOnLine(line string, symbolIndex int) []int {

  var adjecentNumbers []int
  numbers := extractNumbers(line)

  // loop thorugh numbers and check if they are adjecent to symbolIndex

  for _, number := range numbers {
    if (number.StartIndex >= symbolIndex - 1 && number.StartIndex <= symbolIndex + 1) || (number.EndIndex >= symbolIndex - 1 && number.EndIndex <= symbolIndex + 1)  {
      numberInt, _ := strconv.Atoi(number.Value)

      adjecentNumbers = append(adjecentNumbers, numberInt)
    }
  }

  return adjecentNumbers 
}


type NumberInfo struct {
  Value      string
  StartIndex int
  EndIndex   int
}

func extractNumbers(input string) []NumberInfo {
  var numbers []NumberInfo
  var currentNumber string
  var startIndex, endIndex int
  isInNumber := false

  for i, char := range input {
    if unicode.IsDigit(char) {
      if !isInNumber {
        isInNumber = true
        startIndex = i
      }
      currentNumber += string(char)
    } else {
      if isInNumber {
        isInNumber = false
        endIndex = i - 1 // The index of the last digit
        numbers = append(numbers, NumberInfo{Value: currentNumber, StartIndex: startIndex, EndIndex: endIndex})
        currentNumber = ""
      }
    }
  }

  // Check if the last character is part of a number
  if isInNumber {
    numbers = append(numbers, NumberInfo{Value: currentNumber, StartIndex: startIndex, EndIndex: len(input) - 1})
  }

  return numbers
}




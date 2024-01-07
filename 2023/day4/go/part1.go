package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)



func main() {
  cards, err := readLinesFromField("input.txt")
  if err != nil {
    fmt.Println("Error reading input file: ", err)
  }

  var points int

  for _, card := range cards {
    pointsForCard, err := getPointsForCard(card)
    if err != nil {
      fmt.Println(err)
      return
    }
    points += pointsForCard
  }
 
  fmt.Println(points)

}



func readLinesFromField(pathToFile string) ([]string, error) {
  file, err := os.Open(pathToFile) 
  if err != nil {
    fmt.Println("Error: ", err)
    return nil, err 
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)


  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    return nil, err
  }

  return lines, nil

} 


func toMap(numbersArr []int) map[int]bool {
  var numMap = make(map[int]bool)      
  for _, number := range numbersArr {
    numMap[number] = true 
  } 
  return numMap
}


func getPointsForCard(card string) (int, error) {

  // Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
  cardSplit := strings.Split(card, ": ")    
  _, numbers := cardSplit[0], cardSplit[1]

  numbersSplit := strings.Split(numbers, " | ") 
  winningNumbersString, drawnNumbersString := numbersSplit[0], numbersSplit[1]

  drawnNumbersStrings := strings.Split(drawnNumbersString, " ")
  drawnNumbersInts := convertStringsToInts(drawnNumbersStrings)

  winningNumsStrings := strings.Split(winningNumbersString, " ")
  winningNums := convertStringsToInts(winningNumsStrings)  
  

  winningNumsDict := toMap(winningNums) 

  var points int

  for _, num := range drawnNumbersInts {
    if winningNumsDict[num] {
      if points == 0 {
          points = 1
      } else {
          points *= 2
      }
    }
  }
   
  
  return points, nil
}

func convertStringsToInts(stringArr []string) []int {
    
    var resArr []int 

    for _, numStr := range stringArr {
        numInt, err := strconv.Atoi(numStr)
        if err != nil {
          continue
        }
        resArr = append(resArr, numInt)
    }
    return resArr
}


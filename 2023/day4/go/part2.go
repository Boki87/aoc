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


  var toProcess []int
  for i, _ := range cards {
    toProcess = append(toProcess, i + 1)
  } 


  var count = make(map[int]int)

  for len(toProcess) > 0 {
    idx := toProcess[len(toProcess) - 1]
    toProcess = toProcess[:len(toProcess) - 1]

    if _, exists := count[idx]; exists {
        count[idx] = count[idx] + 1
    } else {
        count[idx] = 1
    }
    
    points, _ := getPointsForCard(cards[idx - 1], idx)
    
    toProcess = append(toProcess, points...)
    // fmt.Println(idx, points)
  }

      
    var sum int
    for _, val := range count {
       sum += val 
    }

    fmt.Println(sum)
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


func getPointsForCard(card string, idx int) ([]int, error) {

  // example card row in input text file
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
      points++
    }
  }
  
  
  result := make([]int, points) 
  for i := 0; i < points; i ++ {
    result[i] = idx + 1 + i     
  } 
  
  return result, nil
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


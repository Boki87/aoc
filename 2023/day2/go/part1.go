package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
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
    game := scanner.Text() 
    gameIndex := ReturnSuccessfullGameIndex(game)
    sum += gameIndex
  }
  fmt.Println(sum)

  if err := scanner.Err(); err != nil {
      fmt.Println("Error: ", err)
  }
}


func ReturnSuccessfullGameIndex(input string) int {

  maxColors := map[string]int{
    "red": 12,
    "green": 13,
    "blue": 14,
  }


  result := strings.Split(input, ":")
  gameTitle, gameScores := result[0], result[1]
  gameIndex := strings.Split(gameTitle, " ")[1]

  individualGames := strings.Split(gameScores, "; ")         
 
  for _, game := range individualGames {
    cubeRes := strings.Split(strings.TrimSpace(game), ", ")
    
    for _, cube := range cubeRes {
      cubeSplit := strings.Split(cube, " ")
      amount, _ := strconv.Atoi(cubeSplit[0])
      color := cubeSplit[1]
      if maxColors[color] < amount {
        return 0
      }
    } 
  }

  index, _ := strconv.Atoi(gameIndex) 
  return index
}  







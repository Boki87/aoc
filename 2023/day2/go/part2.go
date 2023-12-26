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
    sum += MinimumPowerOfCubes(game)
  }
  
  fmt.Println(sum)
  if err := scanner.Err(); err != nil {
      fmt.Println("Error: ", err)
  }
}

func MinimumPowerOfCubes(input string) int {

  wholeGame := strings.Split(input, ":")   
  // gameIndex := strings.Split(wholeGame[0], " ")[1]
  gameScores := strings.Split(strings.TrimSpace(wholeGame[1]), "; ") 

  maxColor := make(map[string]int)
  for _, game := range gameScores {

    individualGame := strings.Split(game, ", ")


    for _, cube := range individualGame {
      gameSplit := strings.Split(cube, " ")
      amount, _ := strconv.Atoi(gameSplit[0])
      color := gameSplit[1]

      if maxColor[color] < amount {
        maxColor[color] = amount 
      }

    }

  }


  sum := 1
  for _, val := range maxColor {
     sum *= val 
  }
  return sum
}



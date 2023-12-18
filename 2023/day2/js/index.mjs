import fs from 'fs'

/* part 1
function run() {
  const lines = fs.readFileSync('./input.txt', 'utf8').trim().split('\n')

  const maxColors = {
    red: 12,
    green: 13,
    blue: 14
  }

  let sum = 0

  lines.forEach(line => {
    const [gameTitle, gameScores] = line.split(":")
    const gameId = gameTitle.split(" ")[1]

    const isPossible = gameScores.split(';').map(game => {
      const scores = game.trim().split(', ').map(score => {
        const [amount, color] = score.split(' ')
        return maxColors[color] >= amount
      })

      return scores.every(score => score)
    }).every(game => game)


    if (isPossible) {
      sum += parseInt(gameId)
    }
  })

  console.log(sum)
}
*/


// part 2

const testInput = [
  'Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green',
  'Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue',
  'Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red',
  'Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red',
  'Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green',
]
function run() {
  const lines = fs.readFileSync('./input.txt', 'utf8').trim().split('\n')
  //const lines = testInput
  let res = 0

  lines.forEach(line => {
    const [gameTitle, scores] = line.split(': ')

    const games = scores.split('; ')

    const maxColors = {
      red: 0,
      green: 0,
      blue: 0
    }
    let power = 1

    games.forEach(game => {
      const cubes = game.split(', ')
      cubes.forEach(color => {
        const [amount, colorName] = color.split(' ')
        //console.log('maxColors', maxColors)
        if (maxColors[colorName] < +amount) {
          maxColors[colorName] = +amount
        }
      })
    })

    //console.log(gameTitle, maxColors)

    for (let prop in maxColors) {
      power *= maxColors[prop]
    }

    res += power
  })

  console.log(res)
}

run()

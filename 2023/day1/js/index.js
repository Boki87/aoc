const fs = require("fs");
const input = fs.readFileSync("input.txt", "utf8").trim().split("\n");



const numbers_arr = []

// loop through each line
for (const line of input) {
  //console.log(line)


  const numbersMap = {
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9
  }


  const regexPatern = new RegExp(Object.keys(numbersMap).join('|'), 'gi')
  //const regexPatern = new RegExp(`/?=(${Object.keys(numbersMap).join('|')})/`, 'gi')


  const newLine = line.replace(regexPatern, (match) => {
    return match + match[match.length - 1]
  })
    .replace(regexPatern, (match) => {
      return numbersMap[match.toLowerCase()] || match
    })


  const allNumbers = newLine.match(/\d+/g).map(num => {
    if (num.length > 1) {
      return num.split('')
    }
    return num
  }).flat()


  if (allNumbers.length === 0) {
    continue
  }

  // get first and last number
  const firstNumber = allNumbers[0]
  const lastNumber = allNumbers[allNumbers.length - 1]
  const number = parseInt(`${firstNumber}${lastNumber}`)

  numbers_arr.push(number)
}

const res = numbers_arr.reduce((acc, x) => acc + x, 0)

console.log(res)


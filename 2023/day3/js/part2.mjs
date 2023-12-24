import fs from "fs";

const lines = fs.readFileSync("./input2.txt", "utf8").trim().split("\n");

function getNumbersOnLine(line) {
  const numbers = [];
  const regex = /\d+/g;
  let match;
  while ((match = regex.exec(line)) != null) {
    numbers.push(match);
  }
  return numbers;
}

function isSymbol(char) {
  return isNaN(char) && char !== ".";
}

function isIntersecting(match, symbolIndex) {
  const numStart = match.index;
  const numEnd = match.index + match[0].length - 1;
  // console.log(match, numEnd, symbolIndex);
  if (
    (numEnd >= symbolIndex - 1 && numEnd <= symbolIndex + 1) ||
    (numStart >= symbolIndex - 1 && numStart <= symbolIndex + 1)
  ) {
    return true;
  }
  return false;
}

function getAllNumbersIntersecting(symbolLineIndex, symbolIndex) {
  const allNumbers = [];

  //above
  if (symbolLineIndex > 0) {
    const targetLine = lines[symbolLineIndex - 1];
    const numbers = getNumbersOnLine(targetLine);
    numbers.forEach((num) => {
      if (isIntersecting(num, symbolIndex)) {
        allNumbers.push(num[0]);
      }
    });
  }

  //under
  if (symbolLineIndex < lines.length - 1) {
    const targetLine = lines[symbolLineIndex + 1];
    const numbers = getNumbersOnLine(targetLine);
    numbers.forEach((num) => {
      if (isIntersecting(num, symbolIndex)) {
        allNumbers.push(num[0]);
      }
    });
  }

  // next to
  const numbersOnSameLine = getNumbersOnLine(lines[symbolLineIndex]);
  numbersOnSameLine.forEach((num) => {
    const numStart = num.index;
    const numEnd = num.index + num[0].length - 1;
    if (numEnd === symbolIndex - 1) {
      allNumbers.push(num[0]);
    }

    if (numStart === symbolIndex + 1) {
      allNumbers.push(num[0]);
    }
  });

  return allNumbers;
}

function run() {
  let sum = 0;
  lines.forEach((line, lineIndex) => {
    for (let x = 0; x < line.length; x++) {
      const char = line[x];
      if (isSymbol(char)) {
        // console.log(char);
        let numbers = getAllNumbersIntersecting(lineIndex, x);
        if (numbers.length === 2) {
          sum += numbers[0] * numbers[1];
        }
      }
    }
  });
  console.log(sum);
}

run();

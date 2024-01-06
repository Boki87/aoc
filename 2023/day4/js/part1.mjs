import fs from "fs";

const lines = fs.readFileSync("input.txt", "utf8").trim().split("\n");

function run() {
  console.log(lines);

  let totalPoints = 0;
  lines.forEach((card) => {
    const [winingCardStr, drawnNumbersStr] = card.split(" | ");

    const [_, winingNumbersStr] = winingCardStr.split(": ");

    const winingNumbersArr = winingNumbersStr
      .trim()
      .replaceAll("  ", " ")
      .split(" ");

    const drawnNumbersArr = drawnNumbersStr
      .trim()
      .replaceAll("  ", " ")
      .split(" ");

    const drawnWiningNumbersArr = drawnNumbersArr.filter((num) => {
      if (winingNumbersArr.includes(num)) {
        return num;
      }
    });

    let cardPoints = 0;
    if (drawnWiningNumbersArr.length > 0) {
      cardPoints = 2 ** Math.floor(drawnWiningNumbersArr.length - 1);
    }

    totalPoints += cardPoints;
  });

  console.log(totalPoints);
}

run();

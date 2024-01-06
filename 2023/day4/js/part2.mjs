import fs from "fs";

const cards = fs.readFileSync("input.txt", "utf8").trim().split("\n");

// console.log(cards);

function toDict(numbersArr) {
  return numbersArr.reduce((acc, num) => {
    acc[num] = true;
    return acc;
  }, {});
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
function getPointsForCard(card, cardIdx) {
  const vals = card.split(":")[1];
  let [winningNums, drawnNums] = vals.split(" | ");

  winningNums = winningNums.split(" ").filter((x) => x);
  const winningNumsDict = toDict(winningNums.map((x) => parseInt(x)));
  drawnNums = drawnNums.split(" ").filter((x) => parseInt(x.trim()));

  // sum how many points we have for the given card
  // we need this to generate the cards that we duplicate after this card
  let sumPoints = 0;
  drawnNums.forEach((num) => {
    if (winningNumsDict[num]) {
      sumPoints++;
    }
  });

  return new Array(sumPoints).fill(cardIdx + 1).map((x, i) => {
    return x + i;
  });
}

// for (let [index, card] of cards.entries()) {
//   const pointsForCard = getPointsForCard(card, index + 1);
//   console.log(pointsForCard);
// }

const toProcess = new Array(cards.length).fill(0).map((_, i) => i + 1);
const seen = {};
const count = {};

while (toProcess.length) {
  const idx = toProcess.pop();
  // console.log(idx, toProcess);

  count[idx] = count[idx] ? count[idx] + 1 : 1;
  // console.log(cards[idx]);
  const points = seen[idx] ? seen[idx] : getPointsForCard(cards[idx - 1], idx);
  // console.log(idx, { points, seen, count });
  // console.log(count);
  toProcess.push(...points);
}

console.log(
  count,
  Object.keys(count).reduce((acc, x) => {
    return acc + count[x];
  }, 0),
);

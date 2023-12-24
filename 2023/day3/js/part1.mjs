import fs from "fs";

const lines = fs.readFileSync("./input.txt", "utf8").trim().split("\n");

function isSymbol(char) {
  return isNaN(char) && char !== ".";
}

function isPart(numPos, lineIndex) {
  const { start, end, number } = numPos;

  function getSafeStart(start) {
    if (start === 0) {
      return 0;
    } else {
      return start - 1;
    }
  }

  function getSafeEnd(end, line) {
    if (end == line.length - 1) {
      return line.length;
    } else {
      return end + 2;
    }
  }

  const chars = [];

  //above
  if (lineIndex > 0) {
    chars.push(
      lines[lineIndex - 1].slice(
        getSafeStart(start),
        getSafeEnd(end, lines[lineIndex - 1]),
      ),
    );
  }

  //on the same line, left and right
  const sameLine = `${lines[lineIndex].charAt(start - 1)}${lines[
    lineIndex
  ].charAt(end + 1)}`;
  chars.push(sameLine);

  //under
  if (lineIndex < lines.length - 1) {
    chars.push(
      lines[lineIndex + 1].slice(
        getSafeStart(start),
        getSafeEnd(end, lines[lineIndex + 1]),
      ),
    );
  }

  // console.log(number, chars, start, end);

  return chars
    .join("")
    .trim()
    .split("")
    .some((c) => isSymbol(c));
}

function run() {
  let sum = 0;

  lines.forEach((line, lineIndex) => {
    const numbers = [];
    const regex = /\d+/g;
    let match;
    while ((match = regex.exec(line)) != null) {
      numbers.push(match);
    }

    if (!numbers) return;

    numbers.forEach((num) => {
      if (!num) return;
      const end = num.index + num[0].length - 1;
      const position = {
        start: num.index,
        number: num[0],
        end,
      };

      let res = isPart(position, lineIndex);
      if (res) {
        sum += parseInt(num);
      }
    });
  });
  console.log(sum);
}

run();

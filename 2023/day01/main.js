const fs = require("node:fs");
const getInput = (name) => fs.readFileSync(name, "utf8").split("\n");

const part1 = getInput("input")
  .map((a) => {
    const n = a.split("");
    return Number(n.find((a) => !isNaN(a)) + n.findLast((a) => !isNaN(a)));
  })
  .filter((a) => a)
  .reduce((a, b) => a + b);

const wordToNum = (a) =>
  !isNaN(a)
    ? a
    : {
        one: 1,
        two: 2,
        three: 3,
        four: 4,
        five: 5,
        six: 6,
        seven: 7,
        eight: 8,
        nine: 9,
      }[a];

const matchAllNumbers = (str) =>
  [
    ...str.matchAll(/(?=(\d|one|two|three|four|five|six|seven|eight|nine))/g),
  ].map((a) => a[1]);

const part2 = getInput("input")
  .filter((a) => a)
  .map((line) => {
    const n = matchAllNumbers(line).map(wordToNum);
    return Number(String(n[0]).concat(n[n.length - 1]));
  })
  .reduce((a, b) => a + b);

console.log({ part1, part2 });

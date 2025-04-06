const fs = require("node:fs");

const fileStream = fs.createWriteStream("test.txt");
const start = Date.now();

for (let i = 0; i < 10_000_000; i++) {
  fileStream.write(` ${i} `);
  if (i % 1_000_000 === 0) {
    console.log(i);
  }
}

const end = Date.now();

console.log(`Time taken: ${end - start}ms`);


const fs = require("node:fs")

console.time("performance")

const readStream = fs.createReadStream("numbers.txt")

readStream.on("data", chunk => {
    console.log(chunk.toString())
})

readStream.on("end", () => {
    console.timeEnd("performance")
})

// const content = fs.readFileSync("numbers.txt")
// console.log(content.toString())

// console.timeEnd("performance")
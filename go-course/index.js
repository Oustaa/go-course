const fs = require("node:fs");

const limit = 1_000_000;

const fileStream = fs.createWriteStream("numbers.txt");

function waitForDrain() {
    return new Promise((resolve) => {
        fileStream.once("drain", resolve);
    });
}

async function writeNumbers() {
    for (let i = 0; i < limit; i++) {
        if (!fileStream.write(` ${i} `)) {
            await waitForDrain();
        }
    }

    fileStream.end();
}
writeNumbers().then(() => {
    console.log("Finished writing to file.");
}).catch((err) => {
    console.error("Error writing to file:", err);
});

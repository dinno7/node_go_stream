const { spawn } = require('child_process');
const fs = require('fs');
const { argv, exit } = require('process');
const { pipeline } = require('stream/promises');

if (argv.length < 3) {
	console.error('⭕️ ~ ERROR  ~ Please provide valid input file path');
	exit(1);
}

// > Configuration of normalizer app
const numberFormatter = spawn('./normalizer', ['./dest.txt', '$', '.']);
numberFormatter.stdout.on('data', (chunk) => {
	console.log(`💀 stdout ${chunk}`);
});
numberFormatter.stderr.on('data', (chunk) => {
	console.log(`💀 stderr ${chunk}`);
});
numberFormatter.on('close', (code) => {
	if (code === 0) {
		console.log(`💀 From NodeJs - The process was successful`);
	} else {
		console.log(`💀 From NodeJs - Something bad wrong`, code);
	}
});

// > Reading source file and pipe it to normalizer app(go-lang)
const filePath = argv[2];
const readStream = fs.createReadStream(filePath);
async function run() {
	await pipeline(readStream, numberFormatter.stdin);
	console.log(`💀 From NodeJs - Pipeline succeeded.`);
}

run().catch(console.error);

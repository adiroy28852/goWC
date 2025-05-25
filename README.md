# Go WC

A simple implementation of the Unix `wc` (word count) utility in Go.  
This project was created as a learning exercise to get familiar with the Go programming language.

## Features

- Counts lines, words, and characters in files or standard input.
- Supports flags for line (`-l`), word (`-w`), and character (`-c`) counts.
- Mimics the basic behavior of the Unix `wc` command.

## Usage

### Installation

Clone the repository:

```sh
git clone https://github.com/yourusername/go_wc.git
cd go_wc
```

### Run

To count words in a file:

```sh
go run main.go -w testingWC.txt
```

To count lines:

```sh
go run main.go -l testingWC.txt
```

To count characters:

```sh
go run main.go -c testingWC.txt
```

You can combine flags as needed. If no flags are provided, all counts (lines, words, and characters) are shown:

```sh
go run main.go testingWC.txt
```

To use with standard input:

```sh
go run main.go

${enter your text here. You can Paste from ClipBoard, or type here. Once done, exit via Ctrl + C. Results show up automatically}
```

## Requirements

- Go 1.18 or newer

## License

MIT License

# Go WC

A simple implementation of the Unix `wc` (word count) utility in Go.  
This project was created as a learning exercise to get familiar with the Go programming language.

## Features

- Counts lines, words, and bytes in files or standard input.
- Mimics the basic behavior of the Unix `wc` command.

## Usage

### Installation

Clone the repository:

```sh
git clone https://github.com/yourusername/go_wc.git
cd go_wc
```


### Build

```sh
go build -o go_wc
```

### Run

To count lines, words, and bytes in a file:

```sh
./go_wc filename.txt
```

To use with standard input:

```sh
echo "Hello, world!" | ./go_wc
```

You can also pass multiple files:

```sh
./go_wc file1.txt file2.txt
```

## Requirements

- Go 1.18 or newer

## License

MIT License

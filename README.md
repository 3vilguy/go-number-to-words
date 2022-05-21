# Number to words (in Golang)

A CLI tool that takes a number, from 0 to 100000, and converts the number into grammatically correct English words.

## Running Your Application

There is a binary file that you can execute. No dependencies should be required:

```bash
$ ./numbers-to-words [number]
```

Examples:

```bash
$ ./numbers-to-words 10
ten

$ ./numbers-to-words 123
one hundred and twenty-three
```

## Prerequisites for building

Make sure you have installed all of the following prerequisites on your development machine:

* [The Go Programming Language](https://go.dev/). I was using version 1.18, but since no fancy stuff is being used, it should work with older versions as well.


## Building

To generate an executable binary:

```bash
$ go build
```

## Testing Your Application

To run the tests:

```bash
$ go test ./convert
```

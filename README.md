# A Hangman Game API, written in Golang

An API for a CLI hangman game.

## Using the API

Clone the repo into `$GOPATH/src` and do the following:

get deps:
```bash
  # inside the repo
  $ go get 
```
build:
```bash
  # inside the repo
  $ go build
```

start the server:
```bash
  #inside the repo
  $ ./hangman-api
```

## Test

Clone the repo and do the following:

get deps:
```bash
  # inside the repo
  $ go get 
```

run the tests
```bash
  # inside the repo
  $ go test 
```
*a note on testing: before any further features are added I would like to unit test the public API of this program with enough variance in data to cover all cases. Due to time restrictions this has not been possible*

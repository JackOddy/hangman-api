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
  $ go test -v ./game
  # right now the only package with tests is game
  
```
*a note on testing: before any further features are added I would like to unit test the public API of this program with enough variance in data to cover all cases. Due to time restrictions this has not been possible*

*a note on concurrency: this project did not seem to lend itself heavily to concurrency. The server handles each request concurrently and due to the nature of the game each step inevitably needs to happen sequentially. However in terms of multiple games existing at once, this is covered by using a syncmap.Map - a threadsafe map (mux) that guards against concurrent read/writes*

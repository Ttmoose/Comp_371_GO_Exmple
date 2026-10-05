# Comp_371_GO_Exmple

A small Go program demonstrating some of the language's distinctive features.

## Run

```sh
go run .       # Find primes up to 100
go run . 10000 # Choose a limit from 2 to 1,000,000
```

The program uses goroutines and channels to distribute prime checks across a
worker pool. It also demonstrates Go's statically typed functions, built-in
concurrency support, standard library, and explicit error handling for input.
Prime results are sorted before printing so the output is consistent despite
the work finishing concurrently.

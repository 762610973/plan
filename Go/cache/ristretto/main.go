package main

import "github.com/dgraph-io/ristretto"

//! This project is for learning ristretto.

func main() {
	_, _ = ristretto.NewCache(&ristretto.Config{})
}

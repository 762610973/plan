package main

import gocache "github.com/patrickmn/go-cache"

func main() {
	_ = gocache.New(gocache.DefaultExpiration, gocache.DefaultExpiration)
}

package main

import "github.com/patrickmn/go-cache"

func main() {
	_ = cache.New(cache.DefaultExpiration, cache.DefaultExpiration)
}

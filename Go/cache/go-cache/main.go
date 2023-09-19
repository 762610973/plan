package main

import c "github.com/patrickmn/go-cache"

func main() {
	_ = c.New(c.DefaultExpiration, c.DefaultExpiration)
}

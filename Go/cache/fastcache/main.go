package main

import "github.com/VictoriaMetrics/fastcache"

func main() {
	_ = fastcache.New(10)
}

package gohealthcheck

import (
	"flag"
	"net/http"
	"os"
)

func Register(addr string) {
	hc := flag.Bool("hc", false, "perform a health check")
	flag.Parse()
	if !*hc {
		return
	}
	resp, err := http.Get(addr)
	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	}
	os.Exit(0)
}

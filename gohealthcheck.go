package gohealthcheck

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func Register(addr string) {
	hc := flag.Bool("hc", false, "perform a health check")
	flag.Parse()
	if !*hc {
		return
	}
	l := log.New(os.Stdout, "health-check ", log.LstdFlags)
	resp, err := http.Get(addr)
	if err != nil || resp.StatusCode != 200 {
		l.Println("FAIL")
		os.Exit(1)
	}
	l.Println("OK")
	os.Exit(0)
}

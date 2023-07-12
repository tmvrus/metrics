package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			list := generate(10)
			body, _ := json.Marshal(list)
			for {
				_, _ = http.Post("http://localhost:8080/save", "application/json", bytes.NewReader(body))
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
			}

		}(i + 10)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

func generate(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = randStr(n * i)
	}

	return res
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

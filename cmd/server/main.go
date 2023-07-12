package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// один хендлер со списком json
	// сохраннеи его в map

	http.HandleFunc("/save", mainHandler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(123)
	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var list []string
	err = json.Unmarshal(data, &list)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := make(map[string]string, len(list))
	for _, s := range list {
		m[hash(s)] = s
	}

	err = saveData(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func saveData(_ map[string]string) error {
	ts := time.Now().Unix()
	if rand.Int63n(10) > 5 && ts%5 == 0 {
		return fmt.Errorf("error")
	}
	return nil
}

func hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

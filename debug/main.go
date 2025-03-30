package main

import (
	"api/tts"
	"golib/server"
	"golib/server/http"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	entries, err := os.ReadDir("voices")
	if err != nil {
		log.Fatal(err)
	}
	voices := make(map[string]string, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if ext := filepath.Ext(name); ext == ".bin" {
			voices[name[:len(name)-4]] = filepath.Join("voices", name)
		}
	}
	if err = tts.Initialize(0x92, "model.bin", "lexicon_zh.txt", "lexicon_en_us.json", "lexicon_en_gb.json", voices, 4, "output.wav"); err != nil {
		log.Fatal(err)
	}
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	http.ReturnErr = true
	log.Fatal(server.ServeHTTP(":3000", http.NewServer(), 5*time.Second))
}

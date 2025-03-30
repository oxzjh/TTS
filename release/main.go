package main

import (
	"api/tts"
	"context"
	"embed"
	"fmt"
	"golib/chrome"
	"golib/server"
	"golib/server/http"
	"log"
	"os"
	"path/filepath"
)

//go:embed dist
var dist embed.FS

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
	l1, port1, err := server.ServeLocal(http.NewServer())
	if err != nil {
		log.Fatal(err)
	}
	defer l1.Close()
	apiURL := fmt.Sprintf("http://127.0.0.1:%d", port1)

	c, err := chrome.New(context.Background(), "", "", 640, 480, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.Bind("fetchURL", func() string {
		return apiURL
	})

	l2, port2, err := server.ServeLocalFile(dist)
	if err != nil {
		log.Fatal(err)
	}
	defer l2.Close()

	c.Load(fmt.Sprintf("http://127.0.0.1:%d/dist", port2))
	// c.DisableContextMenu()
	// c.DisableDevTools()
	<-c.Done()
}

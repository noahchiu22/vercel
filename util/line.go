package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"vercel/models"
)

func Response(res models.Reply, token string) {
	if len(res.Messages) > 5 {
		res.Messages = []models.Message{{
			Type: "text",
			Text: "符合條件過多\n請再試一次",
		}}
	}
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(res)

	req, err := http.NewRequest("POST", ReplyUrl, &buf)
	if err != nil {
		return
	}

	// Set the request content type
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer {%s}", token))

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	sitemap, _ := io.ReadAll(resp.Body)

	fmt.Println("收到的回傳:", string(sitemap))
}

func Push(res models.Push, token string) {
	if len(res.Messages) > 5 {
		res.Messages = []models.Message{{
			Type: "text",
			Text: "符合條件過多\n請再試一次",
		}}
	}
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(res)

	req, err := http.NewRequest("POST", PushUrl, &buf)
	if err != nil {
		return
	}

	// Set the request content type
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer {%s}", token))

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	sitemap, _ := io.ReadAll(resp.Body)

	fmt.Println("收到的回傳:", string(sitemap))
}

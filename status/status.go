package status

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func SetSlackStatusUntilEOD(emoji, text string) {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	SetSlackStatus(emoji, text, date.Unix()-1)
}

func SetSlackStatusForever(emoji, text string) {
	SetSlackStatus(emoji, text, 0)
}

func SetSlackStatus(emoji, text string, expiration int64) {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		return
	}

	emoji = strings.ReplaceAll(emoji, ":", "")

	body, err := json.Marshal(map[string]any{
		"profile": map[string]any{
			"status_text":       text,
			"status_emoji":      fmt.Sprintf(":%s:", emoji),
			"status_expiration": expiration,
		},
	})
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "https://slack.com/api/users.profile.set", bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()
}

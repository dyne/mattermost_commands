package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// The mattermost thing
type Notification struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
	GotoLocation string `json:"goto_location"`
}

// The serverino
func MiniServer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	address := getEncodedAddress(r.FormValue("text"))
	text := fmt.Sprintf(`🚨 **Hey beautiful souls!** 🚨

If you're free, we’ve got a videoconference happening today, and I’d love for you to join! 🎥💻

🗓️ **When?** Today
⏰ **Time?** Right now (or whenever you're ready, coffee ☕️ in hand)
💻 **Where?** %s

It'll be great to have you there—let's connect and make it a good one! 💫💬

– *[%s]* 😊✨
`, address, r.FormValue("user_name"))
	data := Notification{text, "in_channel", address}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(&data)
	if err != nil {
		panic(err)
	}
}

func getEncodedAddress(path string) string {
	baseURL, err := url.Parse("https://vdc.dyne.org")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
	}

	baseURL.Path += path
	return baseURL.String()
}

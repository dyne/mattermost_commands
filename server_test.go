package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGETVDC(t *testing.T) {
	cases := []struct {
		got  string
		want string
	}{
		{"", "https://vdc.dyne.org"},
		{"12", "https://vdc.dyne.org/12"},
		{"mimmo", "https://vdc.dyne.org/mimmo"},
		{"a/b", "https://vdc.dyne.org/a/b"},
		{"a?b=1", "https://vdc.dyne.org/a%3Fb=1"},
		{"⛅", "https://vdc.dyne.org/%E2%9B%85"},
		{"me space", "https://vdc.dyne.org/me%20space"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("return a correct link for %q is %q", test.got, test.want), func(t *testing.T) {
			form := url.Values{}
			form.Add("channel_id", "fukxanjgjbnp7ng383at53k1sy")
			form.Add("channel_name", "town-square")
			form.Add("command", "/weather")
			form.Add("response_url", "http://localhost:8066/hooks/commands/i11f6nnfgfyk8eg56x9omc6dpa")
			form.Add("team_domain", "team-awesome")
			form.Add("team_id", "wx4zz8t4ttgmtxqiwfohijayzc")
			form.Add("text", test.got)
			form.Add("token", "qzgakf1nx3yt9dr4n8585ihbxy")
			form.Add("trigger_id", "ZWZ5ZjRndzR4YmJxOHJlZWh4MXpkaHozbnI6ZXJqNnFjazNyZmd0dWpzODZ3NXI2cmNremg6MTY2MjA0MTY5Njg5NjpNRVFDSUQ5cTZ3MkRHU1RaNjhyaDh1TGl1STlSVHh2R1czSXZ5aGVRYjhkWThuZnlBaUI2YnlPR2ZpWlczR1FmVkdIODlreEp4MmlVT0UxMm9LMjlkZ1d0RC8xbjZRPT0=")
			form.Add("user_id", "erj6qck3rfgtujs86w5r6rckzh")
			form.Add("user_name", "alan")

			request, _ := http.NewRequest(http.MethodPost, test.got, bytes.NewBufferString(form.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			request.Header.Set("Accept", "application/json")
			request.Header.Set("Accept-Encoding", "gzip")
			request.Header.Set("Authorization", "Token qzgakf1nx3yt9dr4n8585ihbxy")
			request.Header.Set("User-Agent", "Mattermost-Bot/1.1")
			response := httptest.NewRecorder()

			MiniServer(response, request)
			var notification Notification
			err := json.NewDecoder(response.Body).Decode(&notification)

			if err != nil {
				t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
			}

			got := notification.GotoLocation

			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
			assertStatus(t, response.Code, http.StatusOK)
		})
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

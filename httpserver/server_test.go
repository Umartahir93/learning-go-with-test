package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubPlayerStore struct {
	scores map[string]int
}

func (s *stubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGetPlayer(t *testing.T) {
	store := stubPlayerStore{
		map[string]int{
			"umar":  20,
			"robin": 10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("return umar's score", func(t *testing.T) {
		request := getScoreRequest("umar")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)

	})

	t.Run("return robin's score", func(t *testing.T) {
		request := getScoreRequest("robin")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func getScoreRequest(name string) *http.Request {

	endpoint := fmt.Sprintf("/players/%s", name)
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)

	return req
}

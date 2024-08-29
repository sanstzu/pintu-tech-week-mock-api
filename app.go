package main

import (
	"math/rand/v2"
	"net/http"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func main() {
	router := http.NewServeMux()

	quotes := []string{
		"Try not to become a person of success, but rather try to become a person of value. ~Albert Einstein",
		"When you stop chasing the wrong things you give the right things a chance to catch you. ~Lolly Daskal",
		"Don't be afraid to give up the good to go for the great. ~John D. Rockefeller",
		"Life is not about finding yourself. Life is about creating yourself. ~Lolly Daskal",
		"Only put off until tomorrow what you are willing to die having left undone. ~Pablo Picasso",
		"Success is liking yourself, liking what you do, and liking how you do it. ~Maya Angelou",
		"To be successful you must accept all challenges that come your way. You can't just accept the ones you like. ~Mike Gafka",
		"SKIBIDI DOP DOP DOP YES YES ~anonymous",
		"Fortune sides with him who dares. ~Virgil",
	}

	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		idx := randRange(0, len(quotes))

		resp.Write([]byte(quotes[idx]))

	})

	http.ListenAndServe(":8080", router)
}

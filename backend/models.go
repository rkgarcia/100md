package main

type Question struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Answer   []Answer `json:"answers"`
}

type Answer struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Votes       int    `json:"votes"`
}

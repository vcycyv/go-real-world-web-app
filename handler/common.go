package handler

type message struct {
	InvalidMessage string
}

var Message = message{
	InvalidMessage: "Invalid JSON payload received.",
}

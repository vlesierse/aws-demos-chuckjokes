package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ListJokes",
		"GET",
		"/jokes",
		HandleListJokes,
	},
	Route{
		"GetRandomJoke",
		"GET",
		"/jokes/random",
		HandleGetRandomJoke,
	},
	Route{
		"GetJokeById",
		"GET",
		"/jokes/{jokeId}",
		HandleGetJokeById,
	},
}

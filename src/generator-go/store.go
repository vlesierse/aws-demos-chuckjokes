package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Joke struct {
	Id       int    `json:"id"`
	Joke     string `json:"joke"`
	Explicit bool   `json:"explicit"`
}

type Jokes []Joke

var jokes Jokes
var jokesDict map[int]Joke

func init() {
	file, _ := ioutil.ReadFile("data/jokes.json")
	_ = json.Unmarshal([]byte(file), &jokes)

	jokesDict = make(map[int]Joke, len(jokes))
	for _, joke := range jokes {
		jokesDict[joke.Id] = joke
	}
	explicitJokes := filter(jokes, func(joke Joke) bool { return joke.Explicit })
	nonExplicitJokes := filter(jokes, func(joke Joke) bool { return !joke.Explicit })

	fmt.Println("Jokes: ", len(jokes))
	fmt.Println("Explicit Jokes: ", len(explicitJokes))
	fmt.Println("Non-Explicit Jokes: ", len(nonExplicitJokes))
}

func filter(jokes Jokes, test func(Joke) bool) (ret Jokes) {
	for _, joke := range jokes {
		if test(joke) {
			ret = append(ret, joke)
		}
	}
	return
}

func GetRandomJoke() Joke {
	return jokes[rand.Intn(len(jokes))]
}

func GetJokeById(id int) Joke {
	return jokesDict[id]
}

func ListJokes() Jokes {
	return jokes
}

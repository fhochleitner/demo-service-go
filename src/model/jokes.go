package model

import (
	"fmt"
	"math/rand"
)

var jokes map[int]Joke

type Joke struct {
	Joke   string `json:"joke"`
	Rating int    `json:"rating"`
}

func AllJokes() map[int]Joke {
	return jokes
}

func RandomJoke() Joke {
	return jokes[rand.Intn(len(jokes)-1)]
}

func GetJoke(id int) Joke {
	if joke, ok := jokes[id]; ok {
		return joke
	} else {
		return Joke{}
	}
}

func (j *Joke) String() string {
	return fmt.Sprintf("Joke: %s\nRating: %d\n\n", j.Joke, j.Rating)
}

func init() {
	jokes = map[int]Joke{
		1:  {Joke: "Why don't scientists trust atoms? Because they make up everything!", Rating: 4},
		2:  {Joke: "Did you hear about the mathematician who's afraid of negative numbers? He will stop at nothing to avoid them!", Rating: 5},
		3:  {Joke: "What do you call a bear with no teeth? A gummy bear!", Rating: 3},
		4:  {Joke: "Why don't skeletons fight each other? They don't have the guts!", Rating: 4},
		5:  {Joke: "Why did the scarecrow win an award? Because he was outstanding in his field!", Rating: 5},
		6:  {Joke: "I'm reading a book about anti-gravity. It's impossible to put down!", Rating: 4},
		7:  {Joke: "Why did the bicycle fall over? It was two-tired!", Rating: 3},
		8:  {Joke: "How does a penguin build its house? Igloos it together!", Rating: 4},
		9:  {Joke: "Why did the tomato turn red? Because it saw the salad dressing!", Rating: 4},
		10: {Joke: "Why don't eggs tell jokes? Because they might crack up!", Rating: 3},
		11: {Joke: "Why did the golfer bring two pairs of pants? In case he got a hole in one!", Rating: 5},
		12: {Joke: "What do you call a snowman with a six-pack? An abdominal snowman!", Rating: 4},
		13: {Joke: "How do you organize a space party? You planet!", Rating: 3},
		14: {Joke: "I got a job at a bakery because I kneaded dough!", Rating: 4},
		15: {Joke: "Why don't scientists trust atoms? Because they make up everything!", Rating: 4},
		16: {Joke: "Did you hear about the mathematician who's afraid of negative numbers? He will stop at nothing to avoid them!", Rating: 5},
		17: {Joke: "What do you call a bear with no teeth? A gummy bear!", Rating: 3},
		18: {Joke: "Why don't skeletons fight each other? They don't have the guts!", Rating: 4},
		19: {Joke: "Why did the scarecrow win an award? Because he was outstanding in his field!", Rating: 5},
		20: {Joke: "I'm reading a book about anti-gravity. It's impossible to put down!", Rating: 4},
	}
}

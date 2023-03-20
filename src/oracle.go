// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go func() {
		for {
			select {
			case question := <-questions:
				go func() {
					generateAnswers(question, answers)
				}()
			case <-time.After(time.Duration(rand.Intn(10)+10) * time.Second):
				go func() {
					oraclePrediction(answers)
				}()
			}
		}
	}()

	go func() {
		for answer := range answers {
			printAnswer(answer)
		}
	}()

	return questions
}

func generateAnswers(question string, answer chan<- string) {
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	var response string

	if strings.Contains(strings.ToLower(question), "meaning of life") {
		response = "Ah, life! It is a question that has puzzled many a great mind. Some say the meaning of life is to seek happiness, others say it is to fulfill a divine purpose, and some think it's as simple as 42. What do you think?"
	} else if strings.Contains(strings.ToLower(question), "what is love") {
		response = "Love is a complex emotion that can take many forms. Some describe it as a deep connection with another person, while others believe it to be a chemical reaction in the brain. What is your definition of love?"
	} else if strings.Contains(strings.ToLower(question), "what is happiness") {
		response = "Happiness is a subjective experience that means different things to different people. Some find happiness in material possessions, while others find it in their relationships or hobbies. What brings you happiness?"
	} else {
		response = "I'm sorry, I don't know the answer to that question. Can you try asking me something else?"
	}

	//Prophecy generation
	if response == "" {
		longestWord := ""
		words := strings.Fields(question)
		for _, w := range words {
			if len(w) > len(longestWord) {
				longestWord = w
			}
		}
		nonsense := []string{
			"The moon is dark.",
			"The sun is bright.",
			"The sky is blue.",
			"The grass is green.",
			"The sky is falling.",
			"The snow is white.",
			"The rain is wet.",
		}
		response = longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
	}
	answer <- response
}

func oraclePrediction(answer chan<- string) {
	predictions := []string{
		"A great opportunity is coming your way.",
		"You will soon meet someone special.",
		"Money will come to you unexpectedly.",
		"Your hard work will pay off.",
		"You will receive good news soon.",
	}
	prediction := predictions[rand.Intn(len(predictions))]
	answer <- "Oracle prediction: " + prediction
}

func printAnswer(answer string) {
	for _, char := range answer {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type GoodbyeNickMessagePacket struct {
	colleagueName   string
	goodbyeMessages []string
}

func main() {
	nick()
}
func nick() {
	// Default text
	color.Cyan("\n\nHey Nick!")
	color.Cyan("As a token of our appreciation, respect and friendship, we've made you a little something!\nThis is your \"Goodbye Giftcard\" :)")

	time.Sleep(3 * time.Second)

	color.Yellow("\nWe wish you Farewell!")
	color.Yellow("Thank you for being an amazing colleague and friend.")
	color.Yellow("We wish you all the best in your future endeavors!")
	color.Yellow("\nDon't forget to stay in touch!")
	color.Yellow("Sincerely, team Peagsus")
	color.Blue("Team Peagsus, April 11 2023")

	time.Sleep(6 * time.Second)

	//Send goodbyes 🥲 ; random goodbye each time, randmon colleague as well
	sendRandomGoodbyeMessages()
}

func sendRandomGoodbyeMessages() {

	messagePacks := []GoodbyeNickMessagePacket{
		BogdansGoodbyeWordsToNick,
		JennasGoodbyeWordsToNick,
		JoesGoodbyeWordsToNick,
		StephensGoodbyeWordsToNick,
	}

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(messagePacks), func(i, j int) {
		messagePacks[i], messagePacks[j] = messagePacks[j], messagePacks[i]
	})

	for _, pack := range messagePacks {
		color.Set(color.FgHiGreen)
		fmt.Printf("\nFrom %s: ", pack.colleagueName)
		color.Unset()

		msgIndex := rand.Intn(len(pack.goodbyeMessages))
		selectedMessage := pack.goodbyeMessages[msgIndex]

		color.Set(color.FgHiWhite)
		typeMessage(selectedMessage, 50*time.Millisecond)
		color.Unset()

		fmt.Println()
	}
}

func typeMessage(message string, delay time.Duration) {
	for _, char := range message {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
}

package main
import "fmt"
import "math/rand"
import "time"


/*

import "bufio"
import "os"
import "strconv"

Card Dealer (Soon to be black jack, I hope...)

Author: Garett Skaar

Go LW1 program CS-354

*/

//Main will build a deck of cards, shuffle the cards and ask for a number
//of cards to be dealt.
func main() {

	//Build a new deck, unshuffled...
	deck := buildDeck()
	//Shuffle deck...
	shuffleDeck(deck)
	//prompt user
	fmt.Println("Deal how many cards?")
	var userInput int
	//scan standard input and check for errors
	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
            fmt.Println(err)
    }
	//display the desired amount of cards
	Deal(deck, userInput)
	
	/*
	needs implemented for black jack
	
	if strings.Contains(userInput, "y"){
		fmt.Println()
		startGame(deck)
	} else if strings.Contains(userInput, "n"){
		fmt.Println("Maybe next time!")
	} else{
		fmt.Println("Maybe next time!")
	}
	*/
	
}
//Deals the specified number of cards by user
func Deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d[i].value,"of",d[i].suit)
	}
}
//needs work
func startGame(deck Deck) {

	dealer := []Card{}
	player := []Card{}
	//var dealerVal int = 0
	//var playerVal int = 0
	
	//bust bool := false
	fmt.Println("Dealing cards...")
	for i :=0; i < len(deck) - 1;i++ {
		dealer = append(dealer, deck[i])
		player = append(player, deck[i+1])
	}
}

//struct type holds suit type and value (Face, Ace, or Number)
type Card struct {
	value string
	suit string
}

//Array of Card structs
type Deck []Card



//Builds and returns a new Card array
func buildDeck() (deck Deck) {

	//all valid values of a card
	values := []string{"Two", "Three", "Four", "Five", "Six", "Seven",
			"Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
			
	//all valid suits	
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	
	// Loop over each value and suit appending to the deck
	for i := 0; i < len(values); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				value: values[i],
				suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return deck 	
		
}
//Returns a shuffled Deck of cards
func shuffleDeck(deck Deck) Deck {

	//new seed based on current time
	rand.Seed(time.Now().UnixNano())
	
	//Loops through deck, creates random int within deck size and current index, and
	//swaps card postions if the index and random int are different.
	
	for i := 1; i < len(deck); i++ {
		
		random := rand.Intn(i + 1)
		//swap the cards
		if i != random {
			deck[random], deck[i] = deck[i], deck[random]
		}
	return deck
	}
	}
	
package main //This file belongs to the package "main", which makes it executable
import "fmt" //Import all functions from a reusable or library package

//OR import (	//alternative syntax to import multiple packages at once
//	"fmt"
//  "String"
//)

func main() { //Execution starts here (mandatory for executable file)
	//Create a deck of cards, create a hand and also print the remaining cards in the deck
	cardsToFile := newDeck()
	hand, leftDeck := deal(cardsToFile, 3) //Use := because we are initializing the variables before storing the return value
	fmt.Println("The hand dealt in the deal:")
	hand.print() //Reuse the print function to print the values of the deck type
	fmt.Println("The remaining deck:")
	leftDeck.print()

	//Part 9: Write the byte slice to file
	fmt.Println("Part 9: Write the byte slice to file")
	fmt.Println("SaveToFile error:", cardsToFile.saveToFile("my_cards")) // Prints: <Nil> or the error

	//Part 10: Reading from Hard Drive and Handle error
	fmt.Println("Part 10: Reading from Hard Drive and Handle error")
	cardsFromFile := newDeckFromFile("my_cards")
	cardsFromFile.print() //Should print the deck read from the file

	//Part 11: Shufle the deck
	fmt.Println("Part 11: Shufle the deck")
	cardsFromFile.shuffle()
	cardsFromFile.print()

	//Part 12: Shuffle the deck with seed
	fmt.Println("Part 12: Shuffle the deck with seed")
	myCardsToRandomShuffle := newDeck()
	myCardsToRandomShuffle.randomShuffle()
	myCardsToRandomShuffle.print()

}

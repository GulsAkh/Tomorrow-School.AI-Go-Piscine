package piscine

import "fmt"

func printNum(cards []int) {
	for i, card := range cards {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", card)
	}
	fmt.Printf("\n")
}

func DealAPackOfCards(deck []int) {
	n := 12 / 4

	player1 := (deck[:n])
	player2 := deck[n : 2*n]
	player3 := deck[2*n : 3*n]
	player4 := deck[3*n:]

	fmt.Printf("Player 1: ")
	printNum(player1)

	fmt.Printf("Player 2: ")
	printNum(player2)

	fmt.Printf("Player 3: ")
	printNum(player3)

	fmt.Printf("Player 4: ")
	printNum(player4)
}

// func main() {
// 	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	DealAPackOfCards(deck)
// }

package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for _, c := range deck {
		switch c {
		case 1:
			fmt.Printf("Player 1: %v, %v, %v\n", deck[0], deck[1], deck[2])
		case 3:
			fmt.Printf("Player 2: %v, %v, %v\n", deck[3], deck[4], deck[5])
		case 6:
			fmt.Printf("Player 3: %v, %v, %v\n", deck[6], deck[7], deck[8])
		case 9:
			fmt.Printf("Player 4: %v, %v, %v\n", deck[9], deck[10], deck[11])
		}
	}
}

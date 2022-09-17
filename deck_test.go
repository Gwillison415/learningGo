package main

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := newDeck()
	if len(newDeck) != 16 {
		t.Errorf("expected deck length of 16 but got: %d", len(newDeck))
		// 		t.Error("expected deck length of 20 but got:", len(newDeck))
		// same same
		if newDeck[0] != "Ace of spades" {
			t.Errorf("expected first card in deck to be 'Ace of spades' but got %v", newDeck[0])
		}
		if newDeck[len(newDeck)-1] != "four of Clubs" {
			t.Errorf("expected first card in deck to be 'four of Clubs' but got %v", newDeck[len(newDeck)-1])
		}
	}
}

// func Test_deck_toString(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		d    deck
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.d.toString(); got != tt.want {
// 				t.Errorf("deck.toString() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_newDeckFromFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeckFromFile(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeckFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
 
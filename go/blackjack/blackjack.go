package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		fallthrough
	case "jack":
		fallthrough
	case "king":
		fallthrough
	case "queen":
		return 10
	default:
		return 0
	}
}

type Moves struct {
	split string
	stand string
	hit   string
	auto  string
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	var moves Moves = Moves{split: "P", stand: "S", hit: "H", auto: "W"}

	mySum := ParseCard(card1) + ParseCard(card2)
	dealerSum := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return moves.split
	case (mySum == 21) && !(dealerSum >= 10):
		return moves.auto
	case (mySum <= 11):
		fallthrough
	case (mySum >= 12) && (mySum <= 16) && dealerSum >= 7:
		return moves.hit
	case (mySum == 21) && (dealerSum >= 10):
		fallthrough
	case (mySum >= 17) && (mySum <= 20):
		fallthrough
	case (mySum >= 12) && (mySum <= 16):
		fallthrough
	default:
		return moves.stand
	}
}

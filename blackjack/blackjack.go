package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	score := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case score == 22:
		return "P"
	case score == 21 && dealer < 10:
		return "W"
	case score == 21 && dealer >= 10:
		return "S"
	case score >= 17 && score <= 20:
		return "S"
	case score >= 12 && score <= 16 && dealer > 6:
		return "H"
	case score >= 12 && score <= 16 && dealer < 7:
		return "S"
	case score <= 11:
		return "H"
	default:
		return ""
	}
}

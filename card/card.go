package card

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	SuitCount = 4

	SuitDiamonds        = "diamonds"
	SuitDiamondsUnicode = "\u2666"

	SuitClubs        = "clubs"
	SuitClubsUnicode = "\u2663"

	SuitHearts        = "hearts"
	SuitHeartsUnicode = "\u2665"

	SuitSpades        = "spades"
	SuitSpadesUnicode = "\u2660"
)

const (
	FaceCount = 13
	Face2     = "2"
	Face3     = "3"
	Face4     = "4"
	Face5     = "5"
	Face6     = "6"
	Face7     = "7"
	Face8     = "8"
	Face9     = "9"
	Face10    = "10"
	FaceJack  = "J"
	FaceQueen = "Q"
	FaceKing  = "K"
	FaceAce   = "A"
)

type Card struct {
	Suit string
	Face string
}

func isValidSuit(suit string) bool {
	switch suit {
	case SuitClubs, SuitDiamonds, SuitHearts, SuitSpades:
		return true
	default:
		return false
	}
}

func isValidFace(face string) bool {
	switch face {
	case Face2, Face3, Face4, Face5, Face6, Face7, Face8, Face9, Face10, FaceJack, FaceQueen, FaceKing, FaceAce:
		return true
	default:
		return false
	}
}

func randomSuit(rand rand.Rand) string {
	index := rand.Intn(SuitCount)
	suits := []string{SuitHearts, SuitDiamonds, SuitSpades, SuitClubs}
	return suits[index]
}

func randomFace(random rand.Rand) string {
	index := random.Intn(FaceCount)
	suits := []string{Face2, Face3, Face4, Face5, Face6, Face7, Face8, Face9, Face10, FaceJack, FaceQueen, FaceKing, FaceAce}
	return suits[index]
}

func (c Card) SuitUnicode() (string, error) {
	switch c.Suit {
	case SuitClubs:
		return SuitClubsUnicode, nil
	case SuitSpades:
		return SuitSpadesUnicode, nil
	case SuitHearts:
		return SuitHeartsUnicode, nil
	case SuitDiamonds:
		return SuitDiamondsUnicode, nil
	default:
		return "", errors.New(fmt.Sprintf("unrecognized suit %s", c.Suit))
	}
}

func (c Card) ShortRepresentation() (string, error) {
	unicode, err := c.SuitUnicode()
	if err != nil {
		return "", err
	}
	if !isValidFace(c.Face) {
		return "", errors.New("card face is invalid")
	}
	return fmt.Sprintf("%s%s", unicode, c.Face), nil
}

func New(suit string, face string) (*Card, error) {
	if isValidSuit(suit) && isValidFace(face) {
		return &Card{
			Suit: suit,
			Face: face,
		}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Cannot construct Card with suit %s, face %s", suit, face))
	}
}

func Random(random rand.Rand) (*Card, error) {
	suit := randomSuit(random)
	face := randomFace(random)
	return New(suit, face)
}

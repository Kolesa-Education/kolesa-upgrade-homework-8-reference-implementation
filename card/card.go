package card

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"unicode/utf8"
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

const (
	NumericValueJack  = 11
	NumericValueQueen = 12
	NumericValueKing  = 13
	NumericValueAce   = 14
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

func (c Card) IsNumeric() bool {
	switch c.Face {
	case Face2, Face3, Face4, Face5, Face6, Face7, Face8, Face9, Face10:
		return true
	default:
		return false
	}
}

func (c Card) StrictNumericValue() (int, error) {
	if c.IsNumeric() {
		atoi, err := strconv.Atoi(c.Face)
		if err != nil {
			return 0, err
		}
		return atoi, nil
	} else {
		return 0, errors.New("not-numeric cards cannot have numeric value")
	}
}

func (c Card) NumericValue() int {
	if c.IsNumeric() {
		numericValue, _ := c.StrictNumericValue()
		return numericValue
	} else {
		switch c.Face {
		case FaceJack:
			return NumericValueJack
		case FaceQueen:
			return NumericValueQueen
		case FaceKing:
			return NumericValueKing
		case FaceAce:
			return NumericValueAce
		default:
			return 0
		}
	}
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

func FromShortRepresentation(representation string) (*Card, error) {
	representation = strings.Trim(representation, "\n")
	suitUnicode, size := utf8.DecodeRuneInString(representation)
	face := representation[size:]
	suit, err := SuitOfUnicodeSymbol(string(suitUnicode))
	log.Printf("parsing cardCSVEntry {%s}, suit {%s}, face {%s}, len {%d}",
		representation,
		suit,
		face,
		len(representation),
	)
	if err != nil {
		return nil, err
	}
	resultCard := Card{
		Face: face,
		Suit: suit,
	}
	return &resultCard, nil
}

func Random(random rand.Rand) (*Card, error) {
	suit := randomSuit(random)
	face := randomFace(random)
	return New(suit, face)
}

func SuitOfUnicodeSymbol(unicode string) (string, error) {
	switch unicode {
	case SuitSpadesUnicode:
		return SuitSpades, nil
	case SuitHeartsUnicode:
		return SuitHearts, nil
	case SuitClubsUnicode:
		return SuitClubs, nil
	case SuitDiamondsUnicode:
		return SuitDiamonds, nil
	default:
		return "", errors.New("not implemented suit")
	}
}

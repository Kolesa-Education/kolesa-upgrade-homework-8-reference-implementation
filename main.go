package main

import (
	"fmt"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8-reference-implementation/card"
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8-reference-implementation/combinatorics"
	"github.com/samber/lo"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func processDatasetEntry(cards []card.Card) ([]card.PokerCombination, error) {
	var result []card.PokerCombination
	deduplicated := combinatorics.Deduplicate(cards)
	combinations, err := combinatorics.Combinations(deduplicated, card.ValidCombinationSize)
	if err != nil {
		return nil, err
	}
	for _, comb := range combinations {
		combination, err := card.CombinationOf(comb)
		if err != nil {
			return nil, err
		}
		if combination != nil {
			result = append(result, combination)
		}
	}
	return result, nil
}

func readCardsFromCSV(cardsCSV string) ([]card.Card, error) {
	cards := strings.Split(cardsCSV, ",")
	parsed := lo.Map[string, card.Card](cards, func(csvCard string, index int) card.Card {
		parsedCard, err := card.FromShortRepresentation(csvCard)
		if err != nil {
			log.Fatalln(err)
		}
		return *parsedCard
	})
	return parsed, nil
}

func processFile(dirName, fileName, resultDirName string) {
	inputData, err := os.ReadFile(fmt.Sprintf("%s/%s", dirName, fileName))
	if err != nil {
		log.Fatalln(err)
	}

	inputCards, err := readCardsFromCSV(string(inputData))
	if err != nil {
		log.Fatalln(err)
	}

	result, err := processDatasetEntry(inputCards)
	if err != nil {
		log.Fatalln(err)
	}

	resultCards := lo.Map[card.PokerCombination, []card.Card](result, func(t card.PokerCombination, i int) []card.Card {
		return t.Cards()
	})

	resultFile, err := os.OpenFile(
		fmt.Sprintf("%s/%s", resultDirName, fileName),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)

	if err != nil {
		log.Fatalln(err)
	}

	representations := lo.Map[[]card.Card, []string](resultCards, func(cards []card.Card, index int) []string {
		result := lo.Map[card.Card, string](cards, func(card2 card.Card, index2 int) string {
			r, err := card2.ShortRepresentation()
			if err != nil {
				log.Fatalln(err)
			}
			return r
		})
		return result
	})

	lo.ForEach[[]string](representations, func(rs []string, index int) {
		if index < card.ValidCombinationSize {
			_, err = resultFile.Write([]byte(fmt.Sprintf("%s", rs[index])))
			if err != nil {
				log.Fatalln(err)
			}
		}

	})

	_ = resultFile.Close()
}

func main() {
	log.Println("Starting counting combinations...")
	files, err := os.ReadDir("dataset/")

	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll(filepath.Join(".", "results"), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	for _, dirEntry := range files {
		log.Printf("iterating dirEntry %s\n", dirEntry)
		processFile("dataset", dirEntry.Name(), "results")
	}
}

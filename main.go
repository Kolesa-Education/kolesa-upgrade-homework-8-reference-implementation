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

	resultFile, err := os.OpenFile(
		fmt.Sprintf("%s/%s", resultDirName, fileName),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)

	if err != nil {
		log.Fatalln(err)
	}

	lo.ForEach[card.PokerCombination](result, func(combination card.PokerCombination, index int) {
		representation, err := combination.Representation()
		if err != nil {
			log.Fatalln(err)
		}

		_, err = resultFile.WriteString(fmt.Sprintf("%s\n", representation))
		if err != nil {
			log.Fatalln(err)
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

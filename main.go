package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

type shoppingItem struct {
	id     string
	name   string
	price  int
	rating int
	notes  string
}

func New(name string, price int, rating int, notes string) *shoppingItem {
	return &shoppingItem{uuid.New().String(), name, price, rating, notes}

}

func trimNewLine(input string) string {
	return strings.TrimFunc(input, func(r rune) bool { return r == '\n' })
}

func getItemAttributes() (*shoppingItem, error) {

	reader := bufio.NewReader(os.Stdin)

	println("Enter a name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	println("Enter the price: ")
	inputPrice, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	inputPrice = trimNewLine(inputPrice)
	price, err := strconv.Atoi(inputPrice)
	if err != nil {
		return nil, err
	}
	println("Enter the rating: ")
	inputRating, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	inputRating = trimNewLine(inputRating)

	rating, err := strconv.Atoi(inputRating)
	if err != nil {
		return nil, err
	}

	println("Add some notes: ")
	notes, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return New(name, price, rating, notes), nil

}
func addItem(ctx *cli.Context) error {

	item, err := getItemAttributes()
	if err != nil {
		return err
	}
	fmt.Printf("Item: %+v\n", item)

	// TODO: Write to file
	return nil

}

func main() {
	app := &cli.App{
		Name:  "Compare that",
		Usage: "Catalogue items and compare them side by side",
		Commands: []*cli.Command{
			{
				Name:   "add",
				Usage:  "add an item to the comparing 'scope'",
				Action: addItem,
			},
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/adantop/golang-bootcamp-2020/domain"
	"github.com/adantop/golang-bootcamp-2020/service"
	"github.com/docopt/docopt-go"
)

const usage = `Pokedex

Usage:
  pokedex <source> [--name=<NAME>|--id=<ID>]
  pokedex -h | --help

Options:
  -n NAME --name=NAME   The name of the pokemon
  -i ID --id=ID         The id of the pokemon
  -h --help             Show this screen.`

func main() {
	var (
		p        domain.Pokemon
		opts     = parseOpts()
		svc, err = service.New(opts["source"])
	)

	if err != nil {
		log.Fatalln(err)
	}
	defer svc.DS.Close()

	if opts["filter"] == "name" {
		p, err = svc.DS.GetPokemonByName(opts["value"])
	} else {
		number, err := strconv.Atoi(opts["value"])
		if err != nil {
			log.Panicf("ID provided is not a number %s\n", opts["value"])
		}
		p, err = svc.DS.GetPokemonByID(number)
	}

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(p.ExtendedDescription())
}

func parseOpts() map[string]string {
	opts := map[string]string{}

	args, err := docopt.ParseDoc(usage)
	if err != nil {
		log.Fatalln(err)
	}

	if opts["source"], err = args.String("<source>"); err != nil {
		log.Fatalln("Unable to obtain resource")
	}

	if name, _ := args.String("--name"); name != "" {
		opts["filter"] = "name"
		opts["value"] = name
	} else if number, _ := args.String("--number"); number != "" {
		opts["filter"] = "number"
		opts["value"] = number
	} else {
		log.Fatalln("Either pokemon name or ID is required")
	}

	return opts
}

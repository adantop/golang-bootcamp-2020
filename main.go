package main

import (
	"fmt"
	"log"

	"github.com/adantop/golang-bootcamp-2020/service"
	"github.com/docopt/docopt-go"
)

func main() {
	var (
		opts       = parseOpts()
		srcType, _ = opts.String("<SourceType>")
		svc, err   = service.New(srcType)
	)

	if err != nil {
		log.Fatalln(err)
	}
	defer svc.DS.Close()

	name, _ := opts.String("<PokemonName>")
	pokemon, err := svc.DS.GetPokemonByName(name)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pokemon.ExtendedDescription())
}

func parseOpts() docopt.Opts {
	usage := `Pokedex

Usage:
  pokedex <SourceType> <PokemonName>
  pokedex -h | --help

Options:
  -h --help     Show this screen.`

	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		log.Fatalln(err)
	}

	return opts
}

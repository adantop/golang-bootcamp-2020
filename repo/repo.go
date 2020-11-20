package repo

import "github.com/adantop/golang-bootcamp-2020/domain"

// Repo is an interface for collecting pokemon data
type Repo interface {
	GetPokemonByID(id int) (domain.Pokemon, error)
	GetPokemonByName(name string) (domain.Pokemon, error)
	GetAllPokemons() ([]domain.Pokemon, error)
}

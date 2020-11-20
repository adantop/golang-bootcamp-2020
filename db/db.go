package db

import (
	"database/sql"

	"github.com/adantop/golang-bootcamp-2020/domain"
)

const (
	queryPokemonByName = "SELECT * FROM pokemon WHERE name = $1"
	queryPokemonByID   = "SELECT * FROM pokemon WHERE id = $1"
	selectAllPokemon   = "SELECT * FROM pokemon"
)

type database struct {
	db *sql.DB
}

// GetPokemonByID initializes the datasource
func (ds database) GetPokemonByID(id int) (p domain.Pokemon, err error) {
	err = ds.db.QueryRow(queryPokemonByID, id).Scan(
		&p.Number,
		&p.Name,
		&p.Type1,
		&p.Type2,
		&p.HeightM,
		&p.WeightKg,
		&p.Male,
		&p.Female,
		&p.CaptRate,
		&p.HP,
		&p.Attack,
		&p.Defense,
		&p.Special,
		&p.Speed)

	return
}

// GetPokemonByName Get pokemon by name
func (ds database) GetPokemonByName(name string) (p domain.Pokemon, err error) {

	err = ds.db.QueryRow(queryPokemonByName, name).Scan(
		&p.Number,
		&p.Name,
		&p.Type1,
		&p.Type2,
		&p.HeightM,
		&p.WeightKg,
		&p.Male,
		&p.Female,
		&p.CaptRate,
		&p.HP,
		&p.Attack,
		&p.Defense,
		&p.Special,
		&p.Speed)

	return
}

// GetAllPokemons initializes the datasource
func (ds database) GetAllPokemons() (pokes []domain.Pokemon, err error) {

	rows, err := ds.db.Query(selectAllPokemon)

	if err != nil {
		return
	}
	defer rows.Close()

	for isRow := rows.Next(); !isRow; {

		var p domain.Pokemon

		err = rows.Scan(
			&p.Number,
			&p.Name,
			&p.Type1,
			&p.Type2,
			&p.HeightM,
			&p.WeightKg,
			&p.Male,
			&p.Female,
			&p.CaptRate,
			&p.HP,
			&p.Attack,
			&p.Defense,
			&p.Special,
			&p.Speed)

		if err != nil {
			return
		}
		pokes = append(pokes, p)
	}
	return
}

// Close terminates the database connection
func (ds database) Close() {
	ds.db.Close()
}

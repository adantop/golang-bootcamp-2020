package fs

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/adantop/golang-bootcamp-2020/domain"
	"github.com/adantop/golang-bootcamp-2020/repo"
)

// CSV is the object used to read the csv files
type CSV struct{ file *os.File }

// OpenCSV initializes the datasource
func OpenCSV(filename string, ds *repo.DataSource) (err error) {
	file, err := os.Open(filename)

	(*ds) = CSV{file}
	
	return
}

// GetPokemonByID initializes the datasource
func (ds CSV) GetPokemonByID(id int) (p domain.Pokemon, err error) {
	var (
		reader = csv.NewReader(ds.file)
		idStr  = strconv.Itoa(id)
	)

	// skip header
	reader.Read()

	for {
		r, e := reader.Read()

		if e == io.EOF {
			err = fmt.Errorf("Pokemon %d not found", id)
			return
		}

		if e != nil {
			err = e
			return
		}

		if idStr == r[0] {
			err = p.UpdateFromSlice(&r)
			return
		}
	}
}

// GetPokemonByName Get pokemon by name
func (ds CSV) GetPokemonByName(name string) (p domain.Pokemon, err error) {
	reader := csv.NewReader(ds.file)

	// skip header
	reader.Read()

	for {
		r, e := reader.Read()

		if e == io.EOF {
			err = fmt.Errorf("Pokemon %s not found", name)
			return
		}

		if e != nil {
			err = e
			return
		}

		if name == r[1] {
			err = p.UpdateFromSlice(&r)
			return
		}
	}
}

// GetAllPokemons initializes the datasource
func (ds CSV) GetAllPokemons() ([]domain.Pokemon, error) {
	var (
		err    error
		pokes  = []domain.Pokemon{}
		reader = csv.NewReader(ds.file)
	)

	// skip header
	reader.Read()

	for {
		var (
			p    domain.Pokemon
			r, e = reader.Read()
		)

		if e == io.EOF {
			return pokes, nil
		}

		if e != nil {
			err = e
			return nil, err
		}

		if err = p.UpdateFromSlice(&r); err != nil {
			return nil, err
		}

		pokes = append(pokes, p)
	}

}

// Close terminates the database connection
func (ds CSV) Close() {
	ds.file.Close()
}

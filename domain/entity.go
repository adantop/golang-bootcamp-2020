package domain

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

// Pokemon model for the pokemon
type Pokemon struct {
	Number   int
	Name     string
	Type1    sql.NullString
	Type2    sql.NullString
	HeightM  float64
	WeightKg float64
	Male     float64
	Female   float64
	CaptRate float64
	HP       int
	Attack   int
	Defense  int
	Special  int
	Speed    int
}

// ShortDescription returns the pokemon name and type
func (p *Pokemon) ShortDescription() string {
	var types []string

	for _, val := range []sql.NullString{p.Type1, p.Type2} {
		if val.Valid {
			types = append(types, strings.Title(val.String))
		}
	}

	return fmt.Sprintf("{%03d - %v: %v}", p.Number, p.Name, strings.Join(types, " "))
}

// ExtendedDescription shows detailed information on the pokemon
func (p *Pokemon) ExtendedDescription() string {

	return fmt.Sprintf(
		"%v\n  HP:  %4d\n  Atk: %4d\n  Def: %4d\n  Spe: %4d\n  Spd: %4d",
		p.ShortDescription(),
		p.HP,
		p.Attack,
		p.Defense,
		p.Special,
		p.Speed)
}

// UpdateFromSlice will convert the text values into it's appropiate tipe and update Pokemon properties
func (p *Pokemon) UpdateFromSlice(r *[]string) (err error) {

	if p.Number, err = strconv.Atoi((*r)[0]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Number, column 1")
		return
	}

	p.Name = (*r)[1]
	p.Type1 = sql.NullString{String: (*r)[2], Valid: (*r)[2] != ""}
	p.Type2 = sql.NullString{String: (*r)[3], Valid: (*r)[3] != ""}

	if p.HeightM, err = strconv.ParseFloat((*r)[4], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Height, column 5")
		return
	}

	if p.WeightKg, err = strconv.ParseFloat((*r)[5], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Weight, column 6")
		return
	}

	if p.Male, err = strconv.ParseFloat((*r)[6], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Male ratio, column 7")
		return
	}

	if p.Female, err = strconv.ParseFloat((*r)[7], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Female ratio, column 8")
		return
	}

	if p.CaptRate, err = strconv.ParseFloat((*r)[8], 64); err != nil {
		err = fmt.Errorf("Could not parse float for pokemon Capture rate, column 6")
		return
	}

	if p.HP, err = strconv.Atoi((*r)[9]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon HP, column 10")
		return
	}

	if p.Attack, err = strconv.Atoi((*r)[10]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Attack, column 11")
		return
	}

	if p.Defense, err = strconv.Atoi((*r)[11]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Defense, column 12")
		return
	}

	if p.Special, err = strconv.Atoi((*r)[12]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Special, column 13")
		return
	}

	if p.Speed, err = strconv.Atoi((*r)[13]); err != nil {
		err = fmt.Errorf("Could not parse int for pokemon Speed, column 14")
		return
	}
	return
}

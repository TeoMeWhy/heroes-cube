package db

import (
	"database/sql"
)

type Person struct {
	Id          string
	Name        string
	Strength    int
	Agility     int
	Inteligence int
	Damage      int
	HitPoints   int
	Defense     int
	Class       string
	Race        string
	Exp         int
	Level       int
}

func CreatePerson(p *Person, con *sql.DB) error {
	query := `
	INSERT INTO persons (
		Id,
		Name,
		Strength,
		Agility,
		Inteligence,
		Damage,
		HitPoints,
		Defense,
		Class,
		Race,
		Exp,
		Level
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	state, err := con.Prepare(query)
	if err != nil {
		return err
	}

	_, err = state.Exec(
		p.Id,
		p.Name,
		p.Strength,
		p.Agility,
		p.Inteligence,
		p.Damage,
		p.HitPoints,
		p.Defense,
		p.Class,
		p.Race,
		p.Exp,
		p.Level,
	)

	return err

}

func GetPerson(id string, con *sql.DB) (*Person, error) {

	query := `
	SELECT
		Id,
		Name,
		Strength,
		Agility,
		Inteligence,
		Damage,
		HitPoints,
		Defense,
		Class,
		Race,
		Exp,
		Level
	FROM persons
	WHERE Id = ?;`

	state, err := con.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := state.Query(id)
	if err != nil {
		return nil, err
	}

	p := &Person{}
	for rows.Next() {
		rows.Scan(
			&p.Id,
			&p.Name,
			&p.Strength,
			&p.Agility,
			&p.Inteligence,
			&p.Damage,
			&p.HitPoints,
			&p.Defense,
			&p.Class,
			&p.Race,
			&p.Exp,
			&p.Level,
		)
	}

	return p, nil

}

func UpdatePerson(p *Person, con *sql.DB) error {
	query := `
	UPDATE persons
	SET	Id = ?,
		Name = ?,
		Strength = ?,
		Agility = ?,
		Inteligence = ?,
		Damage = ?,
		HitPoints = ?,
		Defense = ?,
		Class = ?,
		Race = ?,
		Exp = ?,
		Level = ?
	WHERE Id = ?`

	state, err := con.Prepare(query)
	if err != nil {
		return err
	}

	_, err = state.Exec(
		p.Id,
		p.Name,
		p.Strength,
		p.Agility,
		p.Inteligence,
		p.Damage,
		p.HitPoints,
		p.Defense,
		p.Class,
		p.Race,
		p.Exp,
		p.Level,
		p.Id,
	)

	return err

}

func DeletePerson(id string, con *sql.DB) error {
	query := `
	DELETE FROM persons
	WHERE Id = ?
	`

	state, err := con.Prepare(query)
	if err != nil {
		return err
	}

	_, err = state.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetPersonIDbyName(name string, con *sql.DB) (string, error) {

	query := `
	SELECT Id
	FROM persons
	WHERE Name = ?
	`

	state, err := con.Prepare(query)
	if err != nil {
		return "", err
	}

	rows, err := state.Query(name)
	if err != nil {
		return "", err
	}

	var idPerson string
	for rows.Next() {
		rows.Scan(&idPerson)
	}

	return idPerson, nil

}

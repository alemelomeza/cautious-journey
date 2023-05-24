package db

import (
	"database/sql"

	"github.com/alemelomeza/cautious-journey/internal/domain"
	"github.com/alemelomeza/cautious-journey/pkg/helper"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteDB struct {
	db *sql.DB
}

func NewSqliteDB(filename string) *sqliteDB {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		panic(err)
	}
	return &sqliteDB{
		db: db,
	}
}

func (r *sqliteDB) SaveChild(child domain.Child) error {
	parentsIDs := helper.ParentsToIDs(child.Parents...)
	_, err := r.db.Exec("INSERT INTO childrens VALUES (?, ?, ?);", child.Name, child.Age, parentsIDs)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqliteDB) GetChild(name string) (*domain.Child, error) {
	return nil, nil
}

func (r *sqliteDB) GetChildrens() ([]*domain.Child, error) {
	return nil, nil
}

func (r *sqliteDB) SaveParent(parent domain.Parent) error {
	_, err := r.db.Exec("INSERT INTO parents VALUES (?, ?);", parent.Name, parent.Relationship)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqliteDB) GetParent(name string) (*domain.Parent, error) {
	return nil, nil
}

func (r *sqliteDB) GetParents(childName string) ([]*domain.Parent, error) {
	return nil, nil
}

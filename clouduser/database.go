package main

import (
	"database/sql"
	"fmt"
	"github.com/pborman/uuid"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const DBNAME string = "CloudCommander.db"

func _getDBObject() (*sql.DB, error) {
	return sql.Open("sqlite3", fmt.Sprintf("/etc/cloud-commander/database/%s", DBNAME))
}

func ensureDB() error {
	return createDB()
}

func createDB() error {
	db, err := _getDBObject()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := `create table if not exists user
			(id integer primary key autoincrement, name text unique,
			access_key text, last_modified integer, active integer default 1,
		    	last_used integer default -1);`

	_, err = db.Exec(sqlStmt)
	return err
}

func insertUser(username string) (string, error) {
	db, err := _getDBObject()
	defer db.Close()
	if err != nil {
		return "", err
	}

	err = ensureDB()
	if err != nil {
		return "", err
	}
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := tx.Prepare("insert into user (name, access_key, last_modified) values (?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	access_key := uuid.New()
	_, err = stmt.Exec(username, access_key, time.Now().Unix())
	if err != nil {
		return "", err
	}
	tx.Commit()
	return access_key, nil
}

func getUser(username string) (*User, error) {
	user := User{}
	db, err := _getDBObject()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	stmt, err := db.Prepare("select id, name, access_key, active, last_modified, last_used from user where name = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	err = unpackUser(row, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func deleteUser(username string) error {
	db, err := _getDBObject()
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("delete from user where name = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	return nil
}

func unpackUser(row *sql.Row, user *User) error {
	err := row.Scan(&user.Id, &user.Name, &user.AccessKey, &user.Active, &user.LastModified, &user.LastUsed)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"log"
	"os"
	"testing"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS air_quality (
	id SERIAL,
	location TEXT NOT NULL,
	CONSTRAINT air_quality_pkey PRIMARY KEY (id)
)`

var a App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM air_quality")
	a.DB.Exec("ALTER SEQUENCE air_quality_id_seq RESTART WITH 1")
}

package main

import (
	"database/sql"
	"errors"
)

type AirQuality struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
}

func (aq *AirQuality) GetAirQuality(db *sql.DB) error {
	return errors.New("not implemented")
}

func (aq *AirQuality) UpdateAirQuality(db *sql.DB) error {
	return errors.New("not implemented")
}

func (aq *AirQuality) DeleteAirQuality(db *sql.DB) error {
	return errors.New("not implemented")
}

func (aq *AirQuality) CreateAirQuality(db *sql.DB) error {
	return errors.New("not implemented")
}

func getAirQuality(db *sql.DB, start, count int) ([]AirQuality, error) {
	return nil, errors.New("not implemented")
}

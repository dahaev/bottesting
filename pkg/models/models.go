package models

import "time"

type Account struct {
	ID          int
	UserName    string
	Location    string
	Region      string
	Rating      int
	Created     time.Time
	Description string
}

type DonAccount struct {
	ID       int
	UserName string
	Rating   int
	Created  time.Time
}

type Review struct {
	ID          int
	LadyName    string
	DonName     string
	Description string
	Rating      int
	Date        time.Time
}

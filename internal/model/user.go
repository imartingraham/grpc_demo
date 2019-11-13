package model

import "log"

// User struct mapping to DB
type User struct {
	ID            int
	ParentID      int
	Email         string
	Prefix        string
	FirstName     string
	MiddleInitial string
	LastName      string
	Suffix        string
	Status        string
	Type          string
	Perms         int
	Password      string
	Created       string // timestamp
	Modified      string // timestamp
	LastLogin     string
	SessionID     string
	ReferralID    int
	OTPCode       string
	OTPSecret     string
	OTPSentCount  int
	OTPSent       string // timestamp
	OTPVerified   string // timestamp
	AutoLogin     string
	RegComplete   string //timestamp
	PassReset     bool
	BruteForce    int
}

// FindUser looks up user by id and return it if found
func FindUser(id int64) (*User, error) {
	user := User{}
	sql := `SELECT id, first_name, last_name FROM users WHERE id = $1`
	err := db.QueryRow(sql, id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		log.Fatalf("Error querying user: %v", err)
		return nil, err
	}
	return &user, nil
}

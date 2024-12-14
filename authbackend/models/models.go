package models

// to be used when authenticating
type Auth struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	NationID  string `json:"national_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// peer represents the investor/lender or the borrower/client
type Peer struct {
	FirstName string
	LastName  string
	UserName  string
	NationID  string
	Email     string
	Balance   float64 // amount the peer has
	LenderId  string  // the id of the lender who gave the peer money
}

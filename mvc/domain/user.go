package domain

//without json we have result as ID,FirstName but with it it looks this way first_name
type User struct{
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email 	  string `json:"email"`
}


package types

import(
	"time"
)

type UserStore interface{
	getuserbyemail(email string )(*types.User, error)
	getuserbyid(id int)(*types.User, error)
	createuser(User) (error)
}	

type user struct {

	ID int
	firstname string `json:"id"`
	lastname string `json:"firstname"`
	email string `json:"firstname"`
	password string `json:"-"`
	createdat time.time `json:"createdat"`
}



type registeruserpayload struct {
	firstname    string `json:"firstname"`
	lastname   string `json:"lastname"`
	email 	string `json:"email"`
	password string `json:"password" '
}
package USER

import (
	"complete backend/types"
	"database/sql"
	"fmt"
)

type store struct{
	db *sql.DB
}

func newstore(db *sql.DB )*store{
	return &Store{db : db}
}

func (s *store) getuserbyemail(email string )(*types.User, error){
	rows, err := s.db.Query("select * from users where email =?",email)

	if err != nil{
		return nil, err
	}

	res,err := scanrowintouser(rows)

	if err != nil{
		return nil, err
	}

	if res.ID == 0{
		return nil,fmt.Errorf("user not found")
	}

	return res,nil
}

func scanrowintouser(rows *sql.Rows)(*types.User, error){
	user := new(types.user)

	if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt); err != nil{
		return nil, err
	}

	return user , nil
}

func (s *store) getuserbyid(id int)(*types.User, error){
	res,err := s.Query("select * from user where id =?",id)

	if err != nil{
		return nil,err
	}
	if res != nil{
		return res, nill
	}
	else{
		return nil,err
	}
}
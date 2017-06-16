package data

import (
	"time"
	"fmt"
)


type User struct {
	Id int
	Uuid string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}

//create new session for existing user
func (user *User) CreateSession() (session Session, err error) {
	//statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	statement := "insert into sessions set uuid=?, email=?, user_id=?, created_at=?" 		// for mysql
	//statement := "insert into sessions(uuid, email, user_id, created_at) values ($, $, $, $)" 	// for mysql

	stmt, err :=Db.Prepare(statement)
	if err !=nil {
		fmt.Println("CreateSession_ERR1 !!")
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	//err=stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	_, err=stmt.Exec(createUUID(), user.Email, user.Id, time.Now())

	if err !=nil {
		fmt.Println("CreateSession_ERR2 !!")
		fmt.Println(err)
		return
	}
	return
}

func (user *User) Session() (session Session, err error) {
	session=Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (session *Session) Check() (valid bool, err error) {

	fmt.Println("Checksession_#1")

	statement := "select id, uuid, email, user_id, created_at from sessions where uuid=$"
	err =Db.QueryRow(statement, &session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.CreatedAt)
	if err !=nil {
		fmt.Println("Checksession_ERR1 !!")
		fmt.Println(err)
		valid=false
		return
	}
	fmt.Println("Checksession_#2")
	if session.Id != 0 {
		valid=true
	}
	return

}
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
	fmt.Println("user.CreateSession()_#1")
	stmt, err :=Db.Prepare(statement)
	if err !=nil {
		fmt.Println("CreateSession_ERR1 !!")
		fmt.Println(err)
		return
	}
	fmt.Println("user.CreateSession()_#2")
	defer stmt.Close()
	newUuid := createUUID()
	//err=stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	_, err=stmt.Exec(newUuid, user.Email, user.Id, time.Now())
	fmt.Println("user.CreateSession()_#3")

	if err !=nil {
		fmt.Println("CreateSession_ERR2 !!")
		fmt.Println(err)
		return
	}
	session.Uuid=newUuid
	session.UserId=user.Id
	session.Email=user.Email
	session.CreatedAt=time.Now()
	fmt.Println("user.CreateSession()_#3")

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
	fmt.Println("Checksession_#1.1 &session.uuid=", &session.Uuid)
	fmt.Println("Checksession_#1.2 session.uuid=", session.Uuid)
	fmt.Println("Checksession_#1.3 session.email=", session.Email)
	fmt.Println("Checksession_#1.4 session.userid=", session.UserId)

	statement := "select id, uuid, email, user_id, created_at from sessions where uuid=?"
	err =Db.QueryRow(statement, &session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err !=nil {
		fmt.Println("Checksession_ERR1 !!")
		fmt.Println(err)
		valid=false
		return
	}
	fmt.Println("Checksession_#2")
	fmt.Println("Checksession_#2.1 session.uuid=", &session.Uuid)
	fmt.Println("Checksession_#2.2 session.uuid=", session.Uuid)
	fmt.Println("Checksession_#2.3 session.email=", session.Email)
	fmt.Println("Checksession_#2.4 session.userid=", session.UserId)
	if session.Id != 0 {
		valid=true
	}
	return

}

func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid=?"
	fmt.Println("DeleteByUUID_#1 : session.uuid=", session.Uuid)
	stmt, err :=Db.Prepare(statement)
	fmt.Println("DeleteByUUID_#2")
	if err != nil {
		fmt.Println("DeleteByUUID_#RR1 !!", err)
		return
	} else  {
		fmt.Println("DeleteByUUID_#2a : OK!")
	}
	defer stmt.Close()
	fmt.Println("DeleteByUUID_#3")
	_, err =stmt.Exec(session.Uuid)
	if err != nil {
		fmt.Println("DeleteByUUID_#RR2 !!", err)
	}
	fmt.Println("DeleteByUUID_#4")
	return



}

//get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err =Db.QueryRow("select id, uuid, name, email, created_at from users where id=?", session.UserId).
		Scan(&user.Id,&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func (user *User) Create() (err error) {
	fmt.Println("user.Create()_#1")
	// statement_postgres := "insert into users (uuid, name, email, password, created_at) values ($1,$2,$3,$4,$5) returning id, uuid, created_at"

	statement := "insert into users (uuid, name, email, password, created_at) values (?,?,?,?,?)"
	stmt, err :=Db.Prepare(statement)
	fmt.Println("user.Create()_#2")
	if err !=nil {
		fmt.Println("user.Create()_ERR1 !!", err)
		return
	}
	defer stmt.Close()

	fmt.Println("user.Create()_#3")

	//err =stmt.QueryRow(createUUID(), user.Name, user.Email, user.Password).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	_, err =stmt.Exec(createUUID(), user.Name, user.Email, user.Password, time.Now())

	if err !=nil {
		fmt.Println("user.Create()_ERR2 !!")
		return
	}
	fmt.Println("user.Create()_#4")

	return

}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	fmt.Println("userByEmail_#1")
	err = Db.QueryRow("select id, uuid, name, email, password,  created_at from users where email=?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	fmt.Println("userByEmail_#2")
	if err !=nil {
		fmt.Println("UserByEmail_ERR1 !!")
		fmt.Println(err)
	}
	fmt.Println("userByEmail_#3")
	return

}


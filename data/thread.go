package data

import (
	"time"
	"fmt"
)

type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}

type Post struct {
	Id int
	Uuid string
	Body string
	UserId int
	ThreadId int
	CreatedAt time.Time
}

func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")

}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")

}

func (thread *Thread) NumReplies () (count int) {
	rows, err :=Db.Query("select count(*) from posts where thread_id=?", thread.Id)
	fmt.Println("numReplies_#1")
	if err !=nil {
		fmt.Println("numReplies_ERR1 !!")
		return
	}
	fmt.Println("numReplies_#2")
	for rows.Next() {
		if err =rows.Scan(&count); err!=nil {
			fmt.Println("numReplies_ERR2 !!")

			return
		}
	}
	rows.Close()
	fmt.Println("numReplies_#3")

	return
}

func (thread* Thread) Posts() (posts []Post, err error) {
	rows, err :=Db.Query("select id, uuid, body, user_id, thread_id, created_at from posts where thread_id=?",  thread.Id)
	if err !=nil {
		return
	}
	for rows.Next() {
		post := Post {}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err !=nil {
			fmt.Println("thread.Posts()_ERR1 !!", err)
			return
		}
		posts =append(posts, post)
	}
	rows.Close()
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	//statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?) returning id, uuid, topic, user_id, created_at"
	statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	//err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	newUuid := createUUID()
	_, err=stmt.Exec(newUuid, topic, user.Id, time.Now())

	return
}

// Create a new post to a thread ????

func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	//statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	//err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	newUuid := createUUID()
	_,err = stmt.Exec(newUuid, body, user.Id, conv.Id, time.Now())
	return
}

func Threads() (threads []Thread, err error) {
	fmt.Println("Threads()_#1")
	rows, err :=Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err !=nil {
		fmt.Println("Threads()_ERR1 !!")
		return
		//add error handling
	}
	fmt.Println("Threads()_#2")
	for rows.Next() {
		fmt.Println("Threads()_row")
		conv := Thread{}
		if err=rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err !=nil {
			return
			//add error handling
		}
		threads=append(threads, conv)
	}
	rows.Close()
	fmt.Println("Threads()_#3")
	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}


// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

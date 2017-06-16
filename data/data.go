package data

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"

)
const (
	DRIVER_NAME = "mysql"
	DATA_SOURCE_NAME = "ayong:ayong@tcp(127.0.0.1:3306)/test?parseTime=true"
)

var Db *sql.DB


func init() {
	var err error
	fmt.Println("init_#1")
	Db, err = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if err !=nil {
		fmt.Println("init_err1")
		log.Println(err)
		log.Fatal(err)
	}
	if err=Db.Ping(); err !=nil {
		fmt.Println("init_err2")
		log.Println(err)
		log.Fatal(err)
	}
	fmt.Println("init_#2")
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

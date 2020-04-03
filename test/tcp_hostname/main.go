package main

import (
	"blazetunnel/db"
	"log"
)

func main() {

	err := (&db.App{
		Appname:  "dsadas",
		Password: "dasd",
	}).CreateApp()

	log.Println("Createapp", err)
	err = (&db.App{
		Appname:  "dsadas",
		Password: "dasddsada",
	}).CreateApp()

	log.Println("Createapp", err)

	ans := (&db.App{
		Appname:  "dasdsa",
		Password: "Dsads",
	}).Authenticate()

	log.Println(ans)

	(&db.App{
		Appname: "dasdsa",
	}).CreateApp()
	ans = (&db.App{
		Appname: "dasdsa",
	}).Authenticate()

	log.Println(ans)

}

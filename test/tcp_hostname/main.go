package main

import (
	"blazetunnel/db"
	"log"
)

func main() {

	err := (&db.App{
		Appname:  "dasdssa",
		Password: "",
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

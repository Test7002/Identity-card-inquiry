package main

import (
	"log"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLoge := log.New(os.Stdout, "INFO\t", log.LstdFlags|log.Lmicroseconds)
	errorLoge := log.New(os.Stderr, "ERROR\t", log.LstdFlags|log.Lmicroseconds)

	app := application{
		errorLog: errorLoge,
		infoLog:  infoLoge,
	}
	r := app.routers()
	err := r.Run(":5555")
	if err != nil {
		errorLoge.Fatal(err)
		return
	}
	infoLoge.Printf("Starting server on %s", "")
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	appClient "mysms4you.com/app/src/repo/das/client"
	"mysms4you.com/app/src/repo/das/constants"
	"mysms4you.com/app/src/repo/route"
)

func main() {

	ctx := context.Background()

	var onLocal bool = true
	if onLocal {
		os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:"+constants.Emulator_Port)
	}

	/* //Init Firebase app with Service account Config File on Non-Google Platform
	sa := option.WithCredentialsFile("./path-to-account.json")
	app, err := firebase.NewApp(ctx, nil, sa) */

	//Init Firebase app with Project-Id on Google platform
	conf := &firebase.Config{ProjectID: constants.ProjectID}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
	}

	//Create Firestore DB Object
	_, err = appClient.CreateFirestoreDB(app)
	if err != nil {
		log.Fatalf("newFirestoreDB: %v", err)
	}
	// set up the HTTP server
	router := route.Handler()

	fmt.Println("The app server is on tap now: http://localhost:8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}

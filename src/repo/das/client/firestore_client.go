package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"mysms4you.com/app/src/repo/das/constants"
)

type firestoreDB struct {
	Client     *firestore.Client
	Collection string
}

var db *firestoreDB

func CreateFirestoreDB(app *firebase.App) (*firestoreDB, error) {
	ctx := context.Background()

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("firestoredb: could not create Firestore Client : %v", err)
	}
	// Verify that we can communicate and authenticate with the Firestore
	// service.
	err = client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("firestoredb: could not connect: %v", err)
	}
	db = &firestoreDB{
		Client:     client,
		Collection: constants.DB_Name,
	}
	return db, nil
}

func GetClient() (*firestoreDB, error) {
	if db == nil {
		return nil, fmt.Errorf("DB is null")
	}
	return db, nil
}

// Close closes the database.
func (db *firestoreDB) Close(context.Context) error {
	return db.Client.Close()
}

package access

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"mysms4you.com/app/src/models"
	appClient "mysms4you.com/app/src/repo/das/client"
	"mysms4you.com/app/src/repo/das/constants"
)

func GetHomeList(ctx context.Context) (models.StatusList, error) {
	statusList := make([]*models.Status, 0)
	response := models.StatusList{Page_id: "0", Total_Records: "0", Results: statusList}

	collection, err := getCollection(constants.SampleCollection)
	if err != nil {
		return response, err
	}

	iter := collection.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return response, fmt.Errorf("homeDAS: could not list status: %v", err)
		}

		s := &models.Status{}
		doc.DataTo(s)
		statusList = append(statusList, s)

	}

	return models.StatusList{Page_id: "0", Total_Records: "1000", Results: statusList}, nil
}

func getCollection(collectionName string) (*firestore.CollectionRef, error) {
	client, err := appClient.GetClient()
	if err != nil {
		return nil, fmt.Errorf("Could not get Client: %v", err)
	}
	return client.Client.Collection(collectionName), nil
}

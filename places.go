package main

import (
	"context"
	"net/http"
)

func placesIndex(w http.ResponseWriter, r *http.Request) {

	client := new(dbHandler).connect()
	defer client.Disconnect(context.TODO())

}

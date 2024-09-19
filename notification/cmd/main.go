package cmd

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func main() {
	opt := option.WithCredentialsFile("../../mikronet-6cf58-firebase-adminsdk-prfvh-573a5582a6.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Messaging(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

}

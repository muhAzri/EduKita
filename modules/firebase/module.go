package firebase

import (
	"EduKita/modules/firebase/middleware"
	"context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func StartFirebaseModule() middleware.FirebaseMiddleware {

	opt := option.WithCredentialsFile("edu-kita-firebase-admin.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		panic(err)
	}

	client, err := app.Auth(context.Background())

	if err != nil {
		panic(err)
	}

	middleware := middleware.FirebaseMiddleware{AuthClient: *client}

	return middleware
}

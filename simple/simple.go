package simple

import (
	"context"
	"fmt"
	"log"

	"github.com/steebchen/photon-example/db"
)

/*
query Raw {
	findOneUser(where: { id: 420 }) {
		id
		firstName
		lastName
	}
}
*/

var userID = 420

func Raw(ctx context.Context, client *db.PrismaClient) db.UserModel {
	var result []db.UserModel
	err := client.Raw(`SELECT * FROM "User" WHERE id = $1 LIMIT 1`, userID).Exec(ctx, &result)
	if err != nil {
		panic(err)
	}
	if len(result) > 0 {
		if result[0].ID != userID {
			log.Fatalf("expect %d, actual %d", userID, result[0].ID)
		}
		return result[0]
	}
	panic(fmt.Errorf("didn't find a result"))
}

func Query(ctx context.Context, client *db.PrismaClient) db.UserModel {
	user, err := client.User.FindOne(db.User.ID.Equals(userID)).Exec(ctx)
	if err != nil {
		panic(err)
	}
	if user.ID != userID {
		log.Fatalf("expect %d, actual %d", userID, user.ID)
	}
	return user
}

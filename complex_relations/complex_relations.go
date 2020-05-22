package complex_relations

import (
	"context"
	"log"

	"github.com/steebchen/photon-example/db"
)

/*
query SimplerComplex {
	findOneComment(where: { id: 420 }) {
		content
		likes(where: { id_gt: 10000 }) {
			id
			post {
				id
				content
			}
		}
	}
}
*/

const UserID = 4301
const LikeIDMin = 400
const CommentIDResult = 1

func Raw(ctx context.Context, client *db.PrismaClient) interface{} {
	var result interface{}
	// language=PostgreSQL
	err := client.Raw(`
		SELECT
				"Comment"."id" AS "commentId",
				"Comment"."content",
				"Like"."id" AS "likeId",
				"Post"."id" AS "postId",
				"Post"."content" AS "postContent"
		FROM "Comment"
			LEFT JOIN "Like" ON "Comment"."id" = "Like"."comment"
			LEFT JOIN "Post" ON "Like"."post" = "Post"."id"
		WHERE
			"Comment"."author" = $1
			AND
			"Like"."id" > $2
	`, UserID, LikeIDMin).Exec(ctx, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("result %+v", result)
	return result
}

func Query(ctx context.Context, client *db.PrismaClient) []db.CommentModel {
	comment, err := client.Comment.FindMany(
		db.Comment.ID.Equals(UserID),
		db.Comment.VUser.Where(
			db.User.ID.Equals(UserID),
		),
		db.Comment.VLike.Some(
			db.Like.ID.GT(LikeIDMin),
		),
	).With(
		db.Comment.VLike.Fetch().With(
			db.Like.VPost.Fetch(),
		),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	// if comment.ID != CommentIDResult {
	// 	log.Fatalf("expect %d, actual %d", CommentIDResult, comment.ID)
	// }
	return comment
}

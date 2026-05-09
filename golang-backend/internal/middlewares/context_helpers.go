package middlewares

import "context"

type contextKey string

const userIDKey contextKey = "userID"

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}

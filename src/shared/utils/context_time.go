package utils

import (
	"context"
)

func GetContext() context.Context {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	return ctx
}

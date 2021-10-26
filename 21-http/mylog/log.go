package mylog

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int64
const k key = 66

func Println(ctx context.Context, msg string) {

	id, ok := ctx.Value(k).(int64)
	if !ok {
		log.Println("unknown", msg)
	}
	log.Println(id, msg)
}

func Middleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {


		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, k, id)
		next(w, r.WithContext(ctx))

	}
}

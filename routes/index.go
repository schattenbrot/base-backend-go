package routes

import (
	"context"
	"fmt"
	"net/http"
)

// HTTP middleware setting a value on the request context
func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", "123")
		fmt.Println("Middleware fired")

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func With() []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{MyMiddleware}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "test")))
}

func Test() {
	fmt.Println("test")
}

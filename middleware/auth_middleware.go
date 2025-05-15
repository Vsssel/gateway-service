package middleware

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if (strings.Contains(r.URL.Path, "/me") || strings.Contains(r.URL.Path, "/users") ) {
     
      authHeader := r.Header.Get("Authorization")
      if authHeader == "" {
        http.Error(w, "Authorization header missing", http.StatusUnauthorized)
        return
      }
    }

    next.ServeHTTP(w, r)
  })
}

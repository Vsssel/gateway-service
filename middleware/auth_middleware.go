package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/register" || r.URL.Path == "/login" {
      next.ServeHTTP(w, r)
      return
    }

    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
      http.Error(w, "Authorization header missing", http.StatusUnauthorized)
      return
    }


    next.ServeHTTP(w, r)
  })
}

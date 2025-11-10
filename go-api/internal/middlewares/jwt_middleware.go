package middlewares

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(secret string) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            auth := r.Header.Get("Authorization")
            if auth == "" {
                http.Error(w, "missing auth", http.StatusUnauthorized)
                return
            }
            parts := strings.Fields(auth)
            if len(parts) != 2 || parts[0] != "Bearer" {
                http.Error(w, "invalid auth", http.StatusUnauthorized)
                return
            }
            tokenStr := parts[1]
            token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
                return []byte(secret), nil
            })
            if err != nil || !token.Valid {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }
            // optionally set user id in context
            next.ServeHTTP(w, r)
        })
    }
}

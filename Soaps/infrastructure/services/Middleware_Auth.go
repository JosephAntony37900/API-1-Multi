package service

import (
	"net/http"
	"strings"
	"fmt" 
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		fmt.Println("Authorization header recibido:", authHeader)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token not found"})
			return
		}
		fmt.Println("Token extraído:", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("Error al parsear el token:", err) 
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		fmt.Println("Token validado:", token)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Error al obtener claims del token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		userID, exists := claims["userId"]
if !exists {
    fmt.Println("user_id no encontrado en los claims")
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id no encontrado"})
    return
}

userIDInt, ok := userID.(float64) 
if ok {
    c.Set("userID", int(userIDInt)) 
} else {
    fmt.Println("Error al convertir userID a int")
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userID no es válido"})
    return
}

		c.Set("userID", userID) 
		c.Next()
	}
}


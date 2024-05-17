# gin-auth

This repository is to use gin to implementation authentication logic

## package install

```shell
go get -u golang.org/x/crypto/bcrypt 
```

## mysql setup

```
services:
  db: 
    image: mysql:8
    container_name: mysql
    restart: always
    environment:
      MYSQL_DATABASE: "${MYSQL_DATABASE}"
      MYSQL_USER: "${MYSQL_USER}"
      MYSQL_PASSWORD: "${MYSQL_PASSWORD}"
      MYSQL_ROOT_PASSWORD: "${MYSQL_ROOT_PASSWORD}"
    ports:
      - "${PORT}:${PORT}"
    expose:
      - "${PORT}"
    volumes:
      - ./data:/var/lib/mysql
    logging:
      driver: "json-file"
      options: 
        max-size: "1k"
        max-file: "3"
```

## setup middleware for jwt token check

```golang
func AuthMiddleware(c *gin.Context) {
	// Redirect the cookie from the request
	tokenString, err := c.Cookie("Auth")
	if err != nil {
		c.HTML(
			http.StatusUnauthorized,
			"errors/error.tpl",
			gin.H{"error": "No auth token"},
		)
		c.Abort()
		return
	}
	// Extract the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		hmacSampleSecret := config.C.JWT_SIGN_SECRET
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		c.HTML(
			http.StatusUnauthorized,
			"errors/error.tpl",
			gin.H{"error": "Failed to parse JWT token"},
		)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.HTML(
			http.StatusUnauthorized,
			"errors/error.tpl",
			gin.H{"error": "JWT Claims failed"},
		)
		c.Abort()
		return
	}
	// Check expiry of the token
	if claims["ttl"].(float64) < float64(time.Now().Unix()) {
		c.HTML(
			http.StatusUnauthorized,
			"errors/error.tpl",
			gin.H{"error": "JWT token expired"},
		)
		c.Abort()
		return
	}
	// Extract user from token
	var user models.User
	models.DB.Where("id = ?", claims["userID"]).First(&user)
	if user.ID == 0 {
		c.HTML(http.StatusUnauthorized,
			"errors/error.tpl",
			gin.H{"error": "Could not find the user!"})
		c.Abort()
		return
	}
  // set back to context
	c.Set("user", user)
	// go to the next chain
	c.Next()
}
```
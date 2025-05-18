package auth

var jwtSecret string
var hashSalt string

// Init initializes the JWT secret key and hash salt.
func Init(jwtSecretKey string, hashSaltValue string) {
	jwtSecret = jwtSecretKey
	hashSalt = hashSaltValue
}

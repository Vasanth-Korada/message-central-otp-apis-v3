package auth

import "sync"

var (
	authToken string
	mutex     sync.Mutex
)

// SetAuthToken sets the authToken
func SetAuthToken(token string) {
	mutex.Lock()
	defer mutex.Unlock()
	authToken = token
}

// GetAuthToken returns the authToken
func GetAuthToken() string {
	mutex.Lock()
	defer mutex.Unlock()
	return authToken
}

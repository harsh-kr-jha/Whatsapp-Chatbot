package state

import "sync"

// We'll store user states in memory for now
var (
	userStates   = make(map[string]string)
	userJobPosts = make(map[string]map[string]string)
	mu           sync.RWMutex
)

// GetUserState returns the current state for the user
func GetUserState(userID string) string {
	mu.RLock()
	defer mu.RUnlock()
	return userStates[userID]
}

// SetUserState updates the state for the user
func SetUserState(userID, state string) {
	mu.Lock()
	defer mu.Unlock()
	userStates[userID] = state
}

// GetUserJobData returns the job data map for the user
func GetUserJobData(userID string) map[string]string {
	mu.Lock()
	defer mu.Unlock()
	if userJobPosts[userID] == nil {
		userJobPosts[userID] = make(map[string]string)
	}
	return userJobPosts[userID]
}

// ResetUserState clears the state & job data for a user
func ResetUserState(userID string) {
	mu.Lock()
	defer mu.Unlock()
	delete(userStates, userID)
	delete(userJobPosts, userID)
}

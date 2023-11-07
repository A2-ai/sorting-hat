package users

import (
	"log/slog"
	"os"
	"os/user"
	"sync"
)

type User struct {
	Username string
	GroupIDs []string
	User     *user.User
}

type Users struct {
	Users map[string]User
	mu    sync.RWMutex
}

func NewUsers() *Users {
	return &Users{
		Users: make(map[string]User),
	}
}

func (u *Users) Add(username string, user User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Users[username] = user
}

func getPotentialUsersByDirName(dir string) ([]string, error) {
	var potentialUsers []string
	dirs, err := os.ReadDir(dir)
	if err != nil {
		return potentialUsers, err
	}
	for _, d := range dirs {
		if d.IsDir() {
			potentialUsers = append(potentialUsers, d.Name())
		}
	}
	return potentialUsers, nil
}

func lookupUser(u string, logger *slog.Logger) (User, error) {
	user, err := user.Lookup(u)
	if err != nil {
		return User{}, err
	}
	grpIds, err := user.GroupIds()
	if err != nil {
		logger.Warn("user.GroupIds failed for potential user", "user", u)
	}
	return User{
		Username: user.Username,
		GroupIDs: grpIds,
		User:     user,
	}, nil
}

// GetPotentialUsersByDirName returns a list of potential users based on the
// directory name.  This allows us to identify users potentially using the
// system through the nature of them having a directory (aka a home dir)
func GetPotentialUsersByDirName(dir string, logger *slog.Logger, concurrency int) (*Users, error) {
	users := NewUsers()
	potentialUsers, err := getPotentialUsersByDirName(dir)
	if err != nil {
		return users, err
	}
	if concurrency == 0 {
		concurrency = 3
	}
	// bound how many goroutines we can run at once
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	for _, u := range potentialUsers {
		wg.Add(1)
		go func(usr string) {
			sem <- struct{}{}
			defer func() {
				wg.Done()
				<-sem
			}()
			user, err := lookupUser(usr, logger)
			if err != nil {
				logger.Warn("failed to lookup user", "user", usr)
				return
			}
			users.Add(user.Username, user)
		}(u)
	}
	wg.Wait()
	return users, nil
}

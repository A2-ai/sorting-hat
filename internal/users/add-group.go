package users

import (
	"fmt"
	"os/exec"
)

func addUserToGroup(username, group string) error {
	cmd := exec.Command("usermod", "-aG", group, username)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error adding user to group: %w", err)
	}
	return nil
}

package commands

import (
	"fmt"
	"os"

	"github.com/daedaleanai/git-ticket/cache"
	"github.com/daedaleanai/git-ticket/input"
	"github.com/daedaleanai/git-ticket/util/interrupt"
	"github.com/spf13/cobra"
)

func runUserCreate(cmd *cobra.Command, args []string) error {
	backend, err := cache.NewRepoCache(repo)
	if err != nil {
		return err
	}
	defer backend.Close()
	interrupt.RegisterCleaner(backend.Close)

	preName, err := backend.GetUserName()
	if err != nil {
		return err
	}

	name, err := input.PromptDefault("Name", "name", preName, input.Required)
	if err != nil {
		return err
	}

	preEmail, err := backend.GetUserEmail()
	if err != nil {
		return err
	}

	email, err := input.PromptDefault("Email", "email", preEmail, input.Required)
	if err != nil {
		return err
	}

	avatarURL, err := input.Prompt("Avatar URL", "avatar")
	if err != nil {
		return err
	}

	id, err := backend.NewIdentityRaw(name, email, "", avatarURL, nil)
	if err != nil {
		return err
	}

	err = id.CommitAsNeeded()
	if err != nil {
		return err
	}

	set, err := backend.IsUserIdentitySet()
	if err != nil {
		return err
	}

	if !set {
		err = backend.SetUserIdentity(id)
		if err != nil {
			return err
		}
	}

	_, _ = fmt.Fprintln(os.Stderr)
	fmt.Println(id.Id())

	return nil
}

var userCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new identity.",
	PreRunE: loadRepo,
	RunE:    runUserCreate,
}

func init() {
	userCmd.AddCommand(userCreateCmd)
	userCreateCmd.Flags().SortFlags = false
}

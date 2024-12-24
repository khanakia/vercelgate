package main

import (
	"app/gen/ent"
	"app/gen/ent/team"
	"app/pkg/constants"
	"app/pkg/entcfn"
	"app/pkg/entdb"
	"app/pkg/utils"
	"app/pkg/vercelfn"
	"app/pkg/vercelutil"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var version = "dev"

func main() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(switchTeamCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:     "vercelgate",
	Version: version,
	Short:   "Make vercel cli more powerful by adding the ability to switch between multiple accounts.",
	Long:    `You can swithc between multiple accounts without having relogin and logout.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run command `vercelgate --help` for more information`")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Run this command very first time",
	Run: func(cmd *cobra.Command, args []string) {
		globalPath, err := vercelutil.GetGlobalPathConfig()
		if err != nil {
			log.Fatal(err)

			return
		}

		err = utils.IsFileExists(filepath.Join(globalPath, constants.DB_FILE_NAME))
		if err == nil {
			fmt.Println("was initialized already")
			return
		}

		err = entcfn.Migrate()
		if err != nil {
			log.Fatal(err)

			return
		}
		fmt.Println("vercelgate initialized successfully")
	},
}

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch between account",
	Run: func(cmd *cobra.Command, args []string) {

		SwitchCmd(false)
	},
}

var switchTeamCmd = &cobra.Command{
	Use:   "switchteam",
	Short: "Switch between account and teams",
	Run: func(cmd *cobra.Command, args []string) {
		SwitchCmd(true)
	},
}

func SwitchCmd(switchTeam bool) error {
	user, err := promptGetUser()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = vercelutil.SetAuthToken(user.Token)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Switched to user %s\n", user.Name)

	if switchTeam {
		team, err := promptGetTeam(user.ID)
		if err != nil {
			log.Fatal(err)
			return err
		}

		err = vercelutil.SetCurrentTeam(team.ID)
		if err != nil {
			log.Fatal(err)
			return err
		}

		fmt.Printf("Switched to team %s\n", team.Name)
	} else {
		vercelutil.DeleteCurrentTeam()
	}

	return nil
}

func promptGetTeam(userID string) (*ent.Team, error) {
	ctx := context.Background()

	items, err := entdb.Client().Team.Query().Where(team.UserID(userID)).All(ctx)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	itemsList := []string{}

	for _, user := range items {
		itemsList = append(itemsList, user.Name)
	}

	prompt := promptui.Select{
		Label: "Select Team",
		Items: itemsList,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	return items[index], nil
}

func promptGetUser() (*ent.User, error) {
	ctx := context.Background()

	users, err := entdb.Client().User.Query().All(ctx)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	usersList := []string{}

	for _, user := range users {
		name := user.Name
		if len(name) == 0 {
			name = user.Username
		}
		usersList = append(usersList, fmt.Sprintf("%s (%s)", name, user.Email))
	}

	prompt := promptui.Select{
		Label: "Select Account",
		Items: usersList,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	return users[index], nil
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync current logged in account",
	Run: func(cmd *cobra.Command, args []string) {
		err := vercelfn.SyncAuthJson()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("synced successfully")
	},
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Run this to add new vercel client account",
	Run: func(cmd *cobra.Command, args []string) {
		err := NewAccountCmd()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("you can now add new account using `vercel login` and then run `vercelgate sync` again")
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset vercelgate state and will delete all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		err := ResetCmd()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("state reset was successful")
	},
}

func NewAccountCmd() error {
	filePath, err := vercelutil.AuthJsonFile()
	if err != nil {
		return err
	}

	err = os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to remove auth.json file: %w", err)
	}
	return nil
}

func ResetCmd() error {
	ctx := context.Background()
	_, err := entdb.Client().User.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = entdb.Client().Team.Delete().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

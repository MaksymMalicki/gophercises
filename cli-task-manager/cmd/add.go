package clitaskmanger

import (
	"fmt"
	"strings"

	"github.com/Maksymmalicki/gophercises/clitaskmanager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "add a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print("You must provide the name for the task!")
			return
		}
		taskName := strings.Join(args, " ")
		db.WriteToDB([]byte(taskName))
		fmt.Printf(`Added "%s" to your task list.`, taskName)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

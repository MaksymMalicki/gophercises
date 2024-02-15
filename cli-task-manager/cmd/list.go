package clitaskmanger

import (
	"fmt"
	"log"

	"github.com/Maksymmalicki/gophercises/clitaskmanager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "list all the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadFromDB()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Print("You don't have any tasks!")
			return
		}
		for index, task := range tasks {
			fmt.Printf("%d) %s\n", index+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

package clitaskmanger

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Maksymmalicki/gophercises/clitaskmanager/db"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"d"},
	Short:   "mark task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.ReadFromDB()
		if len(tasks) == 0 {
			fmt.Println("You cannot delete tasks, because there are none!")
		}
		for _, id := range ids {
			if id > len(tasks) || id < 0 {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.RemoveFromDB(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}

		}
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

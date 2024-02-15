package main

import (
	clitaskmanger "github.com/Maksymmalicki/gophercises/clitaskmanager/cmd"
	"github.com/Maksymmalicki/gophercises/clitaskmanager/db"
)

func main() {
	db.InitDB("tasks.db")
	defer db.BoltDB.Close()

	clitaskmanger.Execute()
}

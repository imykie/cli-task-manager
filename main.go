package main

import (
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"task/cmd"
	"task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	cmd.Must(db.Init(dbPath))
	cmd.Execute()
}

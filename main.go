package main

import (
	"taskmanager/internal/tasks"

	"taskmanager/cmd"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "taskmanager_user:taskmanager_user_password@tcp(127.0.0.1:13306)/task"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&tasks.Task{})

	cmd.Execute(db)
}

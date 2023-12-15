package cmd

import (
	"fmt"
	"os"
	"taskmanager/internal/tasks"
	"taskmanager/utils"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Execute(db *gorm.DB) {
	var cmdAdd = &cobra.Command{
		Use:   "add [sting]",
		Short: "Command to add a task",
		Long:  "This command add a task to the task list.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var task = tasks.Task{
				Name: string.Join(args, " "),
			}
			newTask := tasks.Add(db, task)

			fmt.Printf(`Tarea creada com exito: %s(%d)`, newTask.Name, task.ID)
		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Command to list all tasks",
		Long:  "This command list all tasks.",
		Run: func(cmd *cobra.Command, Args []string) {
			tasksList := tasks.GetAll(db)
			if len(tasksList) == 0 {
				fmt.Println("No hay tareas por hacer")
				return
			}
			for i := 0; i < len(tasksList); i++ {
				fmt.Printf(
					"%d = %s (%t)\n",
					tasksList[i].ID,
					tasksList[i].Name,
					tasksList[i].Completed,
				)
			}
		},
	}

	var cmdDetail = &cobra.Command{
		Use:   "detail [id]",
		Short: "Command to show a task detail",
		Long:  "this command show a task detail.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])
			task := tasks.GetByID(db, id)
			fmt.Printf("ID: %d \nNombre: %s \nDescripción: %s \nEstado: %t \n",
				task.ID,
				task.Name,
				task.Description,
				task.Completed)
		},
	}

	var cmdCompleted = &cobra.Command{
		Use:   "completed [id]",
		Short: "Command to mark a task as completed",
		Long:  "this command mark a task as completed",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])
			taskEdited := tasks.GetByID(db, id)
			taskEdited.Completed = true
			task := tasks.UpdateByID(db, id, *taskEdited)
			fmt.Printf("La tarea com ID: %d se ha marcado como completada",
				task.ID)
		},
	}

	var cmdDelete = &cobra.Command{
		Use:   "delete [id]",
		Short: "Command to delete a task",
		Long:  "this command delete a task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])
			tasks.DeleteByID(db, id)
			fmt.Printf("La tarea com ID: %d se ha eliminado", id)
		},
	}

	var cmdUpdate = &cobra.Command{
		Use:   "update [id] [string] [string]",
		Short: "Command to update a task",
		Long:  "this command update a task.",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])
			var tasksEdited = tasks.GetByID(db, id)
			if args[1] == "name" {
				tasksEdited.Name = args[2]
			} else if args[1] == "description" {
				tasksEdited.Description = args[2]
			} else {
				panic("Invalid argument")
			}
			task := tasks.UpdateByID(db, id, *tasksEdited)
			fmt.Printf("Actualización realizada con exito \n\n ID: %d \nNombre: %s \nDescripción: %s \nEstado: %t \n",
				task.ID,
				task.Name,
				task.Description,
				task.Completed)
		},
	}

	var rootCmd = &cobra.Command{Use: "taskmanager"}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var rootCmd = &cobra.Command{Use: "taskmanager"}
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdDetail)
	rootCmd.AddCommand(cmdUpdate)
	rootCmd.AddCommand(cmdCompleted)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

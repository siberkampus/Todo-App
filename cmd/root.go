/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	todoapp "todoapp/todo"

	"github.com/spf13/cobra"
)
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ToDo-App",
	Short: "Simple ToDo App",
	Long:  `You can plan your activities with TODOapp`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: Generate,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todoapp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("add", "a", "", "Add new task")
	rootCmd.Flags().IntP("delete", "d", -1, "Delete task")
	rootCmd.Flags().IntP("complete", "c", -1, "Complete task")
	rootCmd.Flags().IntP("incomplete", "n", -1, "Make task incomplete")
	rootCmd.Flags().BoolP("list", "l", false, "list all tasks")
}

func Generate(cmd *cobra.Command, args []string) {
	var item todoapp.Items
	add, _ := cmd.Flags().GetString("add")
	delete, _ := cmd.Flags().GetInt("delete")
	complete, _ := cmd.Flags().GetInt("complete")
	list, _ := cmd.Flags().GetBool("list")
	incomplete, _ := cmd.Flags().GetInt("incomplete")
	err := todoapp.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	case len(add) > 0:
		item.Add(add)
		err := todoapp.Save()
		if err != nil {
			log.Fatal("Save error",err)
		}
	case delete >= 0:
		err := item.Delete(delete)
		if err != nil {
			log.Fatal("Delete error",err)
		}
		err = todoapp.Save()
		if err != nil {
			log.Fatal("Save error",err)
		}
	case complete >= 0:
		err := item.Complete(complete)
		if err != nil {
			log.Fatal("Complete error",err)
		}
		err = todoapp.Save()
		if err != nil {
			log.Fatal("Save error",err)
		}
	case list:
		err = todoapp.Save()
		if err != nil {
			log.Fatal("Save error",err)
		}
		todoapp.List()
		
	case incomplete > 0:
		err := item.InComplete(incomplete)
		if err != nil {
			log.Fatal("Incomplete error",err)
		}
		err = todoapp.Save()
		if err != nil {
			log.Fatal("Save error",err)
		}
	}

}

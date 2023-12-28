package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var version bool
	var arg1, arg2, arg3 string
	var rootCmd = &cobra.Command{
		Use:   "root [sub]",
		Short: "root command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
			if version {
				fmt.Printf("Version:1.0\n")
			}
		},
	}

	var rootCmd1 = &cobra.Command{
		Use:   "root1 [sub]",
		Short: "root1 command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd1 Run with args: %v\n", args)
			if version {
				fmt.Printf("Version:1.0\n")
			}
		},
	}
	// 添加命令行选项
	flags := rootCmd.Flags()
	flags.BoolVarP(&version, "version", "v", false, "Print version information and quit")

	flags1 := rootCmd1.Flags()
	flags1.StringVarP(&arg1, "aaa", "a", "", "Test arg1")
	flags1.StringVarP(&arg2, "bbb", "b", "", "Test arg2")
	flags1.StringVarP(&arg3, "ccc", "c", "", "Test arg3")
	// 添加子命令
	rootCmd.AddCommand(rootCmd1)
	rootCmd.Execute()
}

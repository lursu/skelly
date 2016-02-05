// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gopath string

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "skelly [sub-command]",
	Short: "build the skeliton of a golang application",
	Long: `skelly is used to build new go projects templates.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(addCommands)
	initCommonFlags(RootCmd)
	initConfig(RootCmd)
}

func addCommands() {
	RootCmd.AddCommand(buildCmd)
	RootCmd.AddCommand(configCmd)
}

// Set up any flags that are going to be needed for any command
func initCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&basePath, "project-root", "r", "root for project", "root project dirrectory relative to $GOPATH/src")
	cmd.Flags().StringVarP(&author, "author", "a", "your name", "author name for copyright")
	cmd.Flags().StringVarP(&email, "email", "e", "your email", "your email adress to show up on the files")
	cmd.Flags().BoolVarP(&license, "license", "l", false, "name of the license to use")
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command) {
	viper.SetConfigName(".skelly") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.SetConfigType("json")
	viper.AutomaticEnv()                  // read in environment variables that match

	fmt.Println("got here")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Make sure that go src var is set
	gopath = viper.GetString("gopath")
	if len(gopath) <= 0 {
		gopath = joinPath(os.Getenv("GOPATH"), "src")
		viper.Set("gopath", gopath)
	}

	if cmd.Flags().Lookup("project-root").Changed {
		viper.Set("project-root", basePath)
	}
	if cmd.Flags().Lookup("author").Changed {
		fmt.Println("adding author")
		viper.Set("author", author)
	}
	if cmd.Flags().Lookup("email").Changed {
		viper.Set("email", email)
	}
	fmt.Println(email)
	if cmd.Flags().Lookup("license").Changed {
		viper.Set("license", license)
	}
}

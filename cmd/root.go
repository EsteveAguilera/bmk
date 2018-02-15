package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bmk",
	Short: "A CLI bookmark manager",
	Long:  `bmk is a CLI bookmark manager.`,
}

var commandColor = color.New(color.FgGreen)
var linkColor = color.New(color.FgYellow).Add(color.Underline)
var nameColor = color.New(color.FgCyan).Add(color.Bold)
var idColor = color.New(color.FgGreen).Add(color.Bold)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bmk.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bmk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bmk")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func printBookmark(bmk bookmark) {
	idString := strconv.Itoa(bmk.id)
	idColor.Print("[" + idString + "]")
	nameColor.Print(" " + bmk.name + " ")
	linkColor.Println(bmk.url)
}

type bookmark struct {
	id   int
	url  string
	name string
	tag  string
}

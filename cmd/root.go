package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
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
}

func initConfig() {

	viper.SetConfigFile("./bmk.yaml")
	viper.SetConfigType("yaml")

	viper.SetDefault("color", true)

	// If no config file is found, create it!
	if err := viper.ReadInConfig(); err != nil {
		viper.WriteConfig()
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

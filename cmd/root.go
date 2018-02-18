package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
var bookmarkFileKey = "filename"

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

	// Use colored output by default
	viper.SetDefault("nocolor", false)
	viper.SetDefault(bookmarkFileKey, "bookmarks.json")

	if err := viper.ReadInConfig(); err != nil {
		viper.WriteConfig()
	}

	// Load color output config
	color.NoColor = viper.GetBool("nocolor")
}

func printBookmark(bmk bookmark) {
	idString := strconv.Itoa(bmk.ID)
	idColor.Print("[" + idString + "]")
	nameColor.Print(" " + bmk.Name + " ")
	linkColor.Print(bmk.URL)
}

func loadBookmarks() []bookmark {
	var bookmarkList []bookmark
	data, err := ioutil.ReadFile(viper.GetString(bookmarkFileKey))
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal(data, &bookmarkList)
		if err != nil {
			fmt.Println("No bookmarks found!")
		}
	}
	return bookmarkList
}

func saveBookmarks(bmkList []bookmark) {
	b, _ := json.Marshal(bmkList)
	_ = ioutil.WriteFile(viper.GetString(bookmarkFileKey), b, 0644)
}

type bookmark struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

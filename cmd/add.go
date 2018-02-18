package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bookmark",
	Long:  `Add a new bookmark`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter name: ")
		name, _ := reader.ReadString('\n')
		fmt.Print("Enter URL: ")
		url, _ := reader.ReadString('\n')
		fmt.Print("Enter tag: ")
		tag, _ := reader.ReadString('\n')
		//newBmk := bookmark{42, url, name, tag}

		bookmarkList := loadBookmarks()
		bookmarkList = append(bookmarkList, bookmark{ID: len(bookmarkList) + 1, URL: url, Name: name, Tag: tag})
		for _, item := range bookmarkList {
			printBookmark(item)

		}
		saveBookmarks(bookmarkList)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

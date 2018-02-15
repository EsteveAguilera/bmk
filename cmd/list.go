package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the saved bookmarks",
	Long:  `List all the saved bookmarks`,
	Run: func(cmd *cobra.Command, args []string) {

		bookmarkList := []bookmark{}
		bookmarkList = append(bookmarkList, bookmark{id: 1, url: "http://google.com", name: "Google", tag: "google"})
		bookmarkList = append(bookmarkList, bookmark{id: 2, url: "http://gmail.com", name: "Gmail", tag: "gmail"})

		for _, item := range bookmarkList {
			printBookmark(item)

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

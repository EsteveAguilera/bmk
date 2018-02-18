package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the saved bookmarks",
	Long:  `List all the saved bookmarks`,
	Run: func(cmd *cobra.Command, args []string) {
		bookmarkList := loadBookmarks()
		for _, item := range bookmarkList {
			printBookmark(item)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

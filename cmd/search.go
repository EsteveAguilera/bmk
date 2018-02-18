package cmd

import (
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search bookmarks for name or tags",
	Long:  `Search bookmarks for name or tags`,
	Run: func(cmd *cobra.Command, args []string) {

		bookmarkList := loadBookmarks()
		var someResult bool

		if len(args) > 0 {
			commandColor.Print("Searching bookmark(s) with tag(s): ")
			nameColor.Println(args[0])

			for _, item := range bookmarkList {
				if strings.Contains(item.Tag, args[0]) {
					printBookmark(item)
					someResult = true
				}
			}

			if !someResult {
				color.Red("No bookmarks with this tag(s)")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

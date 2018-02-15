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

		bookmarkList := []bookmark{}
		bookmarkList = append(bookmarkList, bookmark{id: 1, url: "http://google.com", name: "Google", tag: "google"})
		bookmarkList = append(bookmarkList, bookmark{id: 2, url: "http://gmail.com", name: "Gmail", tag: "gmail"})

		var someResult bool

		if len(args) > 0 {
			commandColor.Print("Searching bookmark(s) with tag(s): ")
			nameColor.Println(args[0])

			for _, item := range bookmarkList {
				if strings.Contains(item.tag, args[0]) {
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

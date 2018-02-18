package cmd

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a bookmark",
	Long: `Open a bookmark using the id value

bmk open 42`,
	Run: func(cmd *cobra.Command, args []string) {
		bookmarkList := loadBookmarks()
		var err error
		var id int

		if id, err = strconv.Atoi(args[0]); err != nil {
			panic(err)
		}

		var urlToOpen string

		for _, item := range bookmarkList {
			if id == item.ID {
				urlToOpen = item.URL
			}
		}

		err = exec.Command("open", urlToOpen).Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

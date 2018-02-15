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
		bookmarkList := []bookmark{}
		bookmarkList = append(bookmarkList, bookmark{id: 1, url: "http://google.com", name: "Google", tag: "google"})
		bookmarkList = append(bookmarkList, bookmark{id: 2, url: "http://gmail.com", name: "Gmail", tag: "gmail"})

		var err error
		var id int

		if id, err = strconv.Atoi(args[0]); err != nil {
			panic(err)
		}

		var urlToOpen string

		for _, item := range bookmarkList {
			if id == item.id {
				urlToOpen = item.url
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

package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
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
			fmt.Println("Not a valid command")
			return
		}

		if id > len(bookmarkList) {
			fmt.Println("No bookmark with this ID")
			return
		}

		openUrl(bookmarkList[(id - 1)].URL)

	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func openUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

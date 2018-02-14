package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type bookmark struct {
	id   int
	url  string
	name string
	tag  string
}

var (
	someResult   bool
	bookmarkList []bookmark
	commandColor = color.New(color.FgGreen)
	linkColor    = color.New(color.FgYellow).Add(color.Underline)
	nameColor    = color.New(color.FgCyan).Add(color.Bold)
	idColor      = color.New(color.FgGreen).Add(color.Bold)
)

func init() {}

func main() {

	switch os.Args[1] {
	case "search":
		if len(os.Args) > 2 {
			searchBookmark(os.Args[2])
		} else {
			printUsage()
		}
	case "add":
		addBookmark()
	case "open":
		if len(os.Args) > 2 {
			s := os.Args[2]
			id, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			openbrowser(id)
		}
	case "help":
		printUsage()
	default:
		printUsage()
	}
}

func searchBookmark(text string) {
	bookmarkList := []bookmark{}
	bookmarkList = append(bookmarkList, bookmark{id: 1, url: "http://google.com", name: "Google", tag: "google"})
	bookmarkList = append(bookmarkList, bookmark{id: 2, url: "http://gmail.com", name: "Gmail", tag: "gmail"})

	commandColor.Print("Searching bookmark(s) with tag(s): ")
	nameColor.Println(text)

	for _, item := range bookmarkList {
		if strings.Contains(item.tag, text) {
			printBookmark(item)
			someResult = true
		}
	}

	if !someResult {
		color.Red("No bookmarks with this tag(s)")
	}
}

func printBookmark(bmk bookmark) {
	idString := strconv.Itoa(bmk.id)
	idColor.Print("[" + idString + "]")
	nameColor.Print(" " + bmk.name + " ")
	linkColor.Println(bmk.url)
}

func printUsage() {
	fmt.Println("Example usage:")
	fmt.Println("  bmk search [TEXT]")
	fmt.Println("  bmk open [ID]")
	fmt.Println("  bmk add")
	fmt.Println("  bmk help")
}

func addBookmark() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter URL: ")
	url, _ := reader.ReadString('\n')
	fmt.Print("Enter tag: ")
	tag, _ := reader.ReadString('\n')
	newBmk := bookmark{42, url, name, tag}
	printBookmark(newBmk)
}

func openbrowser(id int) {
	bookmarkList := []bookmark{}
	bookmarkList = append(bookmarkList, bookmark{id: 1, url: "http://google.com", name: "Google", tag: "google"})
	bookmarkList = append(bookmarkList, bookmark{id: 2, url: "http://gmail.com", name: "Gmail", tag: "gmail"})

	var err error

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
}

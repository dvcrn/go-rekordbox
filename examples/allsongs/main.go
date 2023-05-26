package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/dvcrn/go-rekordbox/rekordbox"
)

func main() {
	ctx := context.Background()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	optionsFilePath := filepath.Join(homeDir, "/Library/Application Support/Pioneer/rekordboxAgent/storage/", "options.json")

	client, err := rekordbox.NewClient(optionsFilePath)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	contents, err := client.AllDjmdContent(ctx)
	if err != nil {
		panic(err)
	}

	sort.Slice(contents, func(i, j int) bool {
		left, _ := strconv.ParseInt(contents[i].ID.String(), 10, 64)
		right, _ := strconv.ParseInt(contents[j].ID.String(), 10, 64)
		return left < right
	})

	for _, content := range contents {
		fmt.Printf("%s (%s): %s\n", content.ID, content.UUID, content.Title)
	}
}

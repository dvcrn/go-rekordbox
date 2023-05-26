package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dvcrn/go-rekordbox/rekordbox"
	"github.com/mattn/go-nulltype"
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

	playlist, err := client.DjmdPlaylistByID(ctx, nulltype.NullStringOf("397309610"))
	if err != nil {
		panic(err)
	}

	if playlist.IsSmartlist() {
		fmt.Println("playlist is a smartlist")
		fmt.Println("title: ", playlist.Name)

		node := playlist.SmartlistNode()

		fmt.Println("conditions in this smartlist:")
		for _, condition := range node.Conditions {
			fmt.Println(condition.PropertyName)
			fmt.Println(condition.ParseValueLeft())
		}
	}
}

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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

	// query all playlists
	playlists, err := client.AllDjmdPlaylist(ctx)
	if err != nil {
		panic(err)
	}

	for _, playlist := range playlists {
		fmt.Printf("%s\n", playlist.Name.String())

		// query all songs that are in the playlist
		playlistSongs, err := client.DjmdSongPlaylistByPlaylistID(ctx, playlist.ID)
		if err != nil {
			panic(err)
		}

		// get all songs within the playlist
		for _, playlistSong := range playlistSongs {
			content, err := client.DjmdContentByID(ctx, playlistSong.ContentID)
			if err != nil {
				panic(err)
			}

			fmt.Printf("\t%s\n", content.Title.String())
		}
	}
}

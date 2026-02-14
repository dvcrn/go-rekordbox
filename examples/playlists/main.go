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
		// query all songs that are in the playlist
		playlistSongs, err := client.DjmdSongPlaylistByPlaylistID(ctx, playlist.ID)
		if err != nil {
fmt.Printf("ERROR: fetching songs for playlist %s (ID: %s): %v\n", playlist.Name.String(), playlist.ID.String(), err)
		}

		// get all songs within the playlist
		for _, playlistSong := range playlistSongs {
			content, err := client.DjmdContentByID(ctx, playlistSong.ContentID)
			if err != nil {
				fmt.Printf("DEBUG: ERROR fetching content %s: %v\n", playlistSong.ContentID.String(), err)
				continue
			}

			fmt.Printf("\t%s\n", content.Title.String())
		}
	}
}

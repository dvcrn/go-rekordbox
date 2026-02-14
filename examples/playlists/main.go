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
	fmt.Println("DEBUG: Fetching all playlists...")
	playlists, err := client.AllDjmdPlaylist(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("DEBUG: Found %d playlists\n\n", len(playlists))

	for i, playlist := range playlists {
		fmt.Printf("DEBUG: [%d] Processing playlist: %s (ID: %s)\n", i+1, playlist.Name.String(), playlist.ID.String())

		// query all songs that are in the playlist
		playlistSongs, err := client.DjmdSongPlaylistByPlaylistID(ctx, playlist.ID)
		if err != nil {
			fmt.Printf("DEBUG: ERROR fetching songs for playlist %s: %v\n", playlist.Name.String(), err)
			continue
		}

		fmt.Printf("DEBUG: Found %d songs in playlist\n", len(playlistSongs))

		// get all songs within the playlist
		for _, playlistSong := range playlistSongs {
			content, err := client.DjmdContentByID(ctx, playlistSong.ContentID)
			if err != nil {
				fmt.Printf("DEBUG: ERROR fetching content %s: %v\n", playlistSong.ContentID.String(), err)
				continue
			}

			fmt.Printf("\t%s\n", content.Title.String())
		}
		fmt.Println()
	}

}

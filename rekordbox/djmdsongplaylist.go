package rekordbox

import (
	"context"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdSongPlaylistByPlaylistID returns every track entry for a playlist.
func (c *Client) DjmdSongPlaylistByPlaylistID(ctx context.Context, playlistID nulltype.NullString) ([]*DjmdSongPlaylist, error) {
	const query = `SELECT * FROM djmdSongPlaylist WHERE PlaylistID = $1`
	logf(query, playlistID)

	rows, err := c.db.QueryContext(ctx, query, playlistID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	playlists, err := scanDjmdSongPlaylistRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return playlists, nil
}

// DjmdSongPlaylistByContentID returns every playlist entry for a track.
func (c *Client) DjmdSongPlaylistByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdSongPlaylist, error) {
	const query = `SELECT * FROM djmdSongPlaylist WHERE ContentID = $1`
	logf(query, contentID)

	rows, err := c.db.QueryContext(ctx, query, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	playlists, err := scanDjmdSongPlaylistRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return playlists, nil
}

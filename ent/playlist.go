// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/playlist"
)

// Playlist is the model entity for the Playlist schema.
type Playlist struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ThumbnailPath holds the value of the "thumbnail_path" field.
	ThumbnailPath string `json:"thumbnail_path,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlaylistQuery when eager-loading is set.
	Edges PlaylistEdges `json:"edges"`
}

// PlaylistEdges holds the relations/edges for other nodes in the graph.
type PlaylistEdges struct {
	// Vods holds the value of the vods edge.
	Vods []*Vod `json:"vods,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VodsOrErr returns the Vods value or an error if the edge
// was not loaded in eager-loading.
func (e PlaylistEdges) VodsOrErr() ([]*Vod, error) {
	if e.loadedTypes[0] {
		return e.Vods, nil
	}
	return nil, &NotLoadedError{edge: "vods"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Playlist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case playlist.FieldName, playlist.FieldDescription, playlist.FieldThumbnailPath:
			values[i] = new(sql.NullString)
		case playlist.FieldUpdatedAt, playlist.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case playlist.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Playlist", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Playlist fields.
func (pl *Playlist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playlist.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pl.ID = *value
			}
		case playlist.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case playlist.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pl.Description = value.String
			}
		case playlist.FieldThumbnailPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_path", values[i])
			} else if value.Valid {
				pl.ThumbnailPath = value.String
			}
		case playlist.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Time
			}
		case playlist.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryVods queries the "vods" edge of the Playlist entity.
func (pl *Playlist) QueryVods() *VodQuery {
	return NewPlaylistClient(pl.config).QueryVods(pl)
}

// Update returns a builder for updating this Playlist.
// Note that you need to call Playlist.Unwrap() before calling this method if this Playlist
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Playlist) Update() *PlaylistUpdateOne {
	return NewPlaylistClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Playlist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Playlist) Unwrap() *Playlist {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Playlist is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Playlist) String() string {
	var builder strings.Builder
	builder.WriteString("Playlist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pl.Description)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_path=")
	builder.WriteString(pl.ThumbnailPath)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pl.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Playlists is a parsable slice of Playlist.
type Playlists []*Playlist

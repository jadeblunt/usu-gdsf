package db

import (
	"github.com/jak103/usu-gdsf/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	game0 = models.Game{
		Name:         "game0",
		Author:       "tester",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.0",
		Tags:         []string{"tag0", "tag1"},
	}

	game1 = models.Game{
		Name:         "game1",
		Author:       "tester",
		CreationDate: time.Date(1900, 1, 2, 0, 0, 0, 0, time.UTC),
		Version:      "0.0.1",
		Tags:         []string{"tag1", "tag2"},
	}
)

func TestMongo_GameID(t *testing.T) {
	_db, _ := NewDatabaseFromEnv()

	// assign IDs on add
	id0A, _ := _db.AddGame(game0)
	id1A, _ := _db.AddGame(game1)

	// find IDs with game details
	id0F, _ := _db.GetGameID(game0)
	id1F, _ := _db.GetGameID(game1)

	// assigned IDs
	game0A, _ := _db.GetGameByID(id0A)
	game1A, _ := _db.GetGameByID(id1A)
	assert.Equal(t, game0, game0A)
	assert.Equal(t, game1, game1A)

	// found IDs
	game0F, _ := _db.GetGameByID(id0F)
	game1F, _ := _db.GetGameByID(id1F)
	assert.Equal(t, game0, game0F)
	assert.Equal(t, game1, game1F)

	// cleanup
	_db.RemoveGame(game0)
	_db.RemoveGame(game1)
}

func TestMongo_Tags(t *testing.T) {
	_db, _ := NewDatabaseFromEnv()
	_db.AddGame(game0)
	_db.AddGame(game1)

	res0, _ := _db.GetGamesByTag("tag0")
	res1, _ := _db.GetGamesByTag("tag1")
	res2, _ := _db.GetGamesByTag("tag2")
	res3, _ := _db.GetGamesByTag("bad tag")

	// result size
	assert.Equal(t, 1, len(res0))
	assert.Equal(t, 2, len(res1))
	assert.Equal(t, 1, len(res2))
	assert.Equal(t, 0, len(res3))

	// result elements
	assert.Contains(t, res0, game0)
	assert.Contains(t, res1, game0)
	assert.Contains(t, res1, game1)
	assert.Contains(t, res2, game1)

	// cleanup
	_db.RemoveGame(game0)
	_db.RemoveGame(game1)
}

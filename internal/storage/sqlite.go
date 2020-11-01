package storage

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/chrusty/tunecast/api"
	"github.com/chrusty/tunecast/internal/configuration"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"
	_ "github.com/mattn/go-sqlite3"
)

// SQLite is a storage implementation backed by SQLite:
type SQLite struct {
	config *configuration.Config
	dbConn *sql.DB
	logger *logrus.Logger
}

// NewSQLite returns a configured SQLite storage implementation:
func NewSQLite(logger *logrus.Logger, config *configuration.Config) (*SQLite, error) {

	logger.Info("Preparing an SQLite storage interface ...")

	// Create the SQLite DB file (if it doesn't exist):
	if _, err := os.Stat(config.Database.SQLiteDBPath()); os.IsNotExist(err) {
		file, err := os.Create(config.Database.SQLiteDBPath())
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Establish an SQLite DB:
	dbConn, err := sql.Open("sqlite3", config.Database.SQLiteDBPath())
	if err != nil {
		return nil, err
	}

	// Put the deps into an SQLite:
	newSQLite := &SQLite{
		config: config,
		dbConn: dbConn,
		logger: logger,
	}

	// Make sure the schema exists:
	if err := newSQLite.MigrateSchema(); err != nil {
		return nil, err
	}

	return newSQLite, nil
}

// MigrateSchema ensures that the schema exists:
func (s *SQLite) MigrateSchema() error {
	libraryItemsTableSQL := `CREATE TABLE IF NOT EXISTS libraryItems (
		"path" TEXT NOT NULL PRIMARY KEY,		
		"uuid" TEXT UNIQUE,
		"folder" BOOLEAN,
		"cover" TEXT,
		"added" DATETIME DEFAULT CURRENT_TIMESTAMP	
	  );`

	// Prepare a statement:
	statement, err := s.dbConn.Prepare(libraryItemsTableSQL)
	if err != nil {
		return err
	}

	// Run the statement:
	_, err = statement.Exec()
	return err
}

// Stop shuts down gracefully:
func (s *SQLite) Stop() error {
	return s.dbConn.Close()
}

// AddLibraryItem adds a library item to the DB:
func (s *SQLite) AddLibraryItem(libraryItem *api.LibraryItem) error {

	// Mint a new UUID for this item:
	newUUID := uuid.New()

	// Prepare a statement:
	addLibraryItemSQL := fmt.Sprintf(`INSERT INTO libraryItems(path, uuid, folder, cover) VALUES('%s', '%s', '%v', '%s') ON CONFLICT(path) DO UPDATE SET cover='%s';`, url.QueryEscape(*libraryItem.Path), newUUID.String(), *libraryItem.IsFolder, *libraryItem.Cover, *libraryItem.Cover)
	statement, err := s.dbConn.Prepare(addLibraryItemSQL)
	if err != nil {
		return err
	}

	// Run the statement:
	_, err = statement.Exec()
	return err
}

// List returns a list of library items beneath a given path:
func (s *SQLite) List(parentPath string, sortBy string) ([]*api.LibraryItem, error) {
	libraryItems := []*api.LibraryItem{}

	// Prepare a query:
	listLibraryItemSQL := fmt.Sprintf(`SELECT path, uuid, folder, cover, added FROM libraryItems LIMIT 100;`)
	rows, err := s.dbConn.Query(listLibraryItemSQL)
	if err != nil {
		return nil, err
	}

	// Process the rows:
	for rows.Next() {
		libraryItem := &api.LibraryItem{}
		var addedTime time.Time
		rows.Scan(&libraryItem.Path, &libraryItem.Uuid, &libraryItem.IsFolder, &libraryItem.Cover, &addedTime)
		libraryItem.Added = &openapiTypes.Date{
			addedTime,
		}
		libraryItems = append(libraryItems, libraryItem)
	}

	return libraryItems, nil
}

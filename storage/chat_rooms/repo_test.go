package chatrooms

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/wesmota/go-jobsity-chat-server/logger"
)

type ChatsTestSuite struct {
	suite.Suite
	repo *Repo
}

func TestChatsSuite(t *testing.T) {
	suite.Run(t, new(ChatsTestSuite))
}

func (suite *ChatsTestSuite) SetupSuite() {
	suite.repo = getTestRepo()
	suite.Require().NotNil(suite.repo)
}

func getTestRepo() *Repo {
	log := logger.NewZerologLogger(zerolog.New(os.Stdout))

	db := newDB(&testing.T{})
	testRepoInstance, _ := NewChatRoomsRepo(db, log)
	return testRepoInstance
}

func newDB(t *testing.T) *sql.DB {

	var db *sql.DB
	var err error
	const (
		DBName = "jobsity_test"
		host   = "localhost"
		port   = 5432
		user   = "root"
		pass   = "password"
	)

	// Open a connection to the root database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, pass)
	db, err = sql.Open("postgres", dsn)
	fmt.Printf("Error: %v\n", err)
	require.NoError(t, err)

	// Drop the DB if it exists
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", DBName))
	require.NoError(t, err)

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", DBName))
	require.NoError(t, err)

	// Close connection to default container database
	require.NoError(t, db.Close())

	// Open a connection to the new database
	dsn = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, DBName)
	db, err = sql.Open("postgres", dsn)
	require.NoError(t, err)

	// Execute the schema file SQL commands on the new database
	schemaFile, err := os.ReadFile(filepath.Join(GetTestRoot(t), "db", "schema.sql"))
	require.NoError(t, err)

	_, err = db.Exec(string(schemaFile))
	//fmt.Printf("Schema: %s\n", string(schemaFile))
	require.NoError(t, err)

	// Execute the seed file SQL commands on the new database
	seedFile, err := os.ReadFile(filepath.Join(GetTestRoot(t), "db/setup/local", "seed.sql"))
	require.NoError(t, err)

	_, err = db.Exec(string(seedFile))
	fmt.Printf("Error: %v\n", err)
	require.NoError(t, err)

	require.NoError(t, db.Close())

	db, err = sql.Open("postgres", dsn)
	require.NoError(t, err)

	return db
}

// GetTestRoot finds and returns and the project root for tests.
func GetTestRoot(t *testing.T) string {
	exists := func(t *testing.T, n string) bool {
		_, err := os.Stat(n)
		if os.IsNotExist(err) {
			return false
		}
		require.NoError(t, err)
		return true
	}
	dir, _ := os.Getwd()
	for {
		root := filepath.Join(dir, ".git")
		if exists(t, root) {
			break
		}
		dir = filepath.Dir(dir)
	}
	require.NotEmpty(t, dir)
	return dir
}

package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // для того, чотбы отработала функция init()
)

// Instance of storage
type Storage struct {
	config *Config
	// DataBase FileDescriptor
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

// Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

// Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}

func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return nil
}

func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return nil
}

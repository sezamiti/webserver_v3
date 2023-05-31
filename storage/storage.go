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
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURI)
	if err != nil {
		log.Println("Error opening database:", err)
		return err
	}
	if err := db.Ping(); err != nil {
		log.Println("Error pinging database:", err)
		return err
	}
	s.db = db
	log.Println("Database connection created successfully!")
	return nil
}

// Close connection
func (s *Storage) Close() {
	s.db.Close()
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

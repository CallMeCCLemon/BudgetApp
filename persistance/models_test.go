package persistance

import (
	"github.com/google/uuid"
	"log"
	"os"
	"testing"
)

func TestStorageDao_WriteAccount(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}

	t.Run("Write something", func(t *testing.T) {
		_, err = dao.WriteAccount(Account{
			Name: "test-account-001",
			ID:   uuid.New(),
		})
		if err != nil {
			t.Fatal("Failed to write account to the MySQL DB!", err)
		}
	})
}

package persistance

import (
	"github.com/google/uuid"
	"gotest.tools/v3/assert"
	"log"
	"os"
	"testing"
)

func TestStorageDao_WriteAccount(t *testing.T) {
	dao, err := NewStorageDao("root", os.Getenv("password"), "127.0.0.1", "budgetApp")
	if err != nil {
		log.Fatal("Failed to connect to the MySQL DB!", err)
	}
	id := uuid.New()
	name := "test-account-001"

	t.Run("Write then Read something", func(t *testing.T) {
		account := Account{
			Name: name,
			ID:   id,
		}
		_, err := dao.WriteAccount(account)
		if err != nil {
			t.Fatal("Failed to write account to the MySQL DB!", err)
		}
		//assert.Equal(t, new_id, id)
		savedAccount, err := dao.ReadAccount(id)
		if err != nil {
			t.Fatal("Failed to Read account from the MySQL DB!", err)
		}
		assert.Equal(t, savedAccount.ID, id)
		assert.Equal(t, savedAccount.Name, name)
	})
}

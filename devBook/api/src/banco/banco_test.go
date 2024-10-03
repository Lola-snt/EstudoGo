package banco_test

import (
	"api/src/banco"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConectar(t *testing.T) {
	t.Run("deve retonar sucesso quando conectar no banco", func(t *testing.T) {
		db, err := banco.Conectar()
		assert.Nil(t, err)
		assert.NotNil(t, db)

	})
}

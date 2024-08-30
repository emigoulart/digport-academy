package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsuarioValidar(t *testing.T) {
	t.Run("Deve retornar erro se nome for vazio", func(t *testing.T) {
		usuario := Usuario{
			Nome:  "",
			Email: "",
			Senha: "123456",
		}
		msgEsperada := "nome não pode ser vazio"

		err := usuario.Validar()

		assert.NotNil(t, err)
		assert.Equal(t, msgEsperada, err.Error())
	})

	t.Run("Deve retornar erro se email for vazio", func(t *testing.T) {
		usuario := Usuario{
			Nome:  "Someone",
			Email: "",
			Senha: "123456",
		}

		msgEsperada := "email não pode ser vazio"

		err := usuario.Validar()

		assert.NotNil(t, err)
		assert.Equal(t, msgEsperada, err.Error())
	})

	t.Run("Deve retornar erro se senha for vazia", func(t *testing.T) {
		usuario := Usuario{
			Nome:  "Someone",
			Email: "someone",
			Senha: "",
		}

		msgEsperada := "senha não pode ser vazia"

		err := usuario.Validar()

		assert.Error(t, err)
		assert.Equal(t, msgEsperada, err.Error())
	})

	t.Run("Deve retornar nil se todos os campos estiverem preenchidos", func(t *testing.T) {
		usuario := Usuario{
			Nome:  "Someone",
			Email: "someone",
			Senha: "123456",
		}

		err := usuario.Validar()

		assert.NoError(t, err)
	})
}

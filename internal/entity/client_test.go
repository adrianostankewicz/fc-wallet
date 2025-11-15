package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Jhon Okonor", "j@o.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Jhon Okonor", client.Name)
	assert.Equal(t, "j@o.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Jhon Okonor", "j@o.com")
	err := client.Update("Jhon Okonor Updated", "j_updated@o.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jhon Okonor Updated", client.Name)
	assert.Equal(t, "j_updated@o.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Jhon Okonor", "j@o.com")
	err := client.Update("", "j_updated@o.com")
	assert.Error(t, err, "name is required")
}

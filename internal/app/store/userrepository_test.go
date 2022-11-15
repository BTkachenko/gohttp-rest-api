package store_test

import (
	"testing"

	"github.com/BTkachenko/gotest/internal/app/model"
	"github.com/BTkachenko/gotest/internal/app/store"
	"github.com/stretchr/testify/assert"
)


func TestUserRepoistory_Create(t *testing.T) {
	s , teardown := store.TestStore(t,databaseURL)

	defer teardown("users")

	u,err := s.User().Create(model.TestUser(t))
	assert.NoError(t,err)
	assert.NotNil(t,u)
}


func TestUserRepository_FindbyEmail(t *testing.T){
	s , teardown := store.TestStore(t,databaseURL)

	defer teardown("users")

	email := "notexist@example.org"
	_,err := s.User().FindByEmail(email)
	assert.Error(t,err)
	

	u:= model.TestUser(t)
	u.Email = email


	 s.User().Create(u)

	u,err = s.User().FindByEmail(email)
	assert.NoError(t,err)
	assert.NotNil(t,u)
}
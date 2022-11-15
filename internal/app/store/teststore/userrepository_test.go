package teststore_test

import (
	"testing"

	"github.com/BTkachenko/gotest/internal/app/model"
	"github.com/BTkachenko/gotest/internal/app/store"
	"github.com/BTkachenko/gotest/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)


func TestUserRepoistory_Create(t *testing.T) {
	

	s:= teststore.New()


	u := model.TestUser(t)
	assert.NoError(t,s.User().Create(u))
	assert.NotNil(t,u)
}


func TestUserRepository_FindbyEmail(t *testing.T){
	
	s:= teststore.New()

	email := "notexist@example.org"
	_,err := s.User().FindByEmail(email)
	assert.EqualError(t,err,store.ErrRecordNotFound.Error())
	

	u:= model.TestUser(t)
	u.Email = email


	 s.User().Create(u)

	u,err = s.User().FindByEmail(email)
	assert.NoError(t,err)
	assert.NotNil(t,u)
}
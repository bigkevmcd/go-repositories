package project

import (
	gc "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

type RepoSuite struct{}

var _ = gc.Suite(&RepoSuite{})

type TestRepository struct{}
type TestModel struct{}

func (s *RepoSuite) TestRegisterType(c *gc.C) {
	repo := New()
	testRepo := TestRepository{}
	repo.Register(TestModel{}, testRepo)
}

func (s *RepoSuite) TestGetRepositoryForModel(c *gc.C) {
	repo := New()
	testRepo := TestRepository{}
	repo.Register(TestModel{}, testRepo)

	fetchedRepo, err := repo.Get(TestModel{})
	c.Assert(err, gc.IsNil)
	c.Assert(fetchedRepo, gc.Equals, testRepo)
}

func (s *RepoSuite) TestRegisterDefault(c *gc.C) {
	testRepo := TestRepository{}
	Register(TestModel{}, testRepo)
}

func (s *RepoSuite) TestGetRepositoryForModelDefault(c *gc.C) {
	testRepo := TestRepository{}
	Register(TestModel{}, testRepo)

	fetchedRepo, err := Get(TestModel{})
	c.Assert(err, gc.IsNil)
	c.Assert(fetchedRepo, gc.Equals, testRepo)
}

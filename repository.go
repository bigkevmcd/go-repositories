package project

import (
	"errors"
	"reflect"
)

// A Repository maintains a mapping between types and values.
type Repository struct {
	repoTypes map[reflect.Type]interface{}
}

// Register associates the repo object with the type of modelType.
func (r *Repository) Register(modelType, repo interface{}) {
	r.repoTypes[reflect.TypeOf(modelType)] = repo
}

// Get returns the repository associated with the type of the provided model returns an error if we can't find one.
func (r *Repository) Get(modelType interface{}) (interface{}, error) {
	repo, ok := r.repoTypes[reflect.TypeOf(modelType)]
	if !ok {
		return nil, errors.New("couldn't find repo for model")
	}
	return repo, nil
}

// New Repository.
func New() *Repository {
	return &Repository{make(map[reflect.Type]interface{})}
}

var defaultRepo = New()

// Register the provided repository with the type uses the default Repository.
func Register(modelType, repo interface{}) {
	defaultRepo.Register(modelType, repo)
}

// Get returns the repository associated with the type of the provided model from the default Repository.
func Get(modelType interface{}) (interface{}, error) {
	return defaultRepo.Get(modelType)
}

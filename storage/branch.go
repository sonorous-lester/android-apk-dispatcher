package storage

import (
	"log"
	"sync"
)

type storage struct {
	cabinet sync.Map
}
var branchStorage *storage
var doOnce sync.Once
func GetInstance() *storage {
	var fn = func() { branchStorage = &storage{}}
	doOnce.Do(fn)
	return branchStorage
}

func (vs *storage) Add(tag string, branch string) {
	vs.cabinet.Store(tag, branch)
}

func (vs *storage) Remove(tag string){
	vs.cabinet.Delete(tag)
}

func (vs *storage) GetBranch(tag string) string{
	value, ok := vs.cabinet.Load(tag)
	if !ok {
		log.Printf("can't find branch name by user id: %s", tag)
	}
	return value.(string)
}

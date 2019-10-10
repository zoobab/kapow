package data

import (
	"errors"
	"sync"

	"github.com/BBVA/kapow/internal/server/model"
)

type safeHandlerMap struct {
	hs map[string]*model.Handler
	m  *sync.RWMutex
}

var Handlers = New()

func New() safeHandlerMap {
	return safeHandlerMap{
		hs: make(map[string]*model.Handler),
		m:  &sync.RWMutex{},
	}
}

func (shm *safeHandlerMap) Add(h *model.Handler) {
	shm.m.Lock()
	shm.hs[h.ID] = h
	shm.m.Unlock()
}

func (shm *safeHandlerMap) Remove(id string) {
	shm.m.Lock()
	delete(shm.hs, id)
	shm.m.Unlock()
}

func (shm *safeHandlerMap) Get(id string) (*model.Handler, bool) {
	shm.m.RLock()
	h, ok := shm.hs[id]
	shm.m.RUnlock()
	return h, ok
}

func (shm *safeHandlerMap) ListIDs() (ids []string) {
	shm.m.RLock()
	defer shm.m.RUnlock()
	for id := range shm.hs {
		ids = append(ids, id)
	}
	return
}

//TODO: Test this mess
type HandlerFunction func(*model.Handler) error

func (shm *safeHandlerMap) ReadSafe(id string, f HandlerFunction) error {
	shm.m.RLock()
	defer shm.m.RUnlock()

	return mapOp(shm, id, f)
}

//TODO: Test this mess
func (shm *safeHandlerMap) WriteSafe(id string, f HandlerFunction) error {
	shm.m.Lock()
	defer shm.m.Unlock()

	return mapOp(shm, id, f)
}

func mapOp(shm *safeHandlerMap, id string, f HandlerFunction) error {
	h, ok := shm.hs[id]
	if !ok {
		return errors.New("no handler found")
	}

	return f(h)
}

//TODO: Test this mess
func (shm *safeHandlerMap) Has(id string) bool {
	shm.m.RLock()
	defer shm.m.RUnlock()
	_, ok := shm.hs[id]
	return ok
}

package mysql

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/util/log"
)

var handlers = make(map[string]*Handler)
var hLock sync.RWMutex

func Init(confMap map[string]*Conf) {
	hLock.Lock()
	defer hLock.Unlock()
	for name, conf := range confMap {
		if _, ok := handlers[name]; ok {
			log.Infof("db handler %s has been initialized\n", name)
			continue
		}
		log.Infof("init db handler %s...\n", name)
		// lazy init, do not connect now
		h := NewHandler(conf)
		handlers[name] = h
	}
}

func GetHandler(name string) *Handler {
	hLock.RLock()
	defer hLock.RUnlock()
	if h, ok := handlers[name]; ok {
		return h
	}
	return nil
}

func CloseHandlers() {
	hLock.Lock()
	defer hLock.Unlock()
	for name, h := range handlers {
		log.Infof("close db handler %s...\n", name)
		delete(handlers, name)
		err := h.Close()
		if err != nil {
			log.Errorf("db handler %s close failed\n", name)
		}
	}
}

type Handler struct {
	db        *gorm.DB
	conf      *Conf
	connected bool
	lock      sync.RWMutex
}

func NewHandler(conf *Conf) *Handler {
	return &Handler{
		conf:      conf,
		connected: false,
	}
}

// DCL
func (h *Handler) Connect() error {
	if h.connected {
		return nil
	}
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.connected {
		return nil
	}
	db, err := gorm.Open("mysql", h.conf.BuildArgs())
	if err != nil {
		return err
	}
	h.db = db
	h.connected = true
	h.db.DB().SetMaxOpenConns(h.conf.MaxOpenConns)
	h.db.DB().SetMaxIdleConns(h.conf.MaxIdleConns)
	h.db.SingularTable(true)
	return nil
}

// thread safe to get gorm.DB
func (h *Handler) GetDb() (*gorm.DB, error) {
	err := h.Connect()
	return h.db, err
}

func (h *Handler) Close() error {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.connected {
		err := h.db.Close()
		if err != nil {
			return err
		}
		h.connected = false
	}
	return nil
}

func (h *Handler) IsConnected() bool {
	return h.connected
}

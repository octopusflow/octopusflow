package mysql

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Handler struct {
	db        *gorm.DB
	conf      *Conf
	connected bool
	mu        sync.RWMutex
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
	h.mu.Lock()
	defer h.mu.Unlock()
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
	h.mu.Lock()
	defer h.mu.Unlock()
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

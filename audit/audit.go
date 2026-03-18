package audit

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type AuditEvent struct {
	ID      uint      `json:"id,omitempty"`
	When    time.Time `json:"when"`
	UserID  uint      `json:"user_id"`
	Action  Action    `json:"action"`
	Path    string    `json:"path"`
	OldPath string    `json:"old_path,omitempty"`
	Details string    `json:"details,omitempty"`
}

var (
	mu          sync.Mutex
	out         *os.File
	currentTime time.Time
)

func Init() error {
	mu.Lock()
	defer mu.Unlock()
	return initLocked()
}

func initLocked() error {
	today := time.Now()
	path := fmt.Sprintf("audit-%s.log", today.Format("20060102"))
	fullPath := filepath.Join("logs", path)

	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}

	f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}

	if out != nil {
		out.Close()
	}

	out = f
	currentTime = today
	return nil
}

func Close() error {
	mu.Lock()
	defer mu.Unlock()

	if out == nil {
		return nil
	}

	err := out.Close()
	out = nil
	return err
}

func Log(userID uint, action Action, path, oldPath string) error {
	mu.Lock()
	defer mu.Unlock()
	
	if !ValidateTime() {
	if err := initLocked(); err != nil {
		return err
	}
}

	if out == nil {
		return nil
	}

	e := AuditEvent{
		When:    time.Now(),
		UserID:  userID,
		Action:  action,
		Path:    path,
		OldPath: oldPath,
	}

	enc := json.NewEncoder(out)
	return enc.Encode(&e)
}

func ValidateTime() bool {
	return currentTime.Format("20060102") == time.Now().Format("20060102")
}
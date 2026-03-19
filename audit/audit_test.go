package audit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestInitLocked(t *testing.T) {
	logDir = ".logs"

	err := initLocked()
	if err != nil {
		t.Fatalf("initLocked failed: %v", err)
	}

	today := time.Now().Format("20060102")
	path := filepath.Join(logDir, fmt.Sprintf("audit-%s.log", today))

	// Check file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("log file not created")
	}

	// Write to log
	msg := "test log entry\n"
	_, err = out.WriteString(msg)
	if err != nil {
		t.Fatalf("write failed: %v", err)
	}

	// Verify contents
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}

	if !strings.Contains(string(data), msg) {
		t.Fatalf("log entry not found")
	}
}

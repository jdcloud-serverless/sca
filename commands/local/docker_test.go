package local

import (
	"testing"
	"time"
)

func TestCreateStartContainer(t *testing.T) {
	hostVolPath := "/root/code"

	dClient, err := NewDockerClient()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if err := dClient.CreateDocker(hostVolPath); err != nil {
		t.Errorf("error: %v", err)
	}
	if err := dClient.StartDocker(); err != nil {
		t.Errorf("error: %v", err)
	}

	time.Sleep(time.Minute * time.Duration(10))

	if err := dClient.StopDocker(); err != nil {
		t.Errorf("error: %v", err)
	}
	if err := dClient.RemoveDocker(); err != nil {
		t.Errorf("error: %v", err)
	}
}

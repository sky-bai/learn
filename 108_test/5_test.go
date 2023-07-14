package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestXidLoadBalance(t *testing.T) {
	sessions := &sync.Map{}
	session1 := &MockSession{remoteAddr: "192.168.0.1:1234", closed: false}
	session2 := &MockSession{remoteAddr: "192.168.0.2:5678", closed: false}
	session3 := &MockSession{remoteAddr: "192.168.0.3:9012", closed: false}
	sessions.Store(session1, nil)
	sessions.Store(session2, nil)
	sessions.Store(session3, nil)
	// Positive test case: valid xid with matching session
	session := XidLoadBalance(sessions, "192.168.0.2:5678:123")
	assert.Equal(t, session2, session)
	// Positive test case: valid xid without matching session
	session = XidLoadBalance(sessions, "192.168.0.4:1111:456")
	assert.NotNil(t, session)
	// Negative test case: invalid xid format
	session = XidLoadBalance(sessions, "invalid_xid_format")
	assert.NotNil(t, session)
}

type MockSession struct {
	remoteAddr string
	closed     bool
}

func (s *MockSession) RemoteAddr() string {
	return s.remoteAddr
}
func (s *MockSession) IsClosed() bool {
	return s.closed
}

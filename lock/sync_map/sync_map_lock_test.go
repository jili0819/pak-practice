package sync_map

import (
	"testing"
)

func TestKeyedMutex_Lock(t *testing.T) {
	unlock := km.Lock("test")
	defer unlock()
	// 按key加锁，颗粒度问题
}

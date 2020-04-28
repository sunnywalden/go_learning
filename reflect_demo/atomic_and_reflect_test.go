package reflect_demo

import (
	"sync/atomic"
	"testing"
	"unsafe"
)


func TestBuffData(t *testing.T) {
	testData := []string{"Henry Zhang", "Mesuit ozil"}
	var tmpData unsafe.Pointer
	atomic.StorePointer(&tmpData, unsafe.Pointer(&testData))
	buffData := atomic.LoadPointer(&tmpData)
	t.Log( *(*[]string)(buffData))
}
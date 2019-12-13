package sync

import "sync"

// Read func returns should execute write func
type ReadFunc func() bool

func Do(rwMutex sync.RWMutex, readFunc ReadFunc, writeFunc func()) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	if readFunc() {
		//       code block start								       code block end
		//             ⬇													  ⬇
		// ReadLock -> {  ReadUnlock -> Lock -> write() -> UnLock -> ReadLock } -> ReadUnlock

		rwMutex.RUnlock()
		defer rwMutex.RLock()
		rwMutex.Lock()
		defer rwMutex.Unlock()
		writeFunc()
	}
}

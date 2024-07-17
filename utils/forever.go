package utils

import "sync"

func GoForever() {
	var done sync.WaitGroup
	done.Add(1)
	go forever(&done)
	done.Wait()
}

func forever(done *sync.WaitGroup) {
	defer done.Done()
	select {}
}

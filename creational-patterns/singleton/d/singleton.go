package singleton

import "sync"

// Singleton merupakan sebuah struct
// dimana hanya satu instance saja yang boleh ada
type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

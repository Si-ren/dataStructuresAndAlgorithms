package singleton

import "sync"

type Singleton struct {
	A int
}

var (
	once          sync.Once
	singletonLazy *Singleton
)

//懒汉模式，当用的时候才初始化
func GetSingleton() *Singleton {
	once.Do(func() {
		singletonLazy = &Singleton{A: 1}
	})
	return singletonLazy
}

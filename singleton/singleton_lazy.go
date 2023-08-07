package singleton

var Singleton1 *Singleton

// 饿汉模式，程序启动时就初始化
func init() {
	Singleton1 = &Singleton{A: 1}
}

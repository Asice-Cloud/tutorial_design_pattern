package singleton

import (
	"sync"
)

// 创建型模式：单例模式

/*
单例模式具有以下特点和优点：

确保类只有一个实例对象
提供全局访问点，方便外部代码获取该实例
避免重复创建实例，节省系统资源
*/

/*
单例模式适用于以下应用场景：

日志记录器：在整个系统中只能有一个日志记录器，以保证日志不会被重复记录。
数据库连接池：在高并发环境下，使用单例模式可以避免频繁创建和销毁数据库连接。
*/

// 懒汉式单例模式实现
type Singleton struct{}

var instance0 *Singleton

// GetInstance 返回懒汉式单例模式的实例。
// 如果实例尚未创建，则创建一个新的实例。
func GetInstance() *Singleton {
	if instance0 == nil {
		instance0 = new(Singleton)
	}
	return instance0
}

// 饿汉式单例模式实现
var instance2 = new(Singleton)

// GetInstance2 返回饿汉式单例模式的实例。
// 实例在程序启动时即已创建。
func GetInstance2() *Singleton {
	return instance2
}

// 通过使用 sync.Once 来保证只有一个 goroutine 执行初始化代码，从而解决了线程安全问题。
var instance3 *Singleton
var once sync.Once

// GetInstance3 返回使用 sync.Once 实现的线程安全单例模式的实例。
// 确保只有一个 goroutine 执行初始化代码。
func GetInstance3() *Singleton {
	once.Do(func() {
		instance3 = new(Singleton)
	})
	return instance3
}

// 另一种线程安全的延迟初始化单例模式的实现方式是使用 sync.Mutex 来加锁，保证只有一个 goroutine 执行初始化操作。
var mu sync.Mutex
var instance4 *Singleton

// GetInstance4 返回使用 sync.Mutex 实现的线程安全单例模式的实例。
// 通过加锁保证只有一个 goroutine 执行初始化操作。
func GetInstance4() *Singleton {
	if instance4 == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance4 == nil {
			instance4 = new(Singleton)
		}
	}
	return instance4
}

package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Errorf("Expected same instance but got different instances")
	}
}

func TestGetInstance2(t *testing.T) {
	instance1 := GetInstance2()
	instance2 := GetInstance2()

	if instance1 != instance2 {
		t.Errorf("Expected same instance but got different instances")
	}
}

func TestGetInstance3(t *testing.T) {
	var wg sync.WaitGroup
	var instance1, instance2 *Singleton

	wg.Add(2)
	go func() {
		defer wg.Done()
		instance1 = GetInstance3()
	}()
	go func() {
		defer wg.Done()
		instance2 = GetInstance3()
	}()
	wg.Wait()

	if instance1 != instance2 {
		t.Errorf("Expected same instance but got different instances")
	}
}

func TestGetInstance4(t *testing.T) {
	var wg sync.WaitGroup
	var instance1, instance2 *Singleton

	wg.Add(2)
	go func() {
		defer wg.Done()
		instance1 = GetInstance4()
	}()
	go func() {
		defer wg.Done()
		instance2 = GetInstance4()
	}()
	wg.Wait()

	if instance1 != instance2 {
		t.Errorf("Expected same instance but got different instances")
	}
}

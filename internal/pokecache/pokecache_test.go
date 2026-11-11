package pokecache

import (
	"testing"
	"time"
)

const (
	reapInterval = time.Second
)

func TestCacheEmpty(t *testing.T) {

	cache := NewCache(reapInterval)

	// test getting a value from empty cache
	value, ok := cache.Get("nonexistent")
	if ok {
		t.Errorf("Expected false for non-existent key, got true")
	}
	if value != nil {
		t.Errorf("Expected nil for non-existent key, got %v", value)
	}
}

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(reapInterval)
	key := "key"
	value := []byte("value")
	cache.Add(key, value)

	got, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected true for existing key, got false")
	}
	if string(got) != string(value) {
		t.Errorf("Expected %v for key %s, got %v", value, key, got)
	}
}

func TestCacheOverwrittenValue(t *testing.T) {
	cache := NewCache(reapInterval)
	key := "key"
	value1 := []byte("value1")
	value2 := []byte("value2")
	cache.Add(key, value1)

	var got []byte
	var ok bool
	got, ok = cache.Get(key)
	if !ok {
		t.Errorf("Expected true for existing key, got false")
	}
	if string(got) != string(value1) {
		t.Errorf("Expected %v for key %s, got %v", value1, key, got)
	}

	cache.Add(key, value2)

	got, ok = cache.Get(key)
	if !ok {
		t.Errorf("Expected true for existing key, got false")
	}
	if string(got) != string(value2) {
		t.Errorf("Expected %v for key %s, got %v", value2, key, got)
	}
}

func TestCacheReaping(t *testing.T) {
	cache := NewCache(reapInterval)
	key := "key"
	value := []byte("value")
	cache.Add(key, value)

	var got []byte
	var ok bool

	got, ok = cache.Get(key)
	if !ok {
		t.Errorf("Expected true for existing key, got false")
	}
	if string(got) != string(value) {
		t.Errorf("Expected %v for key %s, got %v", value, key, got)
	}

	time.Sleep(reapInterval + 1*time.Second)

	got, ok = cache.Get(key)
	if ok {
		t.Errorf("Expected false for reaped key, got true")
	}
	if got != nil {
		t.Errorf("Expected nil for reaped key, got %v", got)
	}
}

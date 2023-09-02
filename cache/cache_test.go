package cache

import (
	"sync"
	"time"
	"testing"
)

func TestNewCache(t *testing.T) {
	cache := New()
	if cache == nil {
		t.Error("New() failed")
	}	
}

func TestGetCache(t *testing.T) {
	cache := New()
	value, ok := cache.Get("test", func() interface{} {
		return "test"
	}, 10).(string)

	if !ok {
		t.Error("cache.Get() failed")
	}

	if value != "test" {
		t.Error("cache.Get() failed")
	}
}

func TestGetCacheWithoutExpire(t *testing.T) {
	cache := New()
	value, ok := cache.Get("test", func() interface{} {
		return "test"
	}, -1).(string)

	if !ok {
		t.Error("cache.Get() failed")
	}

	if value != "test" {
		t.Error("cache.Get() failed")
	}

	value, ok = cache.Get("test", func() interface{} {
		return "test2"
	}, -1).(string)

	if !ok {
		t.Error("cache.Get() failed")
	}

	if value != "test" {
		t.Error("cache.Get() failed")
	}
}

func TestGetCacheWithExpire(t *testing.T) {
	cache := New()
	value, ok := cache.Get("test", func() interface{} {
		return "test"
	}, 2).(string)

	if !ok {
		t.Error("cache.Get() failed")
	}

	if value != "test" {
		t.Error("cache.Get() failed")
	}

	time.Sleep(2 * time.Second)

	value, ok = cache.Get("test", func() interface{} {
		return "test2"
	}, 2).(string)

	if !ok {
		t.Error("cache.Get() failed")
	}

	if value != "test2" {
		t.Error("cache.Get() failed")
	}
}

func TestGetCacheParallelRead(t *testing.T) {
	cache := New()
	value, ok := cache.Get("test", func() interface{} {
		return "test"
	}, -1).(string)

	if !ok {
		t.Error("cache.Get() failed")
	}

	if value != "test" {
		t.Error("cache.Get() failed")
	}

	time.Sleep(2 * time.Second)

	wg := sync.WaitGroup{}
	
	wg.Add(3)
	go func() {
		value, ok = cache.Get("test", func() interface{} {
			return "test2"
		}, -1).(string)

		if !ok {
			t.Error("cache.Get() failed")
		}

		if value != "test" {
			t.Error("cache.Get() failed")
		}
		wg.Done()
	}()

	go func() {
		value, ok = cache.Get("test", func() interface{} {
			return "test3"
		}, -1).(string)

		if !ok {
			t.Error("cache.Get() failed")
		}

		if value != "test" {
			t.Error("cache.Get() failed")
		}
		wg.Done()
	}()

	go func() {
		value, ok = cache.Get("test", func() interface{} {
			return "test4"
		}, -1).(string)

		if !ok {
			t.Error("cache.Get() failed")
		}

		if value != "test" {
			t.Error("cache.Get() failed")
		}
		wg.Done()
	}()

	wg.Wait()
}

func TestGetCacheParallelWrite(t *testing.T) {
	cache := New()
	value, ok := cache.Get("test", func() interface{} {
		return "test"
	}, 2).(string)

	if !ok {
		t.Error("cache.Get() failed")
	}

	if value != "test" {
		t.Error("cache.Get() failed")
	}

	time.Sleep(2 * time.Second)
	

	channels := make([]chan string, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan string, 1)
	}

	go func() {
		value, ok = cache.Get("test", func() interface{} {
			return "test2"
		}, -1).(string)

		if !ok {
			t.Error("cache.Get() failed")
		}

		if value == "test" {
			t.Error("cache.Get() failed")
		}
		channels[0] <- value
	}()

	go func() {
		value, ok = cache.Get("test", func() interface{} {
			return "test3"
		}, -1).(string)

		if !ok {
			t.Error("cache.Get() failed")
		}

		if value == "test" {
			t.Error("cache.Get() failed")
		}
		channels[1] <- value
	}()

	go func() {
		value, ok = cache.Get("test", func() interface{} {
			return "test4"
		}, -1).(string)

		if !ok {
			t.Error("cache.Get() failed")
		}

		if value == "test" {
			t.Error("cache.Get() failed")
		}
		channels[2] <- value
	}()

	go func() {
		value, ok = cache.Get("test", func() interface{} {
			return "test5"
		}, -1).(string)

		if !ok {
			t.Error("cache.Get() failed")
		}

		if value == "test" {
			t.Error("cache.Get() failed")
		}
		channels[3] <- value
	}()

	v := <- channels[0]
	for i := 1; i < 4; i++ {
		if v != <- channels[i] {
			t.Error("cache.Get() failed")
		}
	}
}


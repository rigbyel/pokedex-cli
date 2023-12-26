package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	c := NewCache(10 * time.Millisecond)

	if c.cache == nil {
		t.Error("nil cache initialized")
	}
}

func TestAddGetCache(t *testing.T) {
	myCache := NewCache(10 * time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	} {
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
	}

	for _, c := range cases {
		myCache.Add(c.inputKey, c.inputVal)
		actualVal, ok := myCache.Get(c.inputKey)

		if !ok {
			t.Errorf("%s value not found", c.inputKey)
		}
		if string(actualVal) != string(c.inputVal) {
			t.Errorf("wrong value; expected %s, got %s", c.inputVal, actualVal)
		}	
	}
}

func TestReap(t *testing.T) {
	cacheLifetime := 10*time.Millisecond
	myCache := NewCache(cacheLifetime)
	sampleKey := "key1"
	myCache.Add(sampleKey, []byte("value1"))

	time.Sleep(cacheLifetime + time.Millisecond)
	if _, ok := myCache.Get(sampleKey); ok {
		t.Errorf("%s should have been removed", sampleKey)
	}
}

func TestReapFail(t *testing.T) {
	cacheLifetime := 10*time.Millisecond
	myCache := NewCache(cacheLifetime)
	sampleKey := "key1"
	myCache.Add(sampleKey, []byte("value1"))

	time.Sleep(cacheLifetime / 2)
	if _, ok := myCache.Get(sampleKey); !ok {
		t.Errorf("%s should not have been removed", sampleKey)
	}
}
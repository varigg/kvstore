package kvstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	kvstore "github.com/varigg/kvstore/internal"
)

func TestScenario(t *testing.T) {
	store := kvstore.NewKVStore()
	store.Write("a", "hello")
	result := store.Read("a")
	assert.Equal(t, "hello", result)
	store.Start()
	store.Write("a", "hello-again")
	result = store.Read("a")
	assert.Equal(t, "hello-again", result)
	store.Start()
	store.DeleteKey("a")
	result = store.Read("a")
	assert.Equal(t, nil, result)
	err := store.Commit()
	assert.Nil(t, err)
	result = store.Read("a")
	assert.Equal(t, nil, result)
	store.Write("a", "once-more")
	result = store.Read("a")
	assert.Equal(t, "once-more", result)
	err = store.Abort()
	assert.Nil(t, err)
	result = store.Read("a")
	assert.Equal(t, "hello", result)

}

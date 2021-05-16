package kvstore

import (
	"fmt"
)

type KVStore struct {
	currentStore kvstore
	stack        []kvstore
}
type kvstore map[string]interface{}

//  READ Reads and prints, to stdout, the val associated with key. If the value is not present an error is printed to stderr.
//  WRITE Stores val in key.
//  DELETE Removes a key from store. Future READ commands on that key will return an error.
//  START Start a transaction.
//  COMMIT Commit a transaction. All actions in the current transaction are committed to the parent transaction or the root store.
//   If there is no current transaction an error is output to stderr.
//  ABORT Abort a transaction. All actions in the current transaction are discarded.
//  QUIT Exit the REPL cleanly. A message to stderr may be output

func NewKVStore() *KVStore {
	s := KVStore{}
	s.currentStore = make(kvstore)
	return &s
}

func (s *KVStore) GetStackLevel() int {
	return len(s.stack)
}
func (s *KVStore) Start() {
	snapshot := kvstore{}
	for k, v := range s.currentStore {
		snapshot[k] = v
	}
	s.stack = append(s.stack, snapshot)
}

func (s *KVStore) Commit() error {
	if s.GetStackLevel() == 0 {
		return fmt.Errorf("not in a transaction")
	}
	//snapshot := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	// for k, v := range snapshot {
	// 	if s
	// 	s.currentStore[k] = v
	// }

	return nil
}

func (s *KVStore) Abort() error {
	if s.GetStackLevel() == 0 {
		return fmt.Errorf("not in a transaction")
	}
	s.currentStore = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return nil
}

func (s *KVStore) Write(key string, value string) {
	s.currentStore[key] = value
}

func (s *KVStore) DeleteKey(key string) {
	delete(s.currentStore, key)
}

func (s *KVStore) Read(key string) interface{} {
	return s.currentStore[key]
}

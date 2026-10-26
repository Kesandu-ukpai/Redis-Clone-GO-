package main

import "fmt"
//create a store
type Store struct{
	data map[string]string
}
// make the store work
func NewStore() *Store{
	return &Store{
		data: make(map[string]string)
	}
}
// set a key value pair
func (s *Store) Set(key string, value string){
	s.data[key] = value
}



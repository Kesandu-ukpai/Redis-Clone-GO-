package main

import "fmt"

//create a store
type Store struct{
	data map[string]string
}
// make the store work
func NewStore() *Store{
	return &Store{
		data: make(map[string]string),
	}
}
// set a key value pair
func (s *Store) Set(key string, value string){
	s.data[key] = value
}
// get a value by key
func(s *Store) Get(key string)(string, error){
	value, ok := s.data[key]
	if !ok{
		return "", fmt.Errorf("key not found: %s", key)
	}
	return value, nil
}

func (s *Store) Delete(key string) error{
	_, ok := s.data[key]
if !ok{
	return fmt.Errorf("key not found: %s", key)
}
delete(s.data, key)
return nil
}

//exist what will the algo look like
/**1) you will need to check the key exists in the map. so a method will be needed
	2) you will need to return a boolean after checking the map*/
// func (s *Store) Exist(key string) boolean{
// 	_, ok := s.data[key]
// }
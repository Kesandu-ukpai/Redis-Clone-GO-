package main
 
import "fmt"

func main(){
	s:= NewStore()
	s.Set("name", "John")
	s.Set("age", "30")
	name, err := s.Get("name")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	age, err := s.Get("age")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
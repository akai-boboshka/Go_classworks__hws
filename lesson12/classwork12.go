// package main

// import "fmt"

// func main() {
// m := make(map[string]int)
// m["apple"] = 10
// m["banana"] = 19
// fmt.Println(len(m))

// m := map[int]string{
// 	10: "Hello",
// 	20: "World",
// 	30: "How are you?",
// }
// for key, value := range m {
// 	fmt.Println(key, value)
// }
// fmt.Println(m) // map[key1:10 key2:20]

// original := map[string]int{"key1": 10, "key2": 20}
// copy := make(map[string]int)

// for k, v := range original {
// 	copy[k] = v
// }
// fmt.Println(copy)

// type Person struct {
// 	Name string
// 	Age  int
// }

// m := make(map[string]*Person)
// m["key1"] = &Person{Name: "Ibragim", Age: 18}
// m["key2"] = &Person{Name: "Samad", Age: 17}

// m["key1"].Name = "Boboshka"
// fmt.Println(m["key1"].Name)

// Поиск одинаковых элементов в map

// set := make(map[string]struct{})
// set["key1"] = struct{}{}
// set["key2"] = struct{}{}
// _, exists := set["key3"]
// fmt.Println(exists)

// Изменение структуры во время итерации
// m := map[string]int{
// 	"key1": 10,
// 	"key2": 20,
// 	"key3": 30,
// }
// keysToDelete := []string{}
// for key := range m {
// 	keysToDelete = append(keysToDelete, key)
// }
// for _, key := range keysToDelete {
// 	delete(m, key)
// }
// fmt.Println(m)

// reallocation
// m := map[string]int{
// 	"key1": 10,
// 	"key2": 20,
// 	"key3": 30,
// }
// delete(m, "key1")
// fmt.Println(m)
// newMap := make(map[string]int, len(m))

// for k, v := range m {
// 	newMap[k] = v
// }
// m = newMap
// fmt.Println(m)

// Pointers
// type Person struct {
// 	Name string
// 	Age  int
// }

// m := make(map[string]*Person)
// m["user1"] = &Person{Name: "Ibragim", Age: 18}
// m["user2"] = &Person{Name: "Said", Age: 21}

// m["user1"].Age = 25
// fmt.Println(m["user1"])

// }

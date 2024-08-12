package main

// func main() {
// #2
// m := map[string]int{
// 	"key1": 10,
// 	"key2": 20,
// 	"key3": 30,
// }

// _, exists := m["key2"]
// fmt.Println(exists)

// #3
// type User struct {
// 	Name string
// 	Age  int
// 	ID   int
// }

// m := make(map[string]*User)
// m["firstUser"] = &User{Name: "Nikhod", Age: 17, ID: 10}
// m["secondUser"] = &User{Name: "Ibragim", Age: 18, ID: 99}
// fmt.Println("Before changing value:", *m["firstUser"])

// m["firstUser"].ID = 20
// fmt.Println("After changing value:", *m["firstUser"])

// #4
// ourMap := map[string]int{
// 	"Apple":     20,
// 	"Banana":    10,
// 	"Pineapple": 50,
// }
// fmt.Println("Map before delete the key:", ourMap)

// removeFromMap(ourMap, "Pineapple")
// fmt.Println("Map after delete the key:", ourMap)

// #5
// text := "hello world hello helo world everyone"
// frequency := countWordFrequency(text)

// fmt.Println(frequency)

// #6
// sortMap()

// newM := checkMapToHaveKey("key1")
// fmt.Println(newM)

// #16
// returnSlice()

// #18
// var baseMap map[string]int
// baseMap = map[string]int{
// 	"firstKey":  123,
// 	"secondKey": 456,
// }
// newMap := CreateCopyOfMap(baseMap)
// fmt.Println(newMap)

// #17
// 	newMap := map[string]string{
// 		"item1": "Hello",
// 		"item2": "World",
// 		"item3": "How are you?",

// }

// 	returnFunc := f(newMap)
// 	fmt.Println(returnFunc)
// }

// #4
// func removeFromMap(m map[string]int, key string) {
// 	delete(m, key)
// }

// #5
// func countWordFrequency(text string) map[string]int {
// 	// Создаем пустую карту для хранения частоты слов
// 	wordFrequency := make(map[string]int)

// 	// Разбиваем строку на слова, используя пробел как разделитель
// 	words := strings.Fields(text)

// 	// Проходим по всем словам и считаем их частоту
// 	for _, word := range words {
// 		wordFrequency[word]++
// 	}

// 	return wordFrequency
// }

// #6

// func sortMap() /*map[string]int*/ {
// 	m := map[string]int{"keyOne": 111, "keyTwo": 222}

// 	type Keys struct {
// 		Key   string
// 		Value int
// 	}
// 	var sortedSlice []Keys
// 	for key, value := range m {
// 		sortedSlice = append(sortedSlice, Keys{key, value})
// 	}
// 	fmt.Println(sortedSlice[1])

// }
//
// #7
// func checkMapToHaveKey(key string) map[string]int {
// 	m := map[string]int{
// 		"key1": 20,
// 		"key2": 30,
// 	}
// 	_, exists := m[key]
// 	fmt.Println(exists)

// 	return m
// }

// #16
// func returnSlice() {
// 	m := map[string]int{"keyOne": 111, "keyTwo": 222}
// 	type Keys struct {
// 		Key   string
// 		Value int
// 	}
// 	var sortedSlice []Keys
// 	for key, value := range m {
// 		sortedSlice = append(sortedSlice, Keys{key, value})
// 	}
// 	fmt.Println(sortedSlice[1])
// }

// #18
// func CreateCopyOfMap(baseMap map[string]int) map[string]int {
// 	copy := make(map[string]int)
// 	for k, v := range baseMap {
// 		copy[k] = v
// 	}
// 	return copy
// }

// #17
// func f(m map[string]string) string {
// 	values := make([]string, 0, len(m))

// 	for _, value := range m {
// 		values = append(values, value)
// 	}

// 	return strings.Join(values, ", ")
// }

// type Fruit struct {
// 	Name  string
// 	Price int
// }

// func main() {

// 	m := make(map[string]*Fruit)
// 	m["firstItem"] = &Fruit{"Apple", 20}
// 	m["seconditem"] = &Fruit{"Pear", 18}

// 	m["firstItem"] = &Fruit{"Pineapple", 20}
// 	fmt.Println(m["firstItem"])
// }

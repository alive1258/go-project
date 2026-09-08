package main

func main() {

	myMap := map[string]string{
		"name":    "Jon",
		"address": "Winterfell",
	}

	for key, value := range myMap {
		println(key, value)
	}

	myArray := [3]string{"Jon", "Arya", "Sansa"}
	for index, value := range myArray {
		println(index, value)
	}

	name := "Zamirul"
	for index, value := range name {
		println(index, string(value))
	}

	var byteSlice []byte = []byte(name)
	for index, value := range byteSlice {
		println(index, string(value))
	}

}
// Hoang Hai Ha
package main

import (
	"encoding/json"
	"fmt"
	student "github.com/81300988/GolangNordic/type_system/sutdent"
)

func main() {
	dataJson := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
		]`
	aString := "hello world"

	bs := []byte(aString)
	var arrStudent []student.Student
	_ = json.Unmarshal([]byte(dataJson), &arrStudent)
	fmt.Println(arrStudent)
	fmt.Print(bs)

}

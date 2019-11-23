// Hoang Hai Ha
package main

import (
	"github.com/81300988/GolangNordic/tree/master/type_system/sutdent"
	"encoding/json"
	"fmt"
)

func main() {
	dataJson := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
		]`

	var arrStudent []Student
	_ = json.Unmarshal([]byte(dataJson), &arrStudent)
	fmt.Println(arrStudent)

}

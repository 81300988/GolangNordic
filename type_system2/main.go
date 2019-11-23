// Hoang Hai Ha
package Student

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FirstName string `json:"first_name" bson:"full_name" validate:"required"`
	LastName  string `json:"last_name" bson:"full_name" validate:"required"`
	Age       int
	ClassName string `json:"class_name"`
}

func main() {
	dataJson := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
		]`

	var arrStudent []Student
	_ = json.Unmarshal([]byte(dataJson), &arrStudent)
	fmt.Println(arrStudent)

}

// Hoang Hai Ha
package student

type Student struct {
	FirstName string `json:"first_name" bson:"full_name" validate:"required"`
	LastName  string `json:"last_name" bson:"full_name" validate:"required"`
	Age       int
	ClassName string `json:"class_name"`
}

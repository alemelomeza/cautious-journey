package dto

type ChildDTO struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Parents []ParentDTO `json:"parents"`
}
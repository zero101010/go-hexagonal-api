package domain

type User struct {
	ID   uint   `json:"id" gorm:"primary_key;autoIncrement:true"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

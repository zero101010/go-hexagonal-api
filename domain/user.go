package domain

type User struct {
	// gorm.Model
	ID   uint   `json:"id" gorm:"primary_key;autoIncrement:true"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

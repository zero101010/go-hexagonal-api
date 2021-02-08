package domain

type Quote struct {
	ID       uint    `json:"id" gorm:"primary_key;autoIncrement:true"`
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
	User     *User   `gorm:"-"`
}

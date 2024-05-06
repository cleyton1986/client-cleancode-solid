package entities

type User struct {
    ID        uint   `gorm:"primary_key" json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
}

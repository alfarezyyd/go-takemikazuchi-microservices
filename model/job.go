package model

type Job struct {
	ID          uint64  `gorm:"column:id;autoIncrement;primaryKey"`
	UserId      uint64  `gorm:"column:user_id"`
	CategoryId  uint64  `gorm:"column:category_id"`
	Title       string  `gorm:"column:title"`
	Description string  `gorm:"column:description"`
	Location    string  `gorm:"column:location"`
	Latitude    float64 `gorm:"column:latitude"`
	Longitude   float64 `gorm:"column:longitude"`
	Address     string  `gorm:"column:address"`
	PlaceId     string  `gorm:"column:place_id"`
	Price       float64 `gorm:"column:price"`
	Status      string  `gorm:"column:status"`
	CreatedAt   string  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   string  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

package model

type JobResource struct {
	ID        uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	ImagePath string `gorm:"column:image_path"`
	VideoUrl  string `gorm:"column:video_url"`
	JobId     uint64 `gorm:"column:job_id"`
	Job       Job    `gorm:"foreignKey:job_id;references:id"`
}

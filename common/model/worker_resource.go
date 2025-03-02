package model

type WorkerResource struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	WorkerId uint64 `gorm:"column:worker_id"`
	FilePath string `gorm:"column:file_path"`
	Type     string `gorm:"column:type"`
	Worker   Worker `gorm:"foreignKey:worker_id;references:id"`
}

package dto

type UpdateJobDto struct {
	CategoryId       uint64   `form:"category_id"`
	Title            string   `form:"title"`
	Description      string   `form:"description"`
	Price            float64  `form:"price"`
	DeletedFilesName []string `form:"deleted_files_name"`
}

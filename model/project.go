package model

type Project struct {
	Id                  uint   `form:"id" json:"id" gorm:"AUTO_INCREMENT"`
	Title               string `form:"title" json:"title"`
	Project_description string `form:"project_description" json:"project_description"`
	Project_date        string `form:"project_date" json:"project_date"`
	Project_type        string `form:"project_type" json:"project_type"`
	User_id             int    `form:"user_id" json:"user_id"`
}

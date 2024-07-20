package data

type StudentQuery struct {
	StudentQuestion     string `json:"student_query" binding:"required"`
	StudentCurrentLevel int    `json:"student_current_level" binding:"required"`
}

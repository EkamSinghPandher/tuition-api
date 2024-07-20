package data

type TeacherResponse struct {
	TeacherResponse string `json:"teacher_response" binding:"required"`
}

/*
Contains the body fields for the request to the teacher (LLM). Will be serialised to json.
*/
type TeacherQueryRequest struct {
	Model    string                `json:"model" binding:"required"`
	Messages []TeacherQueryContent `json:"messages" binding:"required"`
}

type TeacherQueryContent struct {
	Role    string `json:"role" binding:"required"`
	Content string `json:"content" binding:"required"`
}

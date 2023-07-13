package dto

type CreatePost struct {
	Title  string `json:"title" binding:"required,min=5"`
	Content     string `json:"content" binding:"min=6"`
	AuthorId uint   `json:"authorId" binding:"required,min=1,max=100"`
}


type UpdatePost struct {
	Title  string `json:"title" binding:"min=5"`
	Content     string `json:"content" binding:"min=6"`
}
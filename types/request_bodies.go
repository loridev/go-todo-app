package types

type TodoCreateRequestBody struct {
	Title string `json:"title" binding:"required,alphanum"`
}

type TodoUpdateRequestBody struct {
	Title string `json:"title" binding:"omitempty,alphanum"`
	Done  *bool  `json:"done" binding:"omitempty,boolean"`
}

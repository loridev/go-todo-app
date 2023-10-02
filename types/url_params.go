package types

type TodoUrlParams struct {
	Id string `uri:"id" binding:"required,numeric"`
}

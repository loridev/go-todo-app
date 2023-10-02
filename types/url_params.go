package types

type TodoURLParams struct {
	ID string `uri:"id" binding:"required,numeric"`
}

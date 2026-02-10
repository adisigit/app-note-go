package dto

type NoteCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteUpdateRequest struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

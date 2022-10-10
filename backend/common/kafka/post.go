package kafka_messages

type PostElement struct {
	Text   *string `json:"text,omitempty"`
	Title  *string `json:"title,omitempty"`
	Photo  *[]byte `json:"photo_file_name,omitempty"`
	IsText bool    `json:"is_text"`
}

type Post struct {
	Elements []PostElement `json:"elements"`
	UserID   int           `json:"user_id"`
	ID       int           `json:"id"`
}

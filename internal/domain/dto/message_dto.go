package dto

type SendMessageDTO struct {
	Message  string `json:"message"`
	StreamID string `json:"stream_id"`
}

type FetchMessageDTO struct {
	StreamID string `json:"stream_id"`
}

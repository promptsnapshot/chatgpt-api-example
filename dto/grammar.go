package dto

type ErrorDetail struct {
	OriginalText string   `json:"original_text"`
	Position     []int    `json:"position"`
	Suggestions  []string `json:"suggestions"`
	ErrorType    string   `json:"error_type"`
}

type GrammarCheckRequestDto struct {
	Paragraph string `json:"paragraph"`
}

type GrammarCheckResponseDto struct {
	CorrectedText string         `json:"corrected_text"`
	ErrorDetails  []*ErrorDetail `json:"error_details"`
}

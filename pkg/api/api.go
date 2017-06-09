package api

// Empty type
type Empty struct{}

// IDRequest type
type IDRequest struct {
	ID int64 `json:"id"`
}

// ListRequest type
type ListRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

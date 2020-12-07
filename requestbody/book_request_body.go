package requestbody

type BookRequestBody struct {
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	CoverImage string `json:"cover_image"`
	Amount	int32	`json:"amount"`
	SKU	string `json:"sku"`
	TypeID uint `json:"book_type_id" `
}

type BookRequestPutBody struct {
	ID uint `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	CoverImage string `json:"cover_image"`
	Price	float64	`json:"price"`
	Amount	int32	`json:"amount"`
	SKU	string `json:"sku"`
	BookTypeID uint `json:"book_type_id" `
}

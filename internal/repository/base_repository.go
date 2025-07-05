package repository

func limitOffset(page, perPage uint) (int, int) {
	limit := int(perPage)
	offset := (int(page) - 1) * limit
	return limit, offset
}

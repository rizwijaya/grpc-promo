package errorVal

func ErrNotFound(data string) string {
	return "Data " + data + " tidak ditemukan."
}

func ErrInvalidArgument(data string) string {
	return data + " yang anda inputkan tidak valid."
}

func ErrInvalidUUID(data string) string {
	return "pq: invalid input syntax for type uuid: \"" + data + "\""
}

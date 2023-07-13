package errorVal

func ErrNotFound(data string) string {
	return "Data " + data + " tidak ditemukan."
}

func ErrInvalidArgument(data string) string {
	return "Data " + data + " yang anda inputkan tidak valid."
}

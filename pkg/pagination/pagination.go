package pagination

import (
	"math"
	promoProto "promo/modules/protobuf/pb"
)

func Pages(total int64, page int64, limit int64) (int64, *promoProto.Pagination) {
	var (
		pagin  promoProto.Pagination
		offset int64
	)

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	pagin.Total = uint64(total)
	pagin.PerPage = uint32(limit)
	pagin.LastPage = uint32(math.Ceil(float64(total) / float64(limit)))
	if pagin.LastPage < uint32(page) {
		pagin.CurrentPage = uint32(pagin.LastPage)
		offset = (int64(pagin.LastPage) - 1) * limit
	} else {
		pagin.CurrentPage = uint32(page)
	}
	return offset, &pagin
}

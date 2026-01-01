package middlewares

import (
	"fmt"
	"gamecheck-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

const DEFAULT_PAGE_LIMIT = 100
const DEFAULT_SIZE_LIMIT = 50

type PaginationParams struct {
	PageLimit uint
	SizeLimit uint
}

func AcceptPaginationParamsMiddleware(limits PaginationParams) gin.HandlerFunc {

	var pageLimit uint = limits.PageLimit
	var sizeLimit uint = limits.SizeLimit

	if pageLimit == 0 {
		pageLimit = DEFAULT_PAGE_LIMIT
	}

	if sizeLimit == 0 {
		sizeLimit = DEFAULT_SIZE_LIMIT
	}

	return func(ctx *gin.Context) {
		page, pageExists := ctx.GetQuery("page")
		size, sizeExists := ctx.GetQuery("size")

		if pageExists != sizeExists { // if either one param is passed
			utils.GinBadResponse(ctx, "Page and Size params are both required or both omitted")
			return
		}

		if pageExists == false && sizeExists == false { // if both params aren't passed
			ctx.Next()
			return
		}

		pageNumber, err := strconv.ParseUint(page, 10, 16)
		sizeNumber, err := strconv.ParseUint(size, 10, 16)

		if err != nil || pageNumber == 0 || sizeNumber == 0 {
			utils.GinBadResponse(ctx, "Invalid values for page or size params")
			return
		}

		if uint(pageNumber) > pageLimit {
			utils.GinBadResponse(ctx, fmt.Sprintf("Page exceeds limit of %d", pageLimit))
			return
		}

		if uint(sizeNumber) > sizeLimit {
			utils.GinBadResponse(ctx, fmt.Sprintf("Size exceeds limit of %d", sizeLimit))
			return
		}

		ctx.Set("page", uint(pageNumber))
		ctx.Set("size", uint(sizeNumber))

		ctx.Next()
	}
}

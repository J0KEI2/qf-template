package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/kku-qf-services/pkg/models"
)

func ExtractPaginationOption(c *fiber.Ctx) models.PaginationOptions {
	options := models.PaginationOptions{}
	options.SetPage(c.QueryInt("page", 1))
	options.SetLimit(c.QueryInt("limit", 0))
	options.SetOrder(c.Query("order", "created_at desc"))
	options.SetSearch(c.Query("search", ""))
	options.DefaultTotal()
	options.DefaultTotalPage()
	return options
}

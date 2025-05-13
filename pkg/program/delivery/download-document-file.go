package delivery

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	// helpers "github.com/zercle/gofiber-helpers"
)

func (h programHandler) DownloadDocumentFile() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		fileName := c.Params("fileName")
		downloadPath := fmt.Sprintf("%v/%v", viper.GetString("file.document_path"), fileName)
		if viper.GetBool("debug_mode") {
			fmt.Println("\n downloadPath: ", downloadPath)
		}

		return c.SendFile(downloadPath)
	}
}

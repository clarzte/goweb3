package merchants

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MerchantsRoutes struct {
	e                *echo.Echo
	merchantsService MerchantsServiceActions
}

type ApiError struct {
	Error string `json:"error"`
}

func NewMerchantsRoutes(e *echo.Echo, merchantsService MerchantsServiceActions) *MerchantsRoutes {
	r := &MerchantsRoutes{e, merchantsService}

	r.getMerchants()

	return r
}

func (r *MerchantsRoutes) getMerchants() {
	r.e.GET("/merchants", func(c echo.Context) error {
		lat, latErr := convertStringToFloat(c.QueryParam("lat"))
		lng, lngErr := convertStringToFloat(c.QueryParam("lng"))
		radius, radErr := convertStringToFloat(c.QueryParam("radius")) //in meters
		keyword := c.QueryParam("keyword")

		if latErr != nil || lngErr != nil || radErr != nil {
			return c.JSON(400, &ApiError{"Invalid query parameters"})
		}

		merchants, err := r.merchantsService.GetMerchants(lat, lng, radius, keyword)

		if err != nil {
			fmt.Printf("Error %s", err.Error())

			return c.JSON(500, &ApiError{"Internal server error"})
		}

		return c.JSON(200, merchants)
	})
}

func convertStringToFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

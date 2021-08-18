package kawalcovid

import (
	"main-backend/bussiness/kawalcovid"
	"main-backend/controller"
	_ "main-backend/controller/kawalcovid/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type KawalCovidController struct {
	kawalCovidUC kawalcovid.Usecase
}

func NewKawalCovidController(e *echo.Echo, cu kawalcovid.Usecase) *KawalCovidController {
	return &KawalCovidController{
		kawalCovidUC: cu,
	}
}

func (ctrl *KawalCovidController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.kawalCovidUC.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

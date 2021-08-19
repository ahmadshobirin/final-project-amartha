package clinic

import (
	"errors"
	"main-backend/bussiness/clinic"
	"main-backend/controller"
	"main-backend/controller/clinic/request"
	"main-backend/controller/clinic/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type ClinicController struct {
	clinicUC clinic.Usecase
}

func NewClinicController(e *echo.Echo, cu clinic.Usecase) *ClinicController {
	return &ClinicController{
		clinicUC: cu,
	}
}

func (ctrl *ClinicController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	cityID, _ := strconv.Atoi(c.QueryParam("city_id"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	resp, count, err := ctrl.clinicUC.Fetch(ctx, cityID, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Clinic{}
	for _, value := range resp {
		responseController = append(responseController, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponseWithTotal(c, responseController, count)
}

func (ctrl *ClinicController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.clinicUC.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&result))

}

func (ctrl *ClinicController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Clinic{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.clinicUC.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

func (ctrl *ClinicController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Clinic{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	err := ctrl.clinicUC.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

func (ctrl *ClinicController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Clinic{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	err := ctrl.clinicUC.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

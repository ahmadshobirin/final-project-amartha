package city

import (
	"errors"
	"main-backend/bussiness/city"
	"main-backend/controller"
	"main-backend/controller/city/request"
	"main-backend/controller/city/response"
	"main-backend/helper/str"
	"strconv"
	"strings"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CityController struct {
	cityUsecase city.Usecase
}

func NewCityController(e *echo.Echo, cu city.Usecase) *CityController {
	return &CityController{
		cityUsecase: cu,
	}
}

func (ctrl *CityController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.cityUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.City{}
	for _, value := range resp {
		responseController = append(responseController, *response.FromDomain(&value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *CityController) FindByID(c echo.Context) error {
	ctx := c.Request().Context()
	ID := str.StringToInt(c.Param("id"))

	resp, err := ctrl.cityUsecase.FindByID(ctx, ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(&resp))
}

func (ctrl *CityController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.City{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.cityUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

func (ctrl *CityController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.City{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	err := ctrl.cityUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

func (ctrl *CityController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.City{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	err := ctrl.cityUsecase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

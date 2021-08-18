package queue

import (
	"errors"
	"main-backend/app/middleware"
	"main-backend/bussiness/queue"
	"main-backend/controller"
	"main-backend/controller/queue/request"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type QueueController struct {
	queueUC queue.Usecase
	jwtAuth *middleware.ConfigJWT
}

func NewQueueController(e *echo.Echo, cu queue.Usecase, jwt *middleware.ConfigJWT) *QueueController {
	return &QueueController{
		queueUC: cu,
		jwtAuth: jwt,
	}
}

func (ctrl *QueueController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	user := middleware.GetUser(c)
	req := request.Queue{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.queueUC.Store(ctx, user.ID, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

func (ctrl *QueueController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Queue{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	err := ctrl.queueUC.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil)
}

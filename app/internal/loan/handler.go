package loan

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	ApplyForLoan(c echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

// ApplyForLoan handles POST /api/v1/loans
func (hdl *handler) ApplyForLoan(c echo.Context) error {
	var req LoanApplication

	// Bind JSON request body to struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
	}

	// Validate request
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := hdl.service.ProcessLoanApplication(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

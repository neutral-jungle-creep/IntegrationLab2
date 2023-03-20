package v1

import (
	"IntegrationLab2/internal/domain"
	"IntegrationLab2/pkg/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

const nAPI = "http://testruz.agtu.ru/RUZService/RUZService.svc/"

type Handler struct {
	log *logger.Logger
}

func NewHandler(log *logger.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) getFaculties(c *gin.Context) {
	var faculties []domain.Faculty

	responseText := httpGetRequest("faculties")
	if err := json.Unmarshal(responseText, &faculties); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.log.Infof("get faculties result = %v", faculties)

	c.JSON(http.StatusOK, faculties)
}

func httpGetRequest(args string) []byte {
	resp, err := http.Get(nAPI + args)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return respText
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, message)
}

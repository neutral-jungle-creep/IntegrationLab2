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

func (h *Handler) getGroups(c *gin.Context) {
	var groups []domain.Group

	responseText := httpGetRequest("groups")
	if err := json.Unmarshal(responseText, &groups); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.log.Infof("get groups result = %v", groups)

	c.JSON(http.StatusOK, groups)
}

func (h *Handler) getLessons(c *gin.Context) {
	var lessons []domain.Lesson
	fromDate, toDate := c.Query("fromdate"), c.Query("todate")
	groupOid := c.Query("groupOid")

	var args = "lessons?" + "fromdate=" + fromDate + "&todate=" + toDate + "&groupOid=" + groupOid

	responseText := httpGetRequest(args)
	if err := json.Unmarshal(responseText, &lessons); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.log.Infof("get lessons result = %v", lessons)

	c.JSON(http.StatusOK, lessons)
}

func httpGetRequest(args string) []byte {
	log.Println(args)
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

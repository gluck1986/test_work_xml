package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/datasource/criteria"
	"gluck1986/test_work_xml/internal/model"
	"log"
	"net/http"
	"strings"
)

// GetNamesHandler echo handler
type GetNamesHandler struct {
	logger *log.Logger
	repo   datasource.ISdnRepository
}

// NewGetNamesHandler constructor
func NewGetNamesHandler(
	repo datasource.ISdnRepository,
	logger *log.Logger,
) *GetNamesHandler {
	return &GetNamesHandler{logger: logger, repo: repo}
}

type GetNamesResultItem struct {
	Uid       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Handle echo handler function
func (t *GetNamesHandler) Handle(ctx echo.Context) error {
	names := ctx.QueryParam("name")
	names = strings.TrimSpace(names)
	pType := ctx.QueryParam("type")
	pType = strings.TrimSpace(pType)
	pType = strings.ToLower(pType)

	repoCriteria := criteria.SdnCriteria{}

	if len(names) < 1 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Errorf("name is required"))
	}
	repoCriteria.MaybeFirstName, repoCriteria.MaybeLastName = t.mapNames(names)
	repoCriteria.Mode = t.mapType(pType)
	models, err := t.repo.ReadMany(repoCriteria)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, t.mapResult(models))
}

func (t *GetNamesHandler) mapNames(names string) (string, string) {
	namesSplited := strings.Split(names, " ")
	namesFiltered := make([]string, 2)
	i := 0
	for _, name := range namesSplited {
		if i == 2 {
			break
		}
		if name != "" {
			namesFiltered[i] = name
			i++
		}
	}
	return namesFiltered[0], namesFiltered[1]
}

func (t *GetNamesHandler) mapType(src string) criteria.SdnMode {
	switch src {
	case "strong":
		return criteria.SdnModeStrong
	case "weak":
		return criteria.SdnModeWeak
	default:
		return criteria.SdnModeWeak
	}
}

func (t *GetNamesHandler) mapResult(src []model.SdnEntity) []GetNamesResultItem {
	result := make([]GetNamesResultItem, len(src))

	for i, sdnEntity := range src {
		result[i] = GetNamesResultItem{
			Uid:       sdnEntity.UID,
			FirstName: sdnEntity.FirstName,
			LastName:  sdnEntity.LastName,
		}
	}
	return result
}

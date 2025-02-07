package http

import (
	"strconv"

	"github.com/freitzzz/gameboy-db-api/internal/errors"
	"github.com/freitzzz/gameboy-db-api/internal/logging"
	"github.com/freitzzz/gameboy-db-api/internal/model"
	"github.com/labstack/echo/v4"
)

func registerHandlers(e *echo.Echo) {
	// /previews
	e.GET(previewsRoute, previewsHandler)

	// /details
	e.GET(detailsIdRoute, detailsHandler)

	e.HTTPErrorHandler = httpErrorHandler()
}

func previewsHandler(ectx echo.Context) error {
	return withServiceContainer(ectx, func(sc serviceContainer) error {
		r := ratingFilter(ectx.QueryParam(ratingFilterQueryParam))
		n := ectx.QueryParam(nameFilterQueryParam)
		pq := ectx.QueryParam(pageFilterQueryParam)

		if len(n) == 0 {
			if r == ratingFilterLowestRated {
				return callAndReply(ectx, sc.Games.LowestRated)
			}

			return callAndReply(ectx, sc.Games.HighestRated)
		}

		p, err := strconv.ParseInt(pq, 10, 32)
		if err != nil {
			logging.Warning("failed to parse page(%s), defaulting to 1, %v", p, err)
			p = 1
		}

		return callAndReply(ectx, func() ([]model.GamePreview, error) {
			return sc.Games.Search(int(p), n)
		})
	})
}

func detailsHandler(ectx echo.Context) error {
	return withServiceContainer(ectx, func(sc serviceContainer) error {
		id, err := strconv.ParseInt(ectx.Param("id"), 10, 32)
		if err != nil {
			logging.Error("failed to parse id (%s), %v", ectx.Param("id"), err)
			return echo.ErrNotFound
		}

		game, err := sc.Games.Find(int(id))
		if err != nil {
			return err
		}

		return ectx.JSON(200, game)
	})
}

func httpErrorHandler() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		// make sure to not process any false positives
		if err == nil {
			return
		}

		logging.Error("handling error... %v", err)

		if err == errors.ErrRecordNotFound {
			c.Response().WriteHeader(404)
			return
		}

		he, ok := err.(*echo.HTTPError)

		// If all cast fail, serve fallback
		if !ok {
			he = echo.NewHTTPError(500, err)
		}

		c.Response().WriteHeader(he.Code)
	}
}

package newscontroller

import (
	"net/http"
	"strconv"

	"github.com/firmanmimang/go-api-trusted-news/helpers"
	"github.com/firmanmimang/go-api-trusted-news/middlewares"
	"github.com/firmanmimang/go-api-trusted-news/services"
)

func Index(w http.ResponseWriter, r *http.Request) {
	db := middlewares.GetDB(r.Context())
	if db == nil {
		helpers.Response(w, http.StatusInternalServerError, "Failed to connect db", nil)
		return
	}

	// Get pagination parameters
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	// Set defaults if not provided
	if limit == 0 {
		limit = 10 // Default limit per page
	}
	if offset < 0 {
		offset = 0
	}

	// Get category slug (optional)
	var categorySlug *string
	if slug := r.URL.Query().Get("category"); slug != "" {
		categorySlug = &slug
	}

	userService := services.NewNewsService(db)
	news, err := userService.GetNews(limit, offset, categorySlug)
	if err != nil {
		helpers.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.Response(w, http.StatusOK, "List News", news)
}

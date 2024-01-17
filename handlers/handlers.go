package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/devhulk/gooser-web/components"
	"github.com/devhulk/gooser-web/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	component := components.Home()
	component.Render(context.Background(), w)
}

func GetSites(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("username")
	result, err := services.GetSiteMap(u)
	if err != nil {
		log.Fatalln(err)
	}

	component := components.Sites(result)
	component.Render(context.Background(), w)
}

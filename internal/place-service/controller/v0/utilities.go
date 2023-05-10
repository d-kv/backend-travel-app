package ictrl_v0

import (
	"golang.org/x/exp/slices"

	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
)

func filterByCategory(pls []model.Place, c *category.Category) []model.Place {
	if c.IsAnyCategory() {
		return pls
	}

	var places []model.Place
	var isAdded bool

	for _, p := range pls {

		isAdded = false

		for _, c := range c.Main {
			if slices.Contains(p.Category.Main, c) {
				places = append(places, p)
				isAdded = true

				break
			}
		}

		if isAdded {
			continue
		}

		for _, cat := range c.Sub {
			if slices.Contains(p.Category.Sub, cat) {
				places = append(places, p)
			}
		}

	}

	return places
}

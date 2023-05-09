package ictrl_v0

import (
	"github.com/d-kv/backend-travel-app/pkg/place-service/model"
	"github.com/d-kv/backend-travel-app/pkg/place-service/model/category"
	"golang.org/x/exp/slices"
)

func filterByCategory(pls []model.Place, m []category.Main, s []category.Sub) []model.Place {
	var places []model.Place
	var isAdded bool

	for _, p := range pls {

		isAdded = false

		for _, c := range m {
			if slices.Contains(p.Category.Main, c) {
				places = append(places, p)
				isAdded = true

				break
			}
		}

		if isAdded {
			continue
		}

		for _, cat := range s {
			if slices.Contains(p.Category.Sub, cat) {
				places = append(places, p)
			}
		}

	}

	return places
}

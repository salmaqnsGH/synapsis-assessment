package helper

import (
	"salmaqnsGH/sysnapsis-assessment/model/domain"
	"salmaqnsGH/sysnapsis-assessment/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

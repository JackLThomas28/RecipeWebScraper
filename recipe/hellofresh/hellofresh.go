package recipe

import (
	"main/myutils"
	"golang.org/x/net/html"
	"encoding/json"
	"log"
)

type Nutrition struct {
	Type string `json:"@type"`
	Calories string `json:"calories"`
	FatContent string `json:"fatContent"`
	SaturatedFatContent string `json:"saturatedFatContent"`
	CarbohydrateContent string `json:"carbohydrateContent"`
	SugarContent string `json:"sugarContent"`
	ProteinContent string `json:"proteinContent"`
	FiberContent string `json:"fiberContent"`
	CholesterolContent string `json:"cholesterolContent"`
	SodiumContent string `json:"sodiumContent"`
}

type Step struct {
	Type string `json:"@type"`
	Text string `json:"text"`
}

type Recipe struct {
	Content string `json:"@context"`
	Type string `json:"@type"`
	Name string `json:"name"`
	Author string `json:"author"`
	Image string `json:"image"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Description string `json:"description"`
	DatePublished string `json:"datePublished"`
	TotalTime string `json:"totalTime"`
	Nutrition Nutrition `json:"nutrition"`
	RecipeInstructions []Step `json:"recipeInstructions"`
	RecipeIngredient []string `json:"recipeIngredient"`
	RecipeYield int `json:"recipeYield"`
	Keywords []string `json:"keywords"`
	RecipeCategory string `json:"recipeCategory"`
	RecipeCuisine string `json:"recipeCuisine"`
}

func GetRecipe(node *html.Node) {
	const ID = "schema-org"

	n := myutils.GetElementById(node, ID)
	var recipe Recipe
	json.Unmarshal([]byte(n.FirstChild.Data), &recipe)

	log.Printf("HelloFresh")
	// The recipe is the second item in the array
	for i,_ := range recipe.RecipeIngredient {
		log.Printf(recipe.RecipeIngredient[i])
	}
}
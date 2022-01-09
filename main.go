package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"encoding/json"
	// hellofresh "main/recipe/hellofresh"
	allrecipes "main/recipe/allrecipes"
)

func main() {
	const FILE_NAME = "allrecipes.json"

	// Read in saved recipes
	body, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatalln(err)
	}
	var savedRecipes []allrecipes.Recipe
	json.Unmarshal([]byte(body), &savedRecipes)

	// Initializations for the first recipe to scrape
	recipeNum := 15195
	index := 0
	const LIMIT = 50

	for index < LIMIT {
		recipeNum += 1

		log.Printf("Recipes scraped so far: %d", index)

		URL := allrecipes.URL + strconv.Itoa(recipeNum) + "/"
		log.Printf("URL: %s", URL)
		// Get the recipe in json form
		recipe, err := allrecipes.GetRecipe(URL)
		if err != nil {
			log.Fatalln(err)
		}

		// Check the URL to see if we've previously saved the recipe
		duplicate := false
		for i,_ := range savedRecipes {
			if recipe.MainEntityOfPage == savedRecipes[i].MainEntityOfPage {
				duplicate = true
			}
		}

		// If we have already saved the recipe, don't re-save
		if !duplicate {
			savedRecipes = append(savedRecipes, recipe)
			index++
		}
	}

	// Save all recipes
	file, _ := json.MarshalIndent(savedRecipes, "", " ")
	_ = ioutil.WriteFile(FILE_NAME, file, 0644)
}
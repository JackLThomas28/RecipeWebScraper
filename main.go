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
	const FILE_NAME = allrecipes.FILE_NAME

	// Read in saved recipes
	body, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatalln(err)
	}
	var savedRecipes []allrecipes.Recipe
	json.Unmarshal([]byte(body), &savedRecipes)

	// Store the saved recipes in a map
	recipeMap := make(map[int]allrecipes.Recipe)
	for i := 0; i < len(savedRecipes); i++ {
		recipeMap[savedRecipes[i].ID] = savedRecipes[i]
	}

	// Initializations for the first recipe to scrape
	recipeID := 15195
	index := 0
	const LIMIT = 5

	for index < LIMIT {
		recipeID += 1

		URL := allrecipes.URL + strconv.Itoa(recipeID) + "/"
		log.Printf("URL: %s", URL)
		// Get the recipe in json form
		recipe, err := allrecipes.GetRecipe(URL, recipeID)
		if err != nil {
			log.Fatalln(err)
		}

		// Check the URL to see if we've previously saved the recipe
		_, duplicate := recipeMap[recipe.ID]

		// If we have already saved the recipe, don't re-save
		if !duplicate {
			savedRecipes = append(savedRecipes, recipe)
			recipeMap[recipe.ID] = recipe
			index++
		}
		log.Printf("Recipes scraped so far: %d", index)
	}

	// Save all recipes
	file, _ := json.MarshalIndent(savedRecipes, "", " ")
	_ = ioutil.WriteFile(FILE_NAME, file, 0644)
}
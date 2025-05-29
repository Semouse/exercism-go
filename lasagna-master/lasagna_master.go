package lasagna

// PreparationTime returns the estimate for the total preparation time
// based on the number of layers
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return avgPrepTime * len(layers)
}

// Quantities returns quantity of noodles and sauce needed to make meal
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// AddSecretIngredient replace "?" ingridient by last ingridient from friendIngredients
func AddSecretIngredient(friendIngredients, ownIngredients []string) {
	for i := 0; i < len(ownIngredients); i++ {
		if ownIngredients[i] == "?" {
			ownIngredients[i] = friendIngredients[len(friendIngredients)-1]
		}
	}
}

// ScaleRecipe return amount of ingredients required for producing n portions
func ScaleRecipe(amounts []float64, portions int) []float64 {
	result := make([]float64, len(amounts))
	for i := 0; i < len(amounts); i++ {
		result[i] = amounts[i] * float64(portions) / 2.0
	}
	return result
}

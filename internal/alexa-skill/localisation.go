package skill

import "github.com/dasjott/alexa-sdk-go"

var Locales = alexa.Localisation{
	"fr-FR": alexa.Translation{
		"WELCOME_MSG":            "Bienvenue sur transport rennes! Vous pouvez me demander quand arrive le prochain bus ?",
		"UPCOMING_ONE_BUS_MSG":   "Le {bus}, arrêt {busstop}, direction {destination}: prochain départ dans {dep1} minutes.",
		"UPCOMING_TWO_BUSES_MSG": "Le {bus}, arrêt {busstop}, direction {destination}: prochain départ dans {dep1} minutes, le suivant dans {dep2} minutes.",
		"TOO_MANY_BUSES":         "Beaucoup de bus sont disponibles à l'arrêt {busstop}. Veuillez préciser votre bus dans la question.",
		"NO_BUS_AVAILABLE":       "Aucun bus {bus} n'est actuellement prévu à l'arrêt {busstop}",
		"FAVORITE_SAVED":         "L'arrêt de bus {busstop} a été ajouté en favori",
		"FAVORITE_DELETED":       "Votre favori a été supprimé",
		"NO_FAVORITE":            "Aucun arrêt de bus n'est enregistré en favori",
		"FAVORITE_UNAVAILABLE":   "Le service de favori est momentanément indisponible",
		"HELP_MSG":               "Je peux vous indiquer quand arrive le prochain bus. Dites simplement: quand arrive le prochain bus ?",
		"GOODBYE_MSG":            "Au revoir!",
		"ERROR_MSG":              "Désolé, je n'ai pas compris. Pouvez-vous reformuler?",
	},
	"en-US": alexa.Translation{
		"WELCOME_MSG":            "Welcome to transport rennes! You can ask me when the next bus is coming?",
		"UPCOMING_ONE_BUS_MSG":   "The {bus}, stop {busstop}, direction {destination}: next departure in {dep1} minutes. ",
		"UPCOMING_TWO_BUSES_MSG": "The {bus}, stop {busstop}, direction {destination}: next departure in {dep1} minutes, the next one in {dep2} minutes.",
		"TOO_MANY_BUSES":         "Many buses are available at stop {busstop}. Please specify your bus in the question.",
		"NO_BUS_AVAILABLE":       "No {bus} bus is currently scheduled at stop {busstop}",
		"FAVORITE_SAVED":         "Bus stop {busstop} has been saved as favorite",
		"FAVORITE_DELETED":       "Your favorite has been deleted",
		"NO_FAVORITE":            "No bus stop is saved as favorite",
		"FAVORITE_UNAVAILABLE":   "Favorite service is temporarily unavailable",
		"HELP_MSG":               "I can tell you when the next bus is coming. Just ask: when is the next bus coming?",
		"GOODBYE_MSG":            "Goodbye!",
		"ERROR_MSG":              "Sorry, I didn't understand. Can you rephrase?",
	},
}

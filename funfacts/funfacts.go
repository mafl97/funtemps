package funfacts

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

var funFacts = FunFacts{
	Sun: []string{
		"Temperaturen i solens kjerne er 15 millioner grader Celsius.",
		"Temperaturen på det ytre laget av solen er 5778 grader Kelvin.",
	},
	Luna: []string{
		"Temperaturen på månens overflate er 106 grader Celsius om dagen.",
		"Temperaturen på månens overflate er -183 grader Celsius om natten.",
	},
	Terra: []string{
		"Den høyeste temperaturen målt på jordens overflate er 56.7 grader Celsius.",
		"Den laveste temperaturen målt på jordens overflate er -89.4 grader Celsius.",
		"Temperaturen i jordens indre kjerne er 9392 grader Kelvin.",
	},
}

func GetFunFacts(about string) []string {
	switch about {
	case "sun":
		return funFacts.Sun
	case "luna":
		return funFacts.Luna
	case "terra":
		return funFacts.Terra
	default:
		return []string{}
	}
}

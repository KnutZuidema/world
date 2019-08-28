package clothing

import "math/rand"

func getBottomTemplates() []ItemTemplate {
	return []ItemTemplate{
		{
			Name:         "breeches",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"short",
				"long",
				"loose",
				"baggy",
			},
			SuffixModifiers: []string{
				"tied at the ankles",
				"bunched at the hips",
				"with flared ends",
				"tied at the waist",
			},
		},
		{
			Name:         "pants",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"ankle-length",
				"baggy",
				"flared",
				"long",
				"narrow",
				"short",
				"straight",
				"tight-fitting",
			},
			SuffixModifiers: []string{
				"tied at the waist",
				"tied at the ankles",
				"gathered at the hips",
			},
		},
		{
			Name:         "pantaloons",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"baggy",
				"wide",
			},
			SuffixModifiers: []string{
				"bunched at the hips",
				"flared at the ends",
			},
		},
		{
			Name:         "skirt",
			Type:         "bottom",
			MaterialType: "fabric",
			PrefixModifiers: []string{
				"knee-length",
				"ankle-length",
				"short",
				"pleated",
			},
			SuffixModifiers: []string{
				"down to the knees",
				"flared at the hem",
				"tied at the waist",
			},
		},
		{
			Name:            "kilt",
			Type:            "bottom",
			MaterialType:    "fabric",
			PrefixModifiers: []string{},
			SuffixModifiers: []string{},
		},
	}
}

func getRandomBottom() Item {
	potentialTemplates := getBottomTemplates()

	template := potentialTemplates[rand.Intn(len(potentialTemplates))]

	return getItemFromTemplate(template)
}
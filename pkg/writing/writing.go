/*
Package writing implements generation of fantasy writing systems
*/
package writing

import (
	"fmt"
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/slices"
	"github.com/ironarachne/world/pkg/words"
)

// System is a system of writing
type System struct {
	Name           string `json:"name"`
	Classification string `json:"classification"`
	StrokeType     string `json:"stroke_type"`
}

func randomStrokeType() string {
	var newStroke string
	strokes := []string{}

	strokeTypes := []string{
		"angular lines",
		"arcs",
		"boxes",
		"circles",
		"dots",
		"flowing lines",
		"half-circles",
		"half-loops",
		"loops",
		"sharp angles",
		"slashes",
		"swirls",
		"triangles",
	}

	for i := 0; i < 3; i++ {
		newStroke = strokeTypes[rand.Intn(len(strokeTypes))]
		if !slices.StringIn(newStroke, strokes) {
			strokes = append(strokes, newStroke)
		}
	}

	return words.CombinePhrases(strokes)
}

// Generate randomly generates a writing system
func Generate() (System, error) {
	var writingSystem System

	classifications := []string{
		"abjad",
		"abugida",
		"alphabet",
		"ideographic",
		"pictographic",
		"semanto-phonetic",
		"syllabary",
	}

	classification, err := random.String(classifications)
	if err != nil {
		err = fmt.Errorf("Could not generate writing system: %w", err)
		return System{}, err
	}
	writingSystem.Classification = classification
	writingSystem.StrokeType = randomStrokeType()

	return writingSystem, nil
}

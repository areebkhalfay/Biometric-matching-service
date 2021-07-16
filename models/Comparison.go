package models

// Comparison Description: A similarity score for a single (1:1) image comparison operation.
type Comparison struct {
	Score float32
	NormalizedScore float32
}
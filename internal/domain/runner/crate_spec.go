package runner

import (
	"time"

	"github.com/google/uuid"
)

// CrateSize represents the size category of a pet crate.
type CrateSize string

const (
	CrateSmall  CrateSize = "small"
	CrateMedium CrateSize = "medium"
	CrateLarge  CrateSize = "large"
	CrateXLarge CrateSize = "xlarge"
)

// PetType represents the type of pet that can be transported.
type PetType string

const (
	PetCat     PetType = "cat"
	PetDog     PetType = "dog"
	PetBird    PetType = "bird"
	PetRabbit  PetType = "rabbit"
	PetReptile PetType = "reptile"
	PetOther   PetType = "other"
)

// CrateSpec is a value object representing the specification of a pet transport crate.
type CrateSpec struct {
	ID                    uuid.UUID
	RunnerID              uuid.UUID
	Size                  CrateSize
	PetTypes              []PetType
	MaxWeightKg           float64
	WidthCm               float64
	HeightCm              float64
	DepthCm               float64
	Ventilated            bool
	TemperatureControlled bool
	CreatedAt             time.Time
}

// NewCrateSpec creates a new CrateSpec value object.
func NewCrateSpec(runnerID uuid.UUID, size CrateSize, petTypes []PetType, maxWeight, w, h, d float64, ventilated, tempCtrl bool) *CrateSpec {
	return &CrateSpec{
		ID:                    uuid.New(),
		RunnerID:              runnerID,
		Size:                  size,
		PetTypes:              petTypes,
		MaxWeightKg:           maxWeight,
		WidthCm:               w,
		HeightCm:              h,
		DepthCm:               d,
		Ventilated:            ventilated,
		TemperatureControlled: tempCtrl,
		CreatedAt:             time.Now().UTC(),
	}
}

// SupportsPetType checks if this crate supports the given pet type.
func (cs *CrateSpec) SupportsPetType(petType PetType) bool {
	for _, pt := range cs.PetTypes {
		if pt == petType {
			return true
		}
	}
	return false
}

// SupportsWeight checks if this crate can hold a pet of the given weight.
func (cs *CrateSpec) SupportsWeight(weightKg float64) bool {
	return weightKg <= cs.MaxWeightKg
}

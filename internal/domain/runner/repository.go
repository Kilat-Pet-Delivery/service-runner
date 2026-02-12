package runner

import (
	"context"

	"github.com/google/uuid"
)

// RunnerRepository defines the persistence interface for Runner aggregates.
type RunnerRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Runner, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (*Runner, error)
	Save(ctx context.Context, runner *Runner) error
	Update(ctx context.Context, runner *Runner) error
	UpdateLocation(ctx context.Context, runnerID uuid.UUID, lat, lng float64) error
	FindNearbyActive(ctx context.Context, lat, lng float64, radiusKm float64, limit int) ([]*Runner, error)
}

// CrateSpecRepository defines the persistence interface for CrateSpec entities.
type CrateSpecRepository interface {
	FindByRunnerID(ctx context.Context, runnerID uuid.UUID) ([]*CrateSpec, error)
	Save(ctx context.Context, spec *CrateSpec) error
	Delete(ctx context.Context, id uuid.UUID) error
}

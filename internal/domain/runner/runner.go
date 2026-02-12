package runner

import (
	"time"

	"github.com/google/uuid"
)

// VehicleType represents the type of vehicle a runner uses.
type VehicleType string

const (
	VehicleCar        VehicleType = "car"
	VehicleVan        VehicleType = "van"
	VehicleMotorcycle VehicleType = "motorcycle"
)

// SessionStatus represents the online/offline status of a runner.
type SessionStatus string

const (
	SessionActive   SessionStatus = "active"
	SessionInactive SessionStatus = "inactive"
)

// Runner is the aggregate root for the runner domain.
type Runner struct {
	id             uuid.UUID
	userID         uuid.UUID
	fullName       string
	phone          string
	vehicleType    VehicleType
	vehiclePlate   string
	vehicleModel   string
	vehicleYear    int
	airConditioned bool
	sessionStatus  SessionStatus
	currentLat     *float64
	currentLng     *float64
	rating         float64
	totalTrips     int
	completionRate float64
	version        int64
	createdAt      time.Time
	updatedAt      time.Time
}

// NewRunner creates a new Runner aggregate.
func NewRunner(userID uuid.UUID, fullName, phone string, vehicleType VehicleType, plate, model string, year int, airCon bool) *Runner {
	now := time.Now().UTC()
	return &Runner{
		id:             uuid.New(),
		userID:         userID,
		fullName:       fullName,
		phone:          phone,
		vehicleType:    vehicleType,
		vehiclePlate:   plate,
		vehicleModel:   model,
		vehicleYear:    year,
		airConditioned: airCon,
		sessionStatus:  SessionInactive,
		rating:         0.0,
		totalTrips:     0,
		completionRate: 1.0,
		version:        1,
		createdAt:      now,
		updatedAt:      now,
	}
}

// --- Getters ---

// ID returns the runner's unique identifier.
func (r *Runner) ID() uuid.UUID { return r.id }

// UserID returns the runner's associated user identifier.
func (r *Runner) UserID() uuid.UUID { return r.userID }

// FullName returns the runner's full name.
func (r *Runner) FullName() string { return r.fullName }

// Phone returns the runner's phone number.
func (r *Runner) Phone() string { return r.phone }

// VehicleType returns the runner's vehicle type.
func (r *Runner) VehicleType() VehicleType { return r.vehicleType }

// VehiclePlate returns the runner's vehicle plate number.
func (r *Runner) VehiclePlate() string { return r.vehiclePlate }

// VehicleModel returns the runner's vehicle model.
func (r *Runner) VehicleModel() string { return r.vehicleModel }

// VehicleYear returns the runner's vehicle year.
func (r *Runner) VehicleYear() int { return r.vehicleYear }

// AirConditioned returns whether the runner's vehicle has air conditioning.
func (r *Runner) AirConditioned() bool { return r.airConditioned }

// SessionStatus returns the runner's current session status.
func (r *Runner) SessionStatus() SessionStatus { return r.sessionStatus }

// CurrentLat returns the runner's current latitude (nil if offline).
func (r *Runner) CurrentLat() *float64 { return r.currentLat }

// CurrentLng returns the runner's current longitude (nil if offline).
func (r *Runner) CurrentLng() *float64 { return r.currentLng }

// Rating returns the runner's average rating.
func (r *Runner) Rating() float64 { return r.rating }

// TotalTrips returns the runner's total completed trips.
func (r *Runner) TotalTrips() int { return r.totalTrips }

// CompletionRate returns the runner's trip completion rate.
func (r *Runner) CompletionRate() float64 { return r.completionRate }

// Version returns the runner's version for optimistic locking.
func (r *Runner) Version() int64 { return r.version }

// CreatedAt returns when the runner was created.
func (r *Runner) CreatedAt() time.Time { return r.createdAt }

// UpdatedAt returns when the runner was last updated.
func (r *Runner) UpdatedAt() time.Time { return r.updatedAt }

// --- Behavior ---

// GoOnline sets the runner's session to active and records the current location.
func (r *Runner) GoOnline(lat, lng float64) error {
	r.sessionStatus = SessionActive
	r.currentLat = &lat
	r.currentLng = &lng
	r.updatedAt = time.Now().UTC()
	return nil
}

// GoOffline sets the runner's session to inactive and clears the location.
func (r *Runner) GoOffline() {
	r.sessionStatus = SessionInactive
	r.currentLat = nil
	r.currentLng = nil
	r.updatedAt = time.Now().UTC()
}

// UpdateLocation updates the runner's current GPS coordinates.
func (r *Runner) UpdateLocation(lat, lng float64) {
	r.currentLat = &lat
	r.currentLng = &lng
	r.updatedAt = time.Now().UTC()
}

// UpdateRating updates the runner's rating and total trip count.
func (r *Runner) UpdateRating(newRating float64, newTotalTrips int) {
	r.rating = newRating
	r.totalTrips = newTotalTrips
	r.updatedAt = time.Now().UTC()
}

// IncrementVersion bumps the version for optimistic locking.
func (r *Runner) IncrementVersion() {
	r.version++
	r.updatedAt = time.Now().UTC()
}

// IsActive returns true if the runner is currently online and active.
func (r *Runner) IsActive() bool {
	return r.sessionStatus == SessionActive
}

// --- Reconstruction from persistence ---

// Reconstruct creates a Runner from persisted data (used by repositories).
func Reconstruct(
	id, userID uuid.UUID,
	fullName, phone string,
	vehicleType VehicleType,
	vehiclePlate, vehicleModel string,
	vehicleYear int,
	airConditioned bool,
	sessionStatus SessionStatus,
	currentLat, currentLng *float64,
	rating float64,
	totalTrips int,
	completionRate float64,
	version int64,
	createdAt, updatedAt time.Time,
) *Runner {
	return &Runner{
		id:             id,
		userID:         userID,
		fullName:       fullName,
		phone:          phone,
		vehicleType:    vehicleType,
		vehiclePlate:   vehiclePlate,
		vehicleModel:   vehicleModel,
		vehicleYear:    vehicleYear,
		airConditioned: airConditioned,
		sessionStatus:  sessionStatus,
		currentLat:     currentLat,
		currentLng:     currentLng,
		rating:         rating,
		totalTrips:     totalTrips,
		completionRate: completionRate,
		version:        version,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

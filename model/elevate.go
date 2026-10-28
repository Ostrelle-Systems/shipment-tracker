package model

import "time"

// ElevateRequest parameters for client elevation.
//
// swagger:model ElevateRequest
type ElevateRequest struct {
	// How long the elevation should last, in seconds. Must be between 1 second
	// and 24 hours (86400 seconds).
	//
	// required: true
	// example: 900
	// min: 1
	// max: 86400
	DurationSeconds int `form:"durationSeconds" query:"durationSeconds" json:"durationSeconds" binding:"required,min=1,max=86400"`
}

var DefaultElevationDuration = time.Hour

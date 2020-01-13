// Package resource defines an interface for recording changes to InfluxDB resources.
//
// A resource is an entity in our system, e.g. an organization, task or bucket.
// A change includes the creation, update or deletion of a resource.
package resource

// Logger records changes to resources.
type Logger interface {
	// Log a change to a resource.
	Log(Change) error
}

// Change to a resource.
type Change struct {
	Type           ChangeType
	ResourceID     string
	ResourceType   string
	OrganizationID string
	ResourceBody   []byte
}

// Type of  change.
type ChangeType string

const (
	// Create a resource.
	Create ChangeType = "create"
	// Update a resource.
	Update = "update"
	// Delete a resource
	Delete = "delete"
)

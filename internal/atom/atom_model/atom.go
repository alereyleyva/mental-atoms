package atom_model

import (
	"github.com/google/uuid"
	"time"
)

type Atom struct {
	ID        uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

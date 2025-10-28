package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID           uuid.UUID `json:"id"`
	ParentID     uuid.UUID `json:"parentID"`
	OwnerID      uuid.UUID `json:"ownerID"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"createdAt"`
	LastModified time.Time `json:"lastModified"`
}

type File struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	MimeType     string    `json:"mimeType"`
	FolderID     uuid.UUID `json:"FolderID"`
	OwnerID      uuid.UUID `json:"ownerID"`
	CreatedAt    time.Time `json:"createdAt"`
	LastModified time.Time `json:"lastModified"`
	LastAccessed time.Time `json:"lastAccessed"`
}

type FileStore interface {
	AddFile(ctx context.Context, file *File) error
}

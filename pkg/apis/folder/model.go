package folder

import "time"

type Folder struct {
	// Deprecated: use UID instead
	ID        int64     `json:"id,omitempty"`
	UID       string    `json:"uid"` // create required
	OrgID     int64     `json:"orgId,omitempty"`
	Title     string    `json:"title"` // create required
	URL       string    `json:"url,omitempty"`
	HasACL    bool      `json:"hasAcl,omitempty"`
	CanSave   bool      `json:"canSave,omitempty"`
	CanEdit   bool      `json:"canEdit,omitempty"`
	CanAdmin  bool      `json:"canAdmin,omitempty"`
	CanDelete bool      `json:"canDelete,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty"`
	Updated   time.Time `json:"updated,omitempty"`
	Version   int       `json:"version,omitempty"`
}

type Folders []Folder

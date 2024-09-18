package folder

import "time"

type Folder struct {
	// Deprecated: use UID instead
	ID        int64     `json:"id"`
	UID       string    `json:"uid"`
	OrgID     int64     `json:"orgId"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasACL    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CanDelete bool      `json:"canDelete"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version,omitempty"`
}

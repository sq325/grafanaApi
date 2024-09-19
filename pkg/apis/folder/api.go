package folder

import (
	"net/url"

	"github.com/sq325/grafanaApi/pkg/common"
)

const (
	// GET
	// POST
	folders = "/api/folders"

	// GET
	// PUT
	// DELETE
	folder = "/api/folders/{uid}"
)

type API interface {
	GetAll() Folders
	Get(uid string) *Folder
	Create(folder *Folder) error
}

type api struct {
	u     *url.URL
	token string
}

var _ API = (*api)(nil)

func NewApi(ep, token string) API {
	u, err := common.Url(ep)
	if err != nil {
	}

	return &api{
		u:     u,
		token: token,
	}
}

// get all folders in this org token
func (a *api) GetAll() Folders {
	
	return nil
}

func (a *api) Get(uid string) *Folder {
	return nil
}

func (a *api) Create(folder *Folder) error {
	return nil
}

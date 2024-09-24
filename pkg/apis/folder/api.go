package folder

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
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
	Get(uids ...string) Folders
	Create(folder Folder) error
}

type api struct {
	u      *url.URL
	token  string
	client *http.Client
}

var _ API = (*api)(nil)

func NewApi(ep, token string) API {
	u, err := common.Url(ep)
	if err != nil {
		slog.Error("url parse failed", "err", err, "ep", ep)
		return nil
	}

	return &api{
		u:      u,
		token:  token,
		client: httpclient.New(),
	}
}

// get all folders in this org token
func (a *api) GetAll() Folders {
	req, err := common.Request(http.MethodGet, a.u, folders, a.token, "", "", nil, nil)
	if err != nil {
		slog.Error("new request failed", "err", err)
		return nil
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		var msg json.RawMessage
		if decodeErr := json.NewDecoder(resp.Body).Decode(&msg); decodeErr != nil {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "msg", string(msg), "url", req.URL)
		}
	}
	defer resp.Body.Close()

	var folders Folders
	if err := json.NewDecoder(resp.Body).Decode(&folders); err != nil {
		slog.Error("decode response body to folders failed", "err", err, "url", req.URL)
		return nil
	}

	return folders
}

func (a *api) Get(uids ...string) Folders {
	if len(uids) == 0 {
		slog.Info("no folder uid")
		return nil
	}

	folders := make(Folders, 0, len(uids))
	for _, uid := range uids {
		req, err := common.Request(http.MethodGet, a.u, strings.Replace(folder, "{uid}", uid, 1), a.token, "", "", nil, nil)
		if err != nil {
			slog.Error("new request failed", "err", err)
			return nil
		}

		resp, err := a.client.Do(req)
		if err != nil || resp.StatusCode >= 400 {
			var msg json.RawMessage
			if decodeErr := json.NewDecoder(resp.Body).Decode(&msg); decodeErr != nil {
				slog.Error("client.Do Failed", "err", err, "url", req.URL)
			} else {
				slog.Error("client.Do Failed", "err", err, "msg", string(msg), "url", req.URL)
			}
		}
		defer resp.Body.Close()

		var folder Folder
		if err := json.NewDecoder(resp.Body).Decode(&folder); err != nil {
			slog.Error("decode response body to folder failed", "err", err, "url", req.URL)
			return nil
		}
		folders = append(folders, folder)
	}

	return folders
}

func (a *api) Create(folder Folder) error {
	body, err := json.Marshal(folder)
	if err != nil {
		slog.Error("marshal folder failed", "err", err)
		return err
	}

	req, err := common.Request(http.MethodPost, a.u, folders, a.token, "", "", body, nil)
	if err != nil {
		slog.Error("new request failed", "err", err)
		return err
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		var msg json.RawMessage
		folderBys, _ := json.Marshal(folder)
		if decodeErr := json.NewDecoder(resp.Body).Decode(&msg); decodeErr != nil {
			slog.Error("client.Do Failed", "err", err, "url", req.URL, "respCode", resp.StatusCode, "folder", string(folderBys))
		} else {
			slog.Error("client.Do Failed", "err", err, "msg", string(msg), "url", req.URL, "respCode", resp.StatusCode, "folder", string(folderBys))
		}
		if err == nil {
			return errors.New("respCode is larger than 400, code=" + strconv.Itoa(resp.StatusCode))
		}
		return err
	}
	defer resp.Body.Close()
	return nil
}

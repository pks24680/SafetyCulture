package folders

import (
	"github.com/gofrs/uuid"
)

type FetchFolderRequestWithPagination struct {
	OrgID  uuid.UUID
	Limit  int
	Offset int
}

func GetAllFoldersWithPagination(req *FetchFolderRequestWithPagination) (*FetchFolderResponse, error) {
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	start := req.Offset
	end := req.Offset + req.Limit

	if start > len(r) {
		start = len(r)
	}
	if end > len(r) {
		end = len(r)
	}

	var fp []*Folder
	for _, v := range r[start:end] {
		fp = append(fp, v)
	}

	return &FetchFolderResponse{Folders: fp}, nil
}

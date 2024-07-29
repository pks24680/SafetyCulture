package folders

import (
	"github.com/gofrs/uuid"
)

// Got rid of unused variables as they cause errors
// GetAllFolders fetches all folders for a given organization ID and returns them in a FetchFolderResponse.
// It first calls FetchAllFoldersByOrgID to get the list of folders and then prepares the response.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Fetch all folders by organization ID
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err // Return an error if fetching folders fails
	}

	// Create a slice of pointers to Folder with the same length as the fetched folders
	fp := make([]*Folder, len(r))
	for i, v := range r {
		fp[i] = v // Copy each folder pointer to the new slice
	}

	// Prepare the response with the fetched folders
	ffr := &FetchFolderResponse{Folders: fp}
	return ffr, nil // Return the response
}

// FetchAllFoldersByOrgID fetches all folders from sample data and filters them by the given organization ID.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // Get sample data which returns a slice of pointers to Folder - this eventually gets its data from sample.json through the method in static.go

	// Create a slice to hold the filtered folders
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder) // Add the folder to the result if the organization ID matches
		}
	}
	return resFolder, nil // Return the filtered folders
}

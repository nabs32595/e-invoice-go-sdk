package document_type

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/faridyusof727/e-invoice-go-sdk/constants"
)

// AllDocumentTypes implements Provider.
func (c *Client) AllDocumentTypes(ctx context.Context) ([]DocumentType, error) {
	dataResponse := &struct {
		Result []DocumentType `json:"result"`
	}{}

	err := requests.
		URL(c.authClient.Config().Url).
		Method("GET").
		Path("/api/v1.0/documenttypes").
		Header("Accept", constants.HttpHeaderContentTypeJson).
		Header("Authorization", fmt.Sprintf("Bearer %s", c.authClient.AccessToken())).
		Header("Accept-Language", "en").
		Header("Content-Type", constants.HttpHeaderContentTypeJson).
		ToJSON(dataResponse).
		Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all document types: %w", err)
	}

	return dataResponse.Result, nil
}

// DocumentType implements Provider.
func (c *Client) DocumentType(ctx context.Context, id int64) (*DocumentType, error) {
	dt := &DocumentType{}

	err := requests.
		URL(c.authClient.Config().Url).
		Method("GET").
		Pathf("/api/v1.0/documenttypes/%d", id).
		Header("Accept", constants.HttpHeaderContentTypeJson).
		Header("Authorization", fmt.Sprintf("Bearer %s", c.authClient.AccessToken())).
		Header("Accept-Language", "en").
		Header("Content-Type", constants.HttpHeaderContentTypeJson).
		ToJSON(dt).
		Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get document type: %w", err)
	}

	return dt, nil
}

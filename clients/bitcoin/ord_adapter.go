package bitcoin

import (
	"encoding/json"
	"fmt"
	"github.com/aura-nw/lotus-core/types"
	"io"
	"net/http"
)

type IOrdAdapter interface {
	GetInscriptionIdsByBlock(blockHeight int64) ([]string, error)
	GetInscription(inscriptionId string) (*types.GetInscriptionResponse, error)
	GetOutput(output string) (*types.GetOutputResponse, error)
	GetContent(inscriptionId string) (*types.GetContentResponse, error)
}

type OrdAdapterImpl struct {
	host string
}

func NewOrdAdapter(ordHost string) (IOrdAdapter, error) {
	return &OrdAdapterImpl{
		host: ordHost,
	}, nil
}

func (a *OrdAdapterImpl) GetInscriptionIdsByBlock(blockHeight int64) ([]string, error) {
	var result *types.GetInscriptionIdsResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/inscriptions/block/%d", a.host, blockHeight), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	result, err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result.Ids, nil
}

func (a *OrdAdapterImpl) GetInscription(inscriptionId string) (*types.GetInscriptionResponse, error) {
	var result *types.GetInscriptionResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/inscription/%s", a.host, inscriptionId), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return parseResponse(resp, result)
}

func (a *OrdAdapterImpl) GetOutput(output string) (*types.GetOutputResponse, error) {
	var result *types.GetOutputResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/output/%s", a.host, output), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return parseResponse(resp, result)
}

func (a *OrdAdapterImpl) GetContent(inscriptionId string) (*types.GetContentResponse, error) {
	var result *types.GetContentResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/content/%s", a.host, inscriptionId), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return parseResponse(resp, result)
}

func (a *OrdAdapterImpl) sendRequest(method, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Accept", "application/json")
	return client.Do(req)
}

type Response interface {
	*types.GetInscriptionIdsResponse | *types.GetContentResponse |
		*types.GetOutputResponse | *types.GetInscriptionResponse
}

func parseResponse[T Response](resp *http.Response, result T) (T, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return nil, err
	}

	return result, nil
}

package adaptors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IOrdAdapter interface {
	GetInscriptionIdsByBlock(blockHeight int64) ([]string, error)
	GetInscription(inscriptionId string) (*GetInscriptionResponse, error)
	GetOutput(output string) (*GetOutputResponse, error)
	GetContent(inscriptionId string) (*GetContentResponse, error)
	GetMetadata(inscriptionId string) (interface{}, error)
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
	var result GetInscriptionIdsResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/inscriptions/block/%d", a.host, blockHeight), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return nil, err
	}

	return result.Ids, nil
}

func (a *OrdAdapterImpl) GetInscription(inscriptionId string) (*GetInscriptionResponse, error) {
	var result *GetInscriptionResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/inscription/%s", a.host, inscriptionId), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
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

func (a *OrdAdapterImpl) GetOutput(output string) (*GetOutputResponse, error) {
	var result *GetOutputResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/output/%s", a.host, output), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
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

func (a *OrdAdapterImpl) GetContent(inscriptionId string) (*GetContentResponse, error) {
	var result *GetContentResponse
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/content/%s", a.host, inscriptionId), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
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

func (a *OrdAdapterImpl) GetMetadata(inscriptionId string) (interface{}, error) {
	var result interface{}
	resp, err := a.sendRequest("GET", fmt.Sprintf("%s/content/%s", a.host, inscriptionId), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
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

func (a *OrdAdapterImpl) sendRequest(method, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Accept", "application/json")
	return client.Do(req)
}

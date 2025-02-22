// generated by 'threeport-sdk gen' - do not edit

package v0

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	tpclient_lib "github.com/threeport/threeport/pkg/client/lib/v0"
	tputil "github.com/threeport/threeport/pkg/util/v0"
	"net/http"
	v0 "wordpress-threeport-module/pkg/api/v0"
)

// GetWordpressDefinitions fetches all wordpress definitions.
// TODO: implement pagination
func GetWordpressDefinitions(apiClient *http.Client, apiAddr string) (*[]v0.WordpressDefinition, error) {
	var wordpressDefinitions []v0.WordpressDefinition

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s", apiAddr, v0.PathWordpressDefinitions),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressDefinitions, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &wordpressDefinitions, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressDefinitions); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressDefinitions, nil
}

// GetWordpressDefinitionByID fetches a wordpress definition by ID.
func GetWordpressDefinitionByID(apiClient *http.Client, apiAddr string, id uint) (*v0.WordpressDefinition, error) {
	var wordpressDefinition v0.WordpressDefinition

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s/%d", apiAddr, v0.PathWordpressDefinitions, id),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &wordpressDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressDefinition, nil
}

// GetWordpressDefinitionsByQueryString fetches wordpress definitions by provided query string.
func GetWordpressDefinitionsByQueryString(apiClient *http.Client, apiAddr string, queryString string) (*[]v0.WordpressDefinition, error) {
	var wordpressDefinitions []v0.WordpressDefinition

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s?%s", apiAddr, v0.PathWordpressDefinitions, queryString),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressDefinitions, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &wordpressDefinitions, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressDefinitions); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressDefinitions, nil
}

// GetWordpressDefinitionByName fetches a wordpress definition by name.
func GetWordpressDefinitionByName(apiClient *http.Client, apiAddr, name string) (*v0.WordpressDefinition, error) {
	var wordpressDefinitions []v0.WordpressDefinition

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s?name=%s", apiAddr, v0.PathWordpressDefinitions, name),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &v0.WordpressDefinition{}, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.WordpressDefinition{}, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressDefinitions); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	switch {
	case len(wordpressDefinitions) < 1:
		return &v0.WordpressDefinition{}, errors.New(fmt.Sprintf("no wordpress definition with name %s", name))
	case len(wordpressDefinitions) > 1:
		return &v0.WordpressDefinition{}, errors.New(fmt.Sprintf("more than one wordpress definition with name %s returned", name))
	}

	return &wordpressDefinitions[0], nil
}

// CreateWordpressDefinition creates a new wordpress definition.
func CreateWordpressDefinition(apiClient *http.Client, apiAddr string, wordpressDefinition *v0.WordpressDefinition) (*v0.WordpressDefinition, error) {
	tpclient_lib.ReplaceAssociatedObjectsWithNil(wordpressDefinition)
	jsonWordpressDefinition, err := tputil.MarshalObject(wordpressDefinition)
	if err != nil {
		return wordpressDefinition, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s", apiAddr, v0.PathWordpressDefinitions),
		http.MethodPost,
		bytes.NewBuffer(jsonWordpressDefinition),
		map[string]string{},
		http.StatusCreated,
	)
	if err != nil {
		return wordpressDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return wordpressDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return wordpressDefinition, nil
}

// UpdateWordpressDefinition updates a wordpress definition.
func UpdateWordpressDefinition(apiClient *http.Client, apiAddr string, wordpressDefinition *v0.WordpressDefinition) (*v0.WordpressDefinition, error) {
	tpclient_lib.ReplaceAssociatedObjectsWithNil(wordpressDefinition)
	// capture the object ID, make a copy of the object, then remove fields that
	// cannot be updated in the API
	wordpressDefinitionID := *wordpressDefinition.ID
	payloadWordpressDefinition := *wordpressDefinition
	payloadWordpressDefinition.ID = nil
	payloadWordpressDefinition.CreatedAt = nil
	payloadWordpressDefinition.UpdatedAt = nil

	jsonWordpressDefinition, err := tputil.MarshalObject(payloadWordpressDefinition)
	if err != nil {
		return wordpressDefinition, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s/%d", apiAddr, v0.PathWordpressDefinitions, wordpressDefinitionID),
		http.MethodPatch,
		bytes.NewBuffer(jsonWordpressDefinition),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return wordpressDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return wordpressDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&payloadWordpressDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	payloadWordpressDefinition.ID = &wordpressDefinitionID
	return &payloadWordpressDefinition, nil
}

// DeleteWordpressDefinition deletes a wordpress definition by ID.
func DeleteWordpressDefinition(apiClient *http.Client, apiAddr string, id uint) (*v0.WordpressDefinition, error) {
	var wordpressDefinition v0.WordpressDefinition

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s/%d", apiAddr, v0.PathWordpressDefinitions, id),
		http.MethodDelete,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressDefinition, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &wordpressDefinition, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressDefinition); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressDefinition, nil
}

// GetWordpressInstances fetches all wordpress instances.
// TODO: implement pagination
func GetWordpressInstances(apiClient *http.Client, apiAddr string) (*[]v0.WordpressInstance, error) {
	var wordpressInstances []v0.WordpressInstance

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s", apiAddr, v0.PathWordpressInstances),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressInstances, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &wordpressInstances, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressInstances); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressInstances, nil
}

// GetWordpressInstanceByID fetches a wordpress instance by ID.
func GetWordpressInstanceByID(apiClient *http.Client, apiAddr string, id uint) (*v0.WordpressInstance, error) {
	var wordpressInstance v0.WordpressInstance

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s/%d", apiAddr, v0.PathWordpressInstances, id),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &wordpressInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressInstance, nil
}

// GetWordpressInstancesByQueryString fetches wordpress instances by provided query string.
func GetWordpressInstancesByQueryString(apiClient *http.Client, apiAddr string, queryString string) (*[]v0.WordpressInstance, error) {
	var wordpressInstances []v0.WordpressInstance

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s?%s", apiAddr, v0.PathWordpressInstances, queryString),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressInstances, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &wordpressInstances, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressInstances); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressInstances, nil
}

// GetWordpressInstanceByName fetches a wordpress instance by name.
func GetWordpressInstanceByName(apiClient *http.Client, apiAddr, name string) (*v0.WordpressInstance, error) {
	var wordpressInstances []v0.WordpressInstance

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s?name=%s", apiAddr, v0.PathWordpressInstances, name),
		http.MethodGet,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &v0.WordpressInstance{}, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data)
	if err != nil {
		return &v0.WordpressInstance{}, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressInstances); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	switch {
	case len(wordpressInstances) < 1:
		return &v0.WordpressInstance{}, errors.New(fmt.Sprintf("no wordpress instance with name %s", name))
	case len(wordpressInstances) > 1:
		return &v0.WordpressInstance{}, errors.New(fmt.Sprintf("more than one wordpress instance with name %s returned", name))
	}

	return &wordpressInstances[0], nil
}

// CreateWordpressInstance creates a new wordpress instance.
func CreateWordpressInstance(apiClient *http.Client, apiAddr string, wordpressInstance *v0.WordpressInstance) (*v0.WordpressInstance, error) {
	tpclient_lib.ReplaceAssociatedObjectsWithNil(wordpressInstance)
	jsonWordpressInstance, err := tputil.MarshalObject(wordpressInstance)
	if err != nil {
		return wordpressInstance, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s", apiAddr, v0.PathWordpressInstances),
		http.MethodPost,
		bytes.NewBuffer(jsonWordpressInstance),
		map[string]string{},
		http.StatusCreated,
	)
	if err != nil {
		return wordpressInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return wordpressInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return wordpressInstance, nil
}

// UpdateWordpressInstance updates a wordpress instance.
func UpdateWordpressInstance(apiClient *http.Client, apiAddr string, wordpressInstance *v0.WordpressInstance) (*v0.WordpressInstance, error) {
	tpclient_lib.ReplaceAssociatedObjectsWithNil(wordpressInstance)
	// capture the object ID, make a copy of the object, then remove fields that
	// cannot be updated in the API
	wordpressInstanceID := *wordpressInstance.ID
	payloadWordpressInstance := *wordpressInstance
	payloadWordpressInstance.ID = nil
	payloadWordpressInstance.CreatedAt = nil
	payloadWordpressInstance.UpdatedAt = nil

	jsonWordpressInstance, err := tputil.MarshalObject(payloadWordpressInstance)
	if err != nil {
		return wordpressInstance, fmt.Errorf("failed to marshal provided object to JSON: %w", err)
	}

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s/%d", apiAddr, v0.PathWordpressInstances, wordpressInstanceID),
		http.MethodPatch,
		bytes.NewBuffer(jsonWordpressInstance),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return wordpressInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return wordpressInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&payloadWordpressInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	payloadWordpressInstance.ID = &wordpressInstanceID
	return &payloadWordpressInstance, nil
}

// DeleteWordpressInstance deletes a wordpress instance by ID.
func DeleteWordpressInstance(apiClient *http.Client, apiAddr string, id uint) (*v0.WordpressInstance, error) {
	var wordpressInstance v0.WordpressInstance

	response, err := tpclient_lib.GetResponse(
		apiClient,
		fmt.Sprintf("%s%s/%d", apiAddr, v0.PathWordpressInstances, id),
		http.MethodDelete,
		new(bytes.Buffer),
		map[string]string{},
		http.StatusOK,
	)
	if err != nil {
		return &wordpressInstance, fmt.Errorf("call to threeport API returned unexpected response: %w", err)
	}

	jsonData, err := json.Marshal(response.Data[0])
	if err != nil {
		return &wordpressInstance, fmt.Errorf("failed to marshal response data from threeport API: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	decoder.UseNumber()
	if err := decoder.Decode(&wordpressInstance); err != nil {
		return nil, fmt.Errorf("failed to decode object in response data from threeport API: %w", err)
	}

	return &wordpressInstance, nil
}

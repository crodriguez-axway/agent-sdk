package apic

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	corecfg "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/config"
	"github.com/tidwall/gjson"
)

type actionType int

const (
	addAPI    actionType = iota
	updateAPI            = iota
	deleteAPI            = iota
)

type serviceExecution int

const (
	addAPIServerSpec serviceExecution = iota + 1
	addAPIServerRevisionSpec
	addAPIServerInstanceSpec
	deleteAPIServerSpec
	addCatalog
	addCatalogImage
	updateCatalog
)

//ServiceBody -
type ServiceBody struct {
	NameToPush       string `json:",omitempty"`
	APIName          string `json:",omitempty"`
	RestAPIID        string `json:",omitempty"`
	URL              string `json:",omitempty"`
	Stage            string `json:",omitempty"`
	TeamID           string `json:",omitempty"`
	Description      string `json:",omitempty"`
	Version          string `json:",omitempty"`
	AuthPolicy       string `json:",omitempty"`
	Swagger          []byte `json:",omitempty"`
	Documentation    []byte `json:",omitempty"`
	Tags             map[string]interface{}
	Buffer           []byte            `json:",omitempty"`
	AgentMode        corecfg.AgentMode `json:",omitempty"`
	ServiceExecution serviceExecution  `json:"omitempty"`
	Image            string
	ImageContentType string
}

//CatalogPropertyValue -
type CatalogPropertyValue struct {
	URL        string `json:"url"`
	AuthPolicy string `json:"authPolicy"`
}

//CatalogProperty -
type CatalogProperty struct {
	Key   string               `json:"key"`
	Value CatalogPropertyValue `json:"value"`
}

//CatalogRevisionProperty -
type CatalogRevisionProperty struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

//CatalogItemInitRevision -
type CatalogItemInitRevision struct {
	ID         string                    `json:"id,omitempty"`
	Properties []CatalogRevisionProperty `json:"properties"`
	Number     int                       `json:"number,omitempty"`
	Version    string                    `json:"version"`
	State      string                    `json:"state"`
	Status     string                    `json:"status,omitempty"`
}

//CatalogItemRevision -
type CatalogItemRevision struct {
	ID string `json:"id,omitempty"`
	// metadata []CatalogRevisionProperty `json:"properties"`
	Number  int    `json:"number,omitempty"`
	Version string `json:"version"`
	State   string `json:"state"`
	Status  string `json:"status,omitempty"`
}

//CatalogSubscription -
type CatalogSubscription struct {
	Enabled         bool                      `json:"enabled"`
	AutoSubscribe   bool                      `json:"autoSubscribe"`
	AutoUnsubscribe bool                      `json:"autoUnsubscribe"`
	Properties      []CatalogRevisionProperty `json:"properties"`
}

//CatalogItemInit -
type CatalogItemInit struct {
	OwningTeamID       string                  `json:"owningTeamId"`
	DefinitionType     string                  `json:"definitionType"`
	DefinitionSubType  string                  `json:"definitionSubType"`
	DefinitionRevision int                     `json:"definitionRevision"`
	Name               string                  `json:"name"`
	Description        string                  `json:"description,omitempty"`
	Properties         []CatalogProperty       `json:"properties,omitempty"`
	Tags               []string                `json:"tags,omitempty"`
	Visibility         string                  `json:"visibility"` // default: RESTRICTED
	Subscription       CatalogSubscription     `json:"subscription,omitempty"`
	Revision           CatalogItemInitRevision `json:"revision,omitempty"`
	CategoryReferences string                  `json:"categoryReferences,omitempty"`
}

// CatalogItemImage -
type CatalogItemImage struct {
	DataType      string `json:"data,omitempty"`
	Base64Content string `json:"base64,omitempty"`
}

//CatalogItem -
type CatalogItem struct {
	ID                 string `json:"id"`
	OwningTeamID       string `json:"owningTeamId"`
	DefinitionType     string `json:"definitionType"`
	DefinitionSubType  string `json:"definitionSubType"`
	DefinitionRevision int    `json:"definitionRevision"`
	Name               string `json:"name"`
	// relationships
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	// metadata
	Visibility string `json:"visibility"` // default: RESTRICTED
	State      string `json:"state"`      // default: UNPUBLISHED
	Access     string `json:"access,omitempty"`
	// availableRevisions
	LatestVersion        int                 `json:"latestVersion,omitempty"`
	TotalSubscriptions   int                 `json:"totalSubscriptions,omitempty"`
	LatestVersionDetails CatalogItemRevision `json:"latestVersionDetails,omitempty"`
	Image                *CatalogItemImage   `json:"image,omitempty"`
	// categories
}

// APIServer -
type APIServer struct {
	Name       string                 `json:"name"`
	Title      string                 `json:"title"`
	Tags       []string               `json:"tags"`
	Attributes map[string]interface{} `json:"attributes"`
	Spec       interface{}            `json:"spec"`
}

// APIServiceSpec -
type APIServiceSpec struct {
	Description string `json:"description"`
}

// APIServiceRevisionSpec -
type APIServiceRevisionSpec struct {
	APIService string             `json:"apiService"`
	Definition RevisionDefinition `json:"definition"`
}

// RevisionDefinition -
type RevisionDefinition struct {
	Type  string `json:"type,omitempty"`
	Value []byte `json:"value,omitempty"`
}

// APIServerInstanceSpec -
type APIServerInstanceSpec struct {
	APIServiceRevision string     `json:"apiServiceRevision,omitempty"`
	InstanceEndPoint   []EndPoint `json:"endpoint,omitempty"`
}

// EndPoint -
type EndPoint struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

const (
	subscriptionSchema = "{\"type\": \"object\", \"$schema\": \"http://json-schema.org/draft-04/schema#\", \"description\": \"Subscription specification for API Key authentication\", \"x-axway-unique-keys\": \"APIC_APPLICATION_ID\", \"properties\": {\"applicationId\": {\"type\": \"string\", \"description\": \"Select an application\", \"x-axway-ref-apic\": \"APIC_APPLICATION_ID\"}}, \"required\":[\"applicationId\"]}"
)

func (c *ServiceClient) deployService(serviceBody ServiceBody, method, url string) (string, error) {
	buffer, err := c.marshalServiceBody(serviceBody)
	if err != nil {
		log.Error("Error creating service item: ", err)
		return "", err
	}

	return c.deployAPI(method, url, buffer)
}

// CreateService - Creates a catalog item or API service for the definition based on the agent mode
func (c *ServiceClient) CreateService(serviceBody ServiceBody) (string, error) {
	if c.cfg.GetAgentMode() == corecfg.Connected {
		return c.addAPICService(serviceBody)
	}
	return c.addCatalog(serviceBody)
}

// AddToAPICServer -
func (c *ServiceClient) addAPICService(serviceBody ServiceBody) (string, error) {
	itemID := ""
	// Verify if the api already exists
	if c.isNewAPI(serviceBody) {
		// add api
		serviceBody.ServiceExecution = addAPIServerSpec
		itemID, err := c.deployService(serviceBody, http.MethodPost, c.cfg.GetAPIServerServicesURL())
		if err != nil {
			return itemID, err
		}
	}

	// add api revision
	serviceBody.ServiceExecution = addAPIServerRevisionSpec
	itemID, err := c.deployService(serviceBody, http.MethodPost, c.cfg.GetAPIServerServicesRevisionsURL())
	if err != nil {
		log.Errorf("Error adding API revision for API %v", serviceBody.NameToPush)
		return c.rollbackAPIService(serviceBody)
	}

	// add api instance
	serviceBody.ServiceExecution = addAPIServerInstanceSpec
	itemID, err = c.deployService(serviceBody, http.MethodPost, c.cfg.GetAPIServerServicesInstancesURL())
	if err != nil {
		log.Errorf("Error adding API instance for API %v", serviceBody.NameToPush)
		return c.rollbackAPIService(serviceBody)
	}

	return itemID, err
}

// rollbackAPIService - if the process to add api/revision/instance fails, delete the api that was created
func (c *ServiceClient) rollbackAPIService(serviceBody ServiceBody) (string, error) {
	// rollback and remove the API service
	serviceBody.ServiceExecution = deleteAPIServerSpec
	return c.deployService(serviceBody, http.MethodDelete, c.cfg.DeleteAPIServerServicesURL()+"/"+sanitizeAPIName(serviceBody.APIName))
}

// AddToAPIC -
func (c *ServiceClient) addCatalog(serviceBody ServiceBody) (string, error) {
	serviceBody.ServiceExecution = addCatalog
	itemID, err := c.deployService(serviceBody, http.MethodPost, c.cfg.GetCatalogItemsURL())
	if err != nil {
		return "", err
	}
	if serviceBody.Image != "" {
		serviceBody.ServiceExecution = addCatalogImage
		_, err = c.deployService(serviceBody, http.MethodPost, c.cfg.GetCatalogItemImageURL(itemID))
		if err != nil {
			log.Warn("Unable to add image to the catalog item. " + err.Error())
		}
	}
	return itemID, nil
}

// UpdateService -
func (c *ServiceClient) UpdateService(ID string, serviceBody ServiceBody) (string, error) {
	serviceBody.ServiceExecution = updateCatalog
	return c.deployService(serviceBody, http.MethodPut, c.cfg.GetCatalogItemsURL()+ID)
}

// CreateService -
func (c *ServiceClient) marshalServiceBody(serviceBody ServiceBody) ([]byte, error) {
	if !isValidAuthPolicy(serviceBody.AuthPolicy) {
		return nil, fmt.Errorf("Unsuppored security policy '%v'. ", serviceBody.AuthPolicy)
	}
	if serviceBody.AgentMode == corecfg.Connected {
		return c.createAPIServerBody(serviceBody)
	}
	return c.createCatalogBody(serviceBody)
}

func (c *ServiceClient) createCatalogBody(serviceBody ServiceBody) ([]byte, error) {
	var spec []byte
	var err error
	switch serviceBody.ServiceExecution {
	case addCatalog:
		spec, err = c.marshalCatalogItemInit(serviceBody)
	case addCatalogImage:
		spec, err = c.marshalCatalogItemImage(serviceBody)
	case updateCatalog:
		spec, err = c.marshalCatalogItem(serviceBody)
	default:
		return nil, errors.New("Invalid catalog operation")
	}
	if err != nil {
		return nil, err
	}
	return spec, nil
}

func (c *ServiceClient) marshalCatalogItemInit(serviceBody ServiceBody) ([]byte, error) {
	newCatalogItem := CatalogItemInit{
		DefinitionType:     "API",
		DefinitionSubType:  "swaggerv2",
		DefinitionRevision: 1,
		Name:               serviceBody.NameToPush,
		OwningTeamID:       serviceBody.TeamID,
		Description:        serviceBody.Description,
		Properties: []CatalogProperty{
			{
				Key: "accessInfo",
				Value: CatalogPropertyValue{
					AuthPolicy: serviceBody.AuthPolicy,
					URL:        serviceBody.URL,
				},
			},
		},

		Tags:       c.mapToTagsArray(serviceBody.Tags),
		Visibility: "RESTRICTED", // default value
		Subscription: CatalogSubscription{
			Enabled:         true,
			AutoSubscribe:   true,
			AutoUnsubscribe: true,
			Properties: []CatalogRevisionProperty{{
				Key:   "profile",
				Value: json.RawMessage([]byte(subscriptionSchema)),
			}},
		},
		Revision: CatalogItemInitRevision{
			Version: serviceBody.Version,
			State:   "PUBLISHED",
			Properties: []CatalogRevisionProperty{
				{
					Key:   "documentation",
					Value: json.RawMessage(string(serviceBody.Documentation)),
				},
				{
					Key:   "swagger",
					Value: json.RawMessage(serviceBody.Swagger),
				},
			},
		},
	}

	return json.Marshal(newCatalogItem)
}

func isValidAuthPolicy(auth string) bool {
	for _, item := range ValidPolicies {
		if item == auth {
			return true
		}
	}
	return false
}

// createAPIServerBody - This function is being used by both the api server creation and api server revision creation
func (c *ServiceClient) createAPIServerBody(serviceBody ServiceBody) ([]byte, error) {
	attributes := make(map[string]interface{})
	attributes["externalAPIID"] = serviceBody.RestAPIID

	// spec needs to adhere to environment schema
	var spec interface{}
	name := sanitizeAPIName(serviceBody.APIName)

	switch serviceBody.ServiceExecution {
	case addAPIServerSpec:
		spec = APIServiceSpec{
			Description: serviceBody.Description,
		}
	case deleteAPIServerSpec:
		spec = APIServiceSpec{
			Description: serviceBody.Description,
		}
		return nil, nil
	case addAPIServerRevisionSpec:
		revisionDefinition := RevisionDefinition{
			Type:  "oas2",
			Value: serviceBody.Swagger,
		}
		spec = APIServiceRevisionSpec{
			APIService: name,
			Definition: revisionDefinition,
		}
		// reset the name here to include the stage
		name = sanitizeAPIName(serviceBody.APIName + serviceBody.Stage)
	case addAPIServerInstanceSpec:
		endPoints := []EndPoint{}
		name += strings.ToLower(serviceBody.Stage)
		swaggerHost := strings.Split(gjson.Get(string(serviceBody.Swagger), "host").String(), ":")
		host := swaggerHost[0]
		port := 443
		if len(swaggerHost) > 1 {
			swaggerPort, err := strconv.Atoi(swaggerHost[1])
			if err == nil {
				port = swaggerPort
			}
		}

		// Iterate through protocols and create endpoints for intances
		protocols := gjson.Get(string(serviceBody.Swagger), "schemes")
		schemes := make([]string, 0)
		json.Unmarshal([]byte(protocols.Raw), &schemes)
		for _, protocol := range schemes {
			endPoint := EndPoint{
				Host:     host,
				Port:     port,
				Protocol: protocol,
			}
			endPoints = append(endPoints, endPoint)
		}

		// reset the name here to include the stage
		name = sanitizeAPIName(serviceBody.APIName + serviceBody.Stage)
		spec = APIServerInstanceSpec{
			APIServiceRevision: name,
			InstanceEndPoint:   endPoints,
		}
	default:
		return nil, errors.New("Error getting execution service -- not set")
	}

	newtags := c.mapToTagsArray(serviceBody.Tags)

	apiServerService := APIServer{
		Name:       name,
		Title:      serviceBody.NameToPush,
		Attributes: attributes,
		Spec:       spec,
		Tags:       newtags,
	}

	return json.Marshal(apiServerService)
}

// Sanitize name to be path friendly and follow this regex: ^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*
func sanitizeAPIName(name string) string {
	// convert all letters to lower first
	newName := strings.ToLower(name)
	// fmt.Println(name)

	// parse name out. All valid parts must be '-', '.', a-z, or 0-9
	re := regexp.MustCompile(`[-\.a-z0-9]*`)
	matches := re.FindAllString(newName, -1)

	// join all of the parts, separated with '-'. This in effect is substituting all illegal chars with a '-'
	newName = strings.Join(matches, "-")

	// The regex rule says that the name must not begin or end with a '-' or '.', so trim them off
	newName = strings.TrimLeft(strings.TrimRight(newName, "-."), "-.")

	// The regex rule also says that the name must not have a sequence of ".-", "-.", or "..", so replace them
	r1 := strings.ReplaceAll(newName, "-.", "--")
	r2 := strings.ReplaceAll(r1, ".-", "--")
	r3 := strings.ReplaceAll(r2, "..", "--")

	return r3
}

// CreateCatalogItemBodyForUpdate -
func (c *ServiceClient) marshalCatalogItem(serviceBody ServiceBody) ([]byte, error) {
	newCatalogItem := CatalogItem{
		DefinitionType:    "API",
		DefinitionSubType: "swaggerv2",

		DefinitionRevision: 1,
		Name:               serviceBody.NameToPush,
		OwningTeamID:       serviceBody.TeamID,
		Description:        serviceBody.Description,
		Tags:               c.mapToTagsArray(serviceBody.Tags),
		Visibility:         "RESTRICTED",  // default value
		State:              "UNPUBLISHED", //default
		LatestVersionDetails: CatalogItemRevision{
			Version: serviceBody.Version,
			State:   "PUBLISHED",
		},
	}

	return json.Marshal(newCatalogItem)
}

// marshals the catalog image body
func (c *ServiceClient) marshalCatalogItemImage(serviceBody ServiceBody) ([]byte, error) {
	catalogImage := CatalogItemImage{
		DataType:      serviceBody.ImageContentType,
		Base64Content: serviceBody.Image,
	}
	return json.Marshal(catalogImage)
}

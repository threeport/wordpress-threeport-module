// generated by 'threeport-sdk gen' but will not be regenerated - intended for modification

package cmd

import (
	"fmt"
	"net/http"
	api_v0 "wordpress-threeport-module/pkg/api/v0"
	config_v0 "wordpress-threeport-module/pkg/config/v0"
)

// outputDescribev0WordpressDefinitionCmd produces the plain description
// output for the 'tptctl describe wordpress-definition' command
func outputDescribev0WordpressDefinitionCmd(
	wordpressDefinition *api_v0.WordpressDefinition,
	wordpressDefinitionConfig *config_v0.WordpressDefinitionConfig,
	apiClient *http.Client,
	apiEndpoint string,
) error {
	// output describe details
	fmt.Printf(
		"* WordpressDefinition Name: %s\n",
		*wordpressDefinitionConfig.WordpressDefinition.Name,
	)
	fmt.Printf(
		"* Created: %s\n",
		*wordpressDefinition.CreatedAt,
	)
	fmt.Printf(
		"* Last Modified: %s\n",
		*wordpressDefinition.UpdatedAt,
	)

	return nil
}

// outputDescribev0WordpressInstanceCmd produces the plain description
// output for the 'tptctl describe wordpress-instance' command
func outputDescribev0WordpressInstanceCmd(
	wordpressInstance *api_v0.WordpressInstance,
	wordpressInstanceConfig *config_v0.WordpressInstanceConfig,
	apiClient *http.Client,
	apiEndpoint string,
) error {
	// output describe details
	fmt.Printf(
		"* WordpressInstance Name: %s\n",
		*wordpressInstanceConfig.WordpressInstance.Name,
	)
	fmt.Printf(
		"* Created: %s\n",
		*wordpressInstance.CreatedAt,
	)
	fmt.Printf(
		"* Last Modified: %s\n",
		*wordpressInstance.UpdatedAt,
	)

	return nil
}

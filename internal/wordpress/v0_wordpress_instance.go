// generated by 'threeport-sdk gen' but will not be regenerated - intended for modification

package wordpress

import (
	"fmt"

	logr "github.com/go-logr/logr"
	tpapi "github.com/threeport/threeport/pkg/api/v0"
	tpclient "github.com/threeport/threeport/pkg/client/v0"
	controller "github.com/threeport/threeport/pkg/controller/v0"
	util "github.com/threeport/threeport/pkg/util/v0"

	v0 "wordpress-threeport-module/pkg/api/v0"
	client "wordpress-threeport-module/pkg/client/v0"
)

// v0WordpressInstanceCreated performs reconciliation when a v0 WordpressInstance
// has been created.
func v0WordpressInstanceCreated(
	r *controller.Reconciler,
	wordpressInstance *v0.WordpressInstance,
	log *logr.Logger,
) (int64, error) {
	// get attached Kubernetes runtime instance ID
	kubernetesRuntimeInstanceId, err := tpclient.GetObjectIdByAttachedObject(
		r.APIClient,
		r.APIServer,
		tpapi.ObjectTypeKubernetesRuntimeInstance,
		v0.ObjectTypeWordpressInstance,
		*wordpressInstance.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get Kubernetes runtime instance by attachment: %w", err)
	}

	// get workload definition attached to wordpress definition
	workloadDefinitionId, err := tpclient.GetObjectIdByAttachedObject(
		r.APIClient,
		r.APIServer,
		tpapi.ObjectTypeWorkloadDefinition,
		v0.ObjectTypeWordpressDefinition,
		*wordpressInstance.WordpressDefinitionID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get workload definition by attachment: %w", err)
	}

	// create workload instance if it doesn't already exist
	nameQuery := fmt.Sprintf("name=%s", *wordpressInstance.Name)
	existingWorkloadInstances, err := tpclient.GetWorkloadInstancesByQueryString(
		r.APIClient,
		r.APIServer,
		nameQuery,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to check for workload instance with name %s: %w", *wordpressInstance.Name, err)
	}
	var createdWorkloadInstance *tpapi.WorkloadInstance
	if len(*existingWorkloadInstances) == 0 {
		workloadInstance := tpapi.WorkloadInstance{
			Instance: tpapi.Instance{
				Name: wordpressInstance.Name,
			},
			KubernetesRuntimeInstanceID: kubernetesRuntimeInstanceId,
			WorkloadDefinitionID:        workloadDefinitionId,
		}
		createdWorkloadInst, err := tpclient.CreateWorkloadInstance(
			r.APIClient,
			r.APIServer,
			&workloadInstance,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to create workload instance: %w", err)
		}
		createdWorkloadInstance = createdWorkloadInst
	}

	// establish attachment between wordpress instance and workload instance
	if err := tpclient.EnsureAttachedObjectReferenceExists(
		r.APIClient,
		r.APIServer,
		tpapi.ObjectTypeWorkloadInstance,
		createdWorkloadInstance.ID,
		v0.ObjectTypeWordpressInstance,
		wordpressInstance.ID,
	); err != nil {
		return 0, fmt.Errorf("failed to attach wordpress instance to workload instance: %w", err)
	}

	// get wordpress definition
	wordpressDefinition, err := client.GetWordpressDefinitionByID(
		r.APIClient,
		r.APIServer,
		*wordpressInstance.WordpressDefinitionID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get wordpress definition: %w", err)
	}

	// get infra provider for kubernetes runtime - needed to determine storage
	// class for wordpress PVC
	infraProvider, err := tpclient.GetInfraProviderByKubernetesRuntimeInstanceID(
		r.APIClient,
		r.APIServer,
		kubernetesRuntimeInstanceId,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to determine infra provider for Kubernetes runtime: %w", err)
	}

	// get the manifest for the PVC
	pvcManifest, err := getPvcManifest(
		*infraProvider,
		*wordpressDefinition.Name,
		*wordpressDefinition.Environment,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get Kubernetes manifest for persistent volume claim: %w", err)
	}

	// create the workload resource instance for the PVC
	pvcWri := tpapi.WorkloadResourceInstance{
		JSONDefinition:     pvcManifest,
		WorkloadInstanceID: createdWorkloadInstance.ID,
	}
	_, err = tpclient.CreateWorkloadResourceInstance(
		r.APIClient,
		r.APIServer,
		&pvcWri,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create workload resource instance for persistent volume claim: %w", err)
	}

	// trigger reconciliation to create the PVC
	unreconciledWorkloadInstance := tpapi.WorkloadInstance{
		Common: tpapi.Common{
			ID: createdWorkloadInstance.ID,
		},
		Reconciliation: tpapi.Reconciliation{
			Reconciled: util.Ptr(false),
		},
	}
	_, err = tpclient.UpdateWorkloadInstance(
		r.APIClient,
		r.APIServer,
		&unreconciledWorkloadInstance,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to update workload instance as unreconciled: %w", err)
	}

	// create relational database instance if requested
	if *wordpressDefinition.ManagedDatabase {
		// get relational database definition ID
		awsRdsDefinitionId, err := tpclient.GetObjectIdByAttachedObject(
			r.APIClient,
			r.APIServer,
			tpapi.ObjectTypeAwsRelationalDatabaseDefinition,
			v0.ObjectTypeWordpressDefinition,
			*wordpressDefinition.ID,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to get AWS relational database definition ID by attachment: %w", err)
		}
		// construct relational database instance
		awsRdsInstance := tpapi.AwsRelationalDatabaseInstance{
			Instance: tpapi.Instance{
				Name: wordpressInstance.Name,
			},
			AwsRelationalDatabaseDefinitionID: awsRdsDefinitionId,
			WorkloadInstanceID:                createdWorkloadInstance.ID,
		}
		// create relational database instance
		createdRdsInstance, err := tpclient.CreateAwsRelationalDatabaseInstance(
			r.APIClient,
			r.APIServer,
			&awsRdsInstance,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to create AWS relational database instance: %w", err)
		}
		// establish attachment between wordpress instance and relational
		// database instance
		if err := tpclient.EnsureAttachedObjectReferenceExists(
			r.APIClient,
			r.APIServer,
			tpapi.ObjectTypeAwsRelationalDatabaseInstance,
			createdRdsInstance.ID,
			v0.ObjectTypeWordpressInstance,
			wordpressInstance.ID,
		); err != nil {
			return 0, fmt.Errorf("failed to attach wordpress instance to AWS relational database instance: %w", err)
		}
	}

	// create gateway and subdomain DNS record if requested
	if wordpressInstance.SubDomain != nil && *wordpressInstance.SubDomain != "" {
		// get attached domain name definition ID
		domainNameDefinitionId, err := tpclient.GetObjectIdByAttachedObject(
			r.APIClient,
			r.APIServer,
			tpapi.ObjectTypeDomainNameDefinition,
			v0.ObjectTypeWordpressDefinition,
			*wordpressDefinition.ID,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to get attached domain name definition: %w", err)
		}
		// contruct gateway definition object
		gatewayDefinition := tpapi.GatewayDefinition{
			Definition: tpapi.Definition{
				Name: util.Ptr("web-service-gateway"),
			},
			HttpPorts: []*tpapi.GatewayHttpPort{
				{
					Port:          util.Ptr(80),
					Path:          util.Ptr("/"),
					HTTPSRedirect: util.Ptr(true),
				}, {
					Port:       util.Ptr(443),
					Path:       util.Ptr("/"),
					TLSEnabled: util.Ptr(true),
				},
			},
			DomainNameDefinitionID: domainNameDefinitionId,
			SubDomain:              wordpressInstance.SubDomain,
			ServiceName:            util.Ptr(getWordpressServiceName(*wordpressDefinition.Name)),
			WorkloadDefinitionID:   workloadDefinitionId,
		}
		// create gateway definition
		createdGatewayDefinition, err := tpclient.CreateGatewayDefinition(
			r.APIClient,
			r.APIServer,
			&gatewayDefinition,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to create gateway definition: %w", err)
		}
		// construct gateway instance
		gatewayInstance := tpapi.GatewayInstance{
			Instance: tpapi.Instance{
				Name: util.Ptr(fmt.Sprintf("%s-gateway", *wordpressInstance.Name)),
			},
			KubernetesRuntimeInstanceID: kubernetesRuntimeInstanceId,
			GatewayDefinitionID:         createdGatewayDefinition.ID,
			WorkloadInstanceID:          createdWorkloadInstance.ID,
		}
		// created gateway instance
		createdGatewayInstance, err := tpclient.CreateGatewayInstance(
			r.APIClient,
			r.APIServer,
			&gatewayInstance,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to create gateway instance: %w", err)
		}
		// create attachment between gateway instance and wordpress instance
		if err := tpclient.EnsureAttachedObjectReferenceExists(
			r.APIClient,
			r.APIServer,
			tpapi.ObjectTypeGatewayInstance,
			createdGatewayInstance.ID,
			v0.ObjectTypeWordpressInstance,
			wordpressInstance.ID,
		); err != nil {
			return 0, fmt.Errorf("failed to create attachment between wordpress instance and gateway instance: %w", err)
		}
		// construct domain name instance
		domainNameInstance := tpapi.DomainNameInstance{
			Instance: tpapi.Instance{
				Name: util.Ptr(fmt.Sprintf("%s-domain-name", *wordpressInstance.Name)),
			},
			DomainNameDefinitionID:      domainNameDefinitionId,
			WorkloadInstanceID:          createdWorkloadInstance.ID,
			KubernetesRuntimeInstanceID: kubernetesRuntimeInstanceId,
		}
		// create domain name instance
		createdDomainNameInstance, err := tpclient.CreateDomainNameInstance(
			r.APIClient,
			r.APIServer,
			&domainNameInstance,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to create domain name instance: %w", err)
		}
		// create attachment between domain name instance and wordpress instance
		if err := tpclient.EnsureAttachedObjectReferenceExists(
			r.APIClient,
			r.APIServer,
			tpapi.ObjectTypeDomainNameInstance,
			createdDomainNameInstance.ID,
			v0.ObjectTypeWordpressInstance,
			wordpressInstance.ID,
		); err != nil {
			return 0, fmt.Errorf("failed to create attachment between wordpress instance and domain name instance: %w", err)
		}
	}

	return 0, nil
}

// v0WordpressInstanceUpdated performs reconciliation when a v0 WordpressInstance
// has been updated.
func v0WordpressInstanceUpdated(
	r *controller.Reconciler,
	wordpressInstance *v0.WordpressInstance,
	log *logr.Logger,
) (int64, error) {
	return 0, nil
}

// v0WordpressInstanceDeleted performs reconciliation when a v0 WordpressInstance
// has been deleted.
func v0WordpressInstanceDeleted(
	r *controller.Reconciler,
	wordpressInstance *v0.WordpressInstance,
	log *logr.Logger,
) (int64, error) {
	// get attached workload instance
	workloadInstanceId, err := tpclient.GetObjectIdByAttachedObject(
		r.APIClient,
		r.APIServer,
		tpapi.ObjectTypeWorkloadInstance,
		v0.ObjectTypeWordpressInstance,
		*wordpressInstance.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to find attached workload instance: %w", err)
	}

	// delete workload instance
	_, err = tpclient.DeleteWorkloadInstance(
		r.APIClient,
		r.APIServer,
		*workloadInstanceId,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to delete workload instance: %w", err)
	}

	// remove workload instance attachment
	if err := tpclient.EnsureAttachedObjectReferenceRemoved(
		r.APIClient,
		r.APIServer,
		tpapi.ObjectTypeWorkloadInstance,
		workloadInstanceId,
		v0.ObjectTypeWordpressInstance,
		wordpressInstance.ID,
	); err != nil {
		return 0, fmt.Errorf("failed to remove attachment to deleted workload instance: %w", err)
	}

	// delete relational database instance if it exists
	awsRdsInstanceIds, err := tpclient.GetObjectIdsByAttachedObject(
		r.APIClient,
		r.APIServer,
		tpapi.ObjectTypeAwsRelationalDatabaseInstance,
		v0.ObjectTypeWordpressInstance,
		*wordpressInstance.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get attached relational database IDs: %w", err)
	}
	for _, rdsInstanceId := range awsRdsInstanceIds {
		_, err := tpclient.DeleteAwsRelationalDatabaseInstance(
			r.APIClient,
			r.APIServer,
			*rdsInstanceId,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to delete AWS RDS instance with id %d: %w", *rdsInstanceId, err)
		}
		if err := tpclient.EnsureAttachedObjectReferenceRemoved(
			r.APIClient,
			r.APIServer,
			tpapi.ObjectTypeAwsRelationalDatabaseInstance,
			rdsInstanceId,
			v0.ObjectTypeWordpressInstance,
			wordpressInstance.ID,
		); err != nil {
			return 0, fmt.Errorf("failed to remove attachment to deleted AWS relational database instance: %w", err)
		}
	}

	return 0, nil
}

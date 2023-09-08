// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Elastic Organization To Azure Subscription Mapping details for the logged-in user.
func GetOrganizationElasticToAzureSubscriptionMapping(ctx *pulumi.Context, args *GetOrganizationElasticToAzureSubscriptionMappingArgs, opts ...pulumi.InvokeOption) (*GetOrganizationElasticToAzureSubscriptionMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetOrganizationElasticToAzureSubscriptionMappingResult
	err := ctx.Invoke("azure-native:elastic/v20230701preview:getOrganizationElasticToAzureSubscriptionMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetOrganizationElasticToAzureSubscriptionMappingArgs struct {
}

// The Azure Subscription ID to which the Organization of the logged in user belongs and gets billed into.
type GetOrganizationElasticToAzureSubscriptionMappingResult struct {
	// The properties of Azure Subscription ID to which the Organization of the logged in user belongs and gets billed into.
	Properties ElasticOrganizationToAzureSubscriptionMappingResponsePropertiesResponse `pulumi:"properties"`
}
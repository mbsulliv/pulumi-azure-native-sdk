// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns Connector resource for a given name.
func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:servicelinker/v20221101preview:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	// The name of resource.
	ConnectorName string `pulumi:"connectorName"`
	// The name of Azure region.
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the target subscription.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// Linker of source and target resource
type LookupConnectorResult struct {
	// The authentication type.
	AuthInfo interface{} `pulumi:"authInfo"`
	// The application client type
	ClientType *string `pulumi:"clientType"`
	// The connection information consumed by applications, including secrets, connection strings.
	ConfigurationInfo *ConfigurationInfoResponse `pulumi:"configurationInfo"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The network solution.
	PublicNetworkSolution *PublicNetworkSolutionResponse `pulumi:"publicNetworkSolution"`
	// connection scope in source service.
	Scope *string `pulumi:"scope"`
	// An option to store secret value in secure place
	SecretStore *SecretStoreResponse `pulumi:"secretStore"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The target service properties
	TargetService interface{} `pulumi:"targetService"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The VNet solution.
	VNetSolution *VNetSolutionResponse `pulumi:"vNetSolution"`
}

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			var s LookupConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	// The name of resource.
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
	// The name of Azure region.
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The ID of the target subscription.
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}

// Linker of source and target resource
type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

// The authentication type.
func (o LookupConnectorResultOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectorResult) interface{} { return v.AuthInfo }).(pulumi.AnyOutput)
}

// The application client type
func (o LookupConnectorResultOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.ClientType }).(pulumi.StringPtrOutput)
}

// The connection information consumed by applications, including secrets, connection strings.
func (o LookupConnectorResultOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *ConfigurationInfoResponse { return v.ConfigurationInfo }).(ConfigurationInfoResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o LookupConnectorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The network solution.
func (o LookupConnectorResultOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *PublicNetworkSolutionResponse { return v.PublicNetworkSolution }).(PublicNetworkSolutionResponsePtrOutput)
}

// connection scope in source service.
func (o LookupConnectorResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// An option to store secret value in secure place
func (o LookupConnectorResultOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *SecretStoreResponse { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The target service properties
func (o LookupConnectorResultOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectorResult) interface{} { return v.TargetService }).(pulumi.AnyOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VNet solution.
func (o LookupConnectorResultOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *VNetSolutionResponse { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}

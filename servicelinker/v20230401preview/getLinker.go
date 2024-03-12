// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns Linker resource for a given name.
func LookupLinker(ctx *pulumi.Context, args *LookupLinkerArgs, opts ...pulumi.InvokeOption) (*LookupLinkerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLinkerResult
	err := ctx.Invoke("azure-native:servicelinker/v20230401preview:getLinker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkerArgs struct {
	// The name Linker resource.
	LinkerName string `pulumi:"linkerName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri string `pulumi:"resourceUri"`
}

// Linker of source and target resource
type LookupLinkerResult struct {
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

func LookupLinkerOutput(ctx *pulumi.Context, args LookupLinkerOutputArgs, opts ...pulumi.InvokeOption) LookupLinkerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkerResult, error) {
			args := v.(LookupLinkerArgs)
			r, err := LookupLinker(ctx, &args, opts...)
			var s LookupLinkerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkerResultOutput)
}

type LookupLinkerOutputArgs struct {
	// The name Linker resource.
	LinkerName pulumi.StringInput `pulumi:"linkerName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupLinkerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerArgs)(nil)).Elem()
}

// Linker of source and target resource
type LookupLinkerResultOutput struct{ *pulumi.OutputState }

func (LookupLinkerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerResult)(nil)).Elem()
}

func (o LookupLinkerResultOutput) ToLookupLinkerResultOutput() LookupLinkerResultOutput {
	return o
}

func (o LookupLinkerResultOutput) ToLookupLinkerResultOutputWithContext(ctx context.Context) LookupLinkerResultOutput {
	return o
}

// The authentication type.
func (o LookupLinkerResultOutput) AuthInfo() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLinkerResult) interface{} { return v.AuthInfo }).(pulumi.AnyOutput)
}

// The application client type
func (o LookupLinkerResultOutput) ClientType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *string { return v.ClientType }).(pulumi.StringPtrOutput)
}

// The connection information consumed by applications, including secrets, connection strings.
func (o LookupLinkerResultOutput) ConfigurationInfo() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *ConfigurationInfoResponse { return v.ConfigurationInfo }).(ConfigurationInfoResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLinkerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLinkerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o LookupLinkerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The network solution.
func (o LookupLinkerResultOutput) PublicNetworkSolution() PublicNetworkSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *PublicNetworkSolutionResponse { return v.PublicNetworkSolution }).(PublicNetworkSolutionResponsePtrOutput)
}

// connection scope in source service.
func (o LookupLinkerResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// An option to store secret value in secure place
func (o LookupLinkerResultOutput) SecretStore() SecretStoreResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *SecretStoreResponse { return v.SecretStore }).(SecretStoreResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupLinkerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLinkerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The target service properties
func (o LookupLinkerResultOutput) TargetService() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupLinkerResult) interface{} { return v.TargetService }).(pulumi.AnyOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLinkerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VNet solution.
func (o LookupLinkerResultOutput) VNetSolution() VNetSolutionResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerResult) *VNetSolutionResponse { return v.VNetSolution }).(VNetSolutionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkerResultOutput{})
}

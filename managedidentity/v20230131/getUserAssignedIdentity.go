// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230131

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the identity.
func LookupUserAssignedIdentity(ctx *pulumi.Context, args *LookupUserAssignedIdentityArgs, opts ...pulumi.InvokeOption) (*LookupUserAssignedIdentityResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUserAssignedIdentityResult
	err := ctx.Invoke("azure-native:managedidentity/v20230131:getUserAssignedIdentity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserAssignedIdentityArgs struct {
	// The name of the Resource Group to which the identity belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the identity resource.
	ResourceName string `pulumi:"resourceName"`
}

// Describes an identity resource.
type LookupUserAssignedIdentityResult struct {
	// The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId string `pulumi:"clientId"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The id of the service principal object associated with the created identity.
	PrincipalId string `pulumi:"principalId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The id of the tenant which the identity belongs to.
	TenantId string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupUserAssignedIdentityOutput(ctx *pulumi.Context, args LookupUserAssignedIdentityOutputArgs, opts ...pulumi.InvokeOption) LookupUserAssignedIdentityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserAssignedIdentityResult, error) {
			args := v.(LookupUserAssignedIdentityArgs)
			r, err := LookupUserAssignedIdentity(ctx, &args, opts...)
			var s LookupUserAssignedIdentityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserAssignedIdentityResultOutput)
}

type LookupUserAssignedIdentityOutputArgs struct {
	// The name of the Resource Group to which the identity belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the identity resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupUserAssignedIdentityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserAssignedIdentityArgs)(nil)).Elem()
}

// Describes an identity resource.
type LookupUserAssignedIdentityResultOutput struct{ *pulumi.OutputState }

func (LookupUserAssignedIdentityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserAssignedIdentityResult)(nil)).Elem()
}

func (o LookupUserAssignedIdentityResultOutput) ToLookupUserAssignedIdentityResultOutput() LookupUserAssignedIdentityResultOutput {
	return o
}

func (o LookupUserAssignedIdentityResultOutput) ToLookupUserAssignedIdentityResultOutputWithContext(ctx context.Context) LookupUserAssignedIdentityResultOutput {
	return o
}

// The id of the app associated with the identity. This is a random generated UUID by MSI.
func (o LookupUserAssignedIdentityResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupUserAssignedIdentityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupUserAssignedIdentityResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupUserAssignedIdentityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the service principal object associated with the created identity.
func (o LookupUserAssignedIdentityResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupUserAssignedIdentityResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupUserAssignedIdentityResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The id of the tenant which the identity belongs to.
func (o LookupUserAssignedIdentityResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupUserAssignedIdentityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserAssignedIdentityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserAssignedIdentityResultOutput{})
}

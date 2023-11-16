// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing KeyGroup within a profile.
func LookupKeyGroup(ctx *pulumi.Context, args *LookupKeyGroupArgs, opts ...pulumi.InvokeOption) (*LookupKeyGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKeyGroupResult
	err := ctx.Invoke("azure-native:cdn/v20230701preview:getKeyGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeyGroupArgs struct {
	// Name of the KeyGroup under the profile.
	KeyGroupName string `pulumi:"keyGroupName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Contains a list of references of UrlSigningKey type secret objects.
type LookupKeyGroupResult struct {
	DeploymentStatus string `pulumi:"deploymentStatus"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Names of UrlSigningKey type secret objects
	KeyReferences []ResourceReferenceResponse `pulumi:"keyReferences"`
	// Resource name.
	Name string `pulumi:"name"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupKeyGroupOutput(ctx *pulumi.Context, args LookupKeyGroupOutputArgs, opts ...pulumi.InvokeOption) LookupKeyGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeyGroupResult, error) {
			args := v.(LookupKeyGroupArgs)
			r, err := LookupKeyGroup(ctx, &args, opts...)
			var s LookupKeyGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKeyGroupResultOutput)
}

type LookupKeyGroupOutputArgs struct {
	// Name of the KeyGroup under the profile.
	KeyGroupName pulumi.StringInput `pulumi:"keyGroupName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKeyGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyGroupArgs)(nil)).Elem()
}

// Contains a list of references of UrlSigningKey type secret objects.
type LookupKeyGroupResultOutput struct{ *pulumi.OutputState }

func (LookupKeyGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeyGroupResult)(nil)).Elem()
}

func (o LookupKeyGroupResultOutput) ToLookupKeyGroupResultOutput() LookupKeyGroupResultOutput {
	return o
}

func (o LookupKeyGroupResultOutput) ToLookupKeyGroupResultOutputWithContext(ctx context.Context) LookupKeyGroupResultOutput {
	return o
}

func (o LookupKeyGroupResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyGroupResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupKeyGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Names of UrlSigningKey type secret objects
func (o LookupKeyGroupResultOutput) KeyReferences() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupKeyGroupResult) []ResourceReferenceResponse { return v.KeyReferences }).(ResourceReferenceResponseArrayOutput)
}

// Resource name.
func (o LookupKeyGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning status
func (o LookupKeyGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o LookupKeyGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKeyGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupKeyGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeyGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeyGroupResultOutput{})
}

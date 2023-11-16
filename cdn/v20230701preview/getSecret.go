// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing Secret within a profile.
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecretResult
	err := ctx.Invoke("azure-native:cdn/v20230701preview:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecretArgs struct {
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Secret under the profile.
	SecretName string `pulumi:"secretName"`
}

// Friendly Secret name mapping to the any Secret or secret related information.
type LookupSecretResult struct {
	DeploymentStatus string `pulumi:"deploymentStatus"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// object which contains secret parameters
	Parameters interface{} `pulumi:"parameters"`
	// The name of the profile which holds the secret.
	ProfileName string `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretResult, error) {
			args := v.(LookupSecretArgs)
			r, err := LookupSecret(ctx, &args, opts...)
			var s LookupSecretResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretResultOutput)
}

type LookupSecretOutputArgs struct {
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Secret under the profile.
	SecretName pulumi.StringInput `pulumi:"secretName"`
}

func (LookupSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretArgs)(nil)).Elem()
}

// Friendly Secret name mapping to the any Secret or secret related information.
type LookupSecretResultOutput struct{ *pulumi.OutputState }

func (LookupSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretResult)(nil)).Elem()
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutput() LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutputWithContext(ctx context.Context) LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupSecretResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSecretResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Name }).(pulumi.StringOutput)
}

// object which contains secret parameters
func (o LookupSecretResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSecretResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

// The name of the profile which holds the secret.
func (o LookupSecretResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o LookupSecretResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o LookupSecretResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecretResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupSecretResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}

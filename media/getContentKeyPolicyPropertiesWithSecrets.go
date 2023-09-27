// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Content Key Policy including secret values
// Azure REST API version: 2023-01-01.
func GetContentKeyPolicyPropertiesWithSecrets(ctx *pulumi.Context, args *GetContentKeyPolicyPropertiesWithSecretsArgs, opts ...pulumi.InvokeOption) (*GetContentKeyPolicyPropertiesWithSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetContentKeyPolicyPropertiesWithSecretsResult
	err := ctx.Invoke("azure-native:media:getContentKeyPolicyPropertiesWithSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetContentKeyPolicyPropertiesWithSecretsArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Content Key Policy name.
	ContentKeyPolicyName string `pulumi:"contentKeyPolicyName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The properties of the Content Key Policy.
type GetContentKeyPolicyPropertiesWithSecretsResult struct {
	// The creation date of the Policy
	Created string `pulumi:"created"`
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The last modified date of the Policy
	LastModified string `pulumi:"lastModified"`
	// The Key Policy options.
	Options []ContentKeyPolicyOptionResponse `pulumi:"options"`
	// The legacy Policy ID.
	PolicyId string `pulumi:"policyId"`
}

func GetContentKeyPolicyPropertiesWithSecretsOutput(ctx *pulumi.Context, args GetContentKeyPolicyPropertiesWithSecretsOutputArgs, opts ...pulumi.InvokeOption) GetContentKeyPolicyPropertiesWithSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContentKeyPolicyPropertiesWithSecretsResult, error) {
			args := v.(GetContentKeyPolicyPropertiesWithSecretsArgs)
			r, err := GetContentKeyPolicyPropertiesWithSecrets(ctx, &args, opts...)
			var s GetContentKeyPolicyPropertiesWithSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContentKeyPolicyPropertiesWithSecretsResultOutput)
}

type GetContentKeyPolicyPropertiesWithSecretsOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The Content Key Policy name.
	ContentKeyPolicyName pulumi.StringInput `pulumi:"contentKeyPolicyName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetContentKeyPolicyPropertiesWithSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContentKeyPolicyPropertiesWithSecretsArgs)(nil)).Elem()
}

// The properties of the Content Key Policy.
type GetContentKeyPolicyPropertiesWithSecretsResultOutput struct{ *pulumi.OutputState }

func (GetContentKeyPolicyPropertiesWithSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContentKeyPolicyPropertiesWithSecretsResult)(nil)).Elem()
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) ToGetContentKeyPolicyPropertiesWithSecretsResultOutput() GetContentKeyPolicyPropertiesWithSecretsResultOutput {
	return o
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) ToGetContentKeyPolicyPropertiesWithSecretsResultOutputWithContext(ctx context.Context) GetContentKeyPolicyPropertiesWithSecretsResultOutput {
	return o
}

func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetContentKeyPolicyPropertiesWithSecretsResult] {
	return pulumix.Output[GetContentKeyPolicyPropertiesWithSecretsResult]{
		OutputState: o.OutputState,
	}
}

// The creation date of the Policy
func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) string { return v.Created }).(pulumi.StringOutput)
}

// A description for the Policy.
func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The last modified date of the Policy
func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The Key Policy options.
func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) Options() ContentKeyPolicyOptionResponseArrayOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) []ContentKeyPolicyOptionResponse {
		return v.Options
	}).(ContentKeyPolicyOptionResponseArrayOutput)
}

// The legacy Policy ID.
func (o GetContentKeyPolicyPropertiesWithSecretsResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetContentKeyPolicyPropertiesWithSecretsResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContentKeyPolicyPropertiesWithSecretsResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datareplication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the policy.
// Azure REST API version: 2021-02-16-preview.
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:datareplication:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	// The policy name.
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName string `pulumi:"vaultName"`
}

// Policy model.
type LookupPolicyResult struct {
	// Gets or sets the Id of the resource.
	Id string `pulumi:"id"`
	// Gets or sets the name of the resource.
	Name string `pulumi:"name"`
	// Policy model properties.
	Properties PolicyModelPropertiesResponse `pulumi:"properties"`
	SystemData PolicyModelResponseSystemData `pulumi:"systemData"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	// The policy name.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// Policy model.
type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyResult] {
	return pulumix.Output[LookupPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the Id of the resource.
func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the name of the resource.
func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy model properties.
func (o LookupPolicyResultOutput) Properties() PolicyModelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPolicyResult) PolicyModelPropertiesResponse { return v.Properties }).(PolicyModelPropertiesResponseOutput)
}

func (o LookupPolicyResultOutput) SystemData() PolicyModelResponseSystemDataOutput {
	return o.ApplyT(func(v LookupPolicyResult) PolicyModelResponseSystemData { return v.SystemData }).(PolicyModelResponseSystemDataOutput)
}

// Gets or sets the type of the resource.
func (o LookupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}

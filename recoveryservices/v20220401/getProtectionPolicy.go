// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Base class for backup policy. Workload-specific backup policies are derived from this class.
func LookupProtectionPolicy(ctx *pulumi.Context, args *LookupProtectionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProtectionPolicyResult, error) {
	var rv LookupProtectionPolicyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20220401:getProtectionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionPolicyArgs struct {
	// Backup policy information to be fetched.
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// Base class for backup policy. Workload-specific backup policies are derived from this class.
type LookupProtectionPolicyResult struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// ProtectionPolicyResource properties
	Properties interface{} `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func LookupProtectionPolicyOutput(ctx *pulumi.Context, args LookupProtectionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupProtectionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectionPolicyResult, error) {
			args := v.(LookupProtectionPolicyArgs)
			r, err := LookupProtectionPolicy(ctx, &args, opts...)
			var s LookupProtectionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProtectionPolicyResultOutput)
}

type LookupProtectionPolicyOutputArgs struct {
	// Backup policy information to be fetched.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupProtectionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionPolicyArgs)(nil)).Elem()
}

// Base class for backup policy. Workload-specific backup policies are derived from this class.
type LookupProtectionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupProtectionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionPolicyResult)(nil)).Elem()
}

func (o LookupProtectionPolicyResultOutput) ToLookupProtectionPolicyResultOutput() LookupProtectionPolicyResultOutput {
	return o
}

func (o LookupProtectionPolicyResultOutput) ToLookupProtectionPolicyResultOutputWithContext(ctx context.Context) LookupProtectionPolicyResultOutput {
	return o
}

// Optional ETag.
func (o LookupProtectionPolicyResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id represents the complete path to the resource.
func (o LookupProtectionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupProtectionPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o LookupProtectionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// ProtectionPolicyResource properties
func (o LookupProtectionPolicyResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags.
func (o LookupProtectionPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o LookupProtectionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectionPolicyResultOutput{})
}
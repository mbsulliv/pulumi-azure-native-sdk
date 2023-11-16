// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a replication policy.
func LookupReplicationPolicy(ctx *pulumi.Context, args *LookupReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupReplicationPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationPolicyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20230401:getReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationPolicyArgs struct {
	// Replication policy name.
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Protection profile details.
type LookupReplicationPolicyResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// The custom data.
	Properties PolicyPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationPolicyOutput(ctx *pulumi.Context, args LookupReplicationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationPolicyResult, error) {
			args := v.(LookupReplicationPolicyArgs)
			r, err := LookupReplicationPolicy(ctx, &args, opts...)
			var s LookupReplicationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationPolicyResultOutput)
}

type LookupReplicationPolicyOutputArgs struct {
	// Replication policy name.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationPolicyArgs)(nil)).Elem()
}

// Protection profile details.
type LookupReplicationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationPolicyResult)(nil)).Elem()
}

func (o LookupReplicationPolicyResultOutput) ToLookupReplicationPolicyResultOutput() LookupReplicationPolicyResultOutput {
	return o
}

func (o LookupReplicationPolicyResultOutput) ToLookupReplicationPolicyResultOutputWithContext(ctx context.Context) LookupReplicationPolicyResultOutput {
	return o
}

// Resource Id
func (o LookupReplicationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The custom data.
func (o LookupReplicationPolicyResultOutput) Properties() PolicyPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) PolicyPropertiesResponse { return v.Properties }).(PolicyPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationPolicyResultOutput{})
}

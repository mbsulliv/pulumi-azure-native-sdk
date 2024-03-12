// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the database data masking policy.
// Azure REST API version: 2021-11-01.
//
// Other available API versions: 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview.
func LookupDataMaskingPolicy(ctx *pulumi.Context, args *LookupDataMaskingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDataMaskingPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataMaskingPolicyResult
	err := ctx.Invoke("azure-native:sql:getDataMaskingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataMaskingPolicyArgs struct {
	// The name of the database for which the data masking policy applies.
	DataMaskingPolicyName string `pulumi:"dataMaskingPolicyName"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A database data masking policy.
type LookupDataMaskingPolicyResult struct {
	// The list of the application principals. This is a legacy parameter and is no longer used.
	ApplicationPrincipals string `pulumi:"applicationPrincipals"`
	// The state of the data masking policy.
	DataMaskingState string `pulumi:"dataMaskingState"`
	// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
	ExemptPrincipals *string `pulumi:"exemptPrincipals"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The kind of Data Masking Policy. Metadata, used for Azure portal.
	Kind string `pulumi:"kind"`
	// The location of the data masking policy.
	Location string `pulumi:"location"`
	// The masking level. This is a legacy parameter and is no longer used.
	MaskingLevel string `pulumi:"maskingLevel"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupDataMaskingPolicyOutput(ctx *pulumi.Context, args LookupDataMaskingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDataMaskingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataMaskingPolicyResult, error) {
			args := v.(LookupDataMaskingPolicyArgs)
			r, err := LookupDataMaskingPolicy(ctx, &args, opts...)
			var s LookupDataMaskingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataMaskingPolicyResultOutput)
}

type LookupDataMaskingPolicyOutputArgs struct {
	// The name of the database for which the data masking policy applies.
	DataMaskingPolicyName pulumi.StringInput `pulumi:"dataMaskingPolicyName"`
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDataMaskingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataMaskingPolicyArgs)(nil)).Elem()
}

// A database data masking policy.
type LookupDataMaskingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDataMaskingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataMaskingPolicyResult)(nil)).Elem()
}

func (o LookupDataMaskingPolicyResultOutput) ToLookupDataMaskingPolicyResultOutput() LookupDataMaskingPolicyResultOutput {
	return o
}

func (o LookupDataMaskingPolicyResultOutput) ToLookupDataMaskingPolicyResultOutputWithContext(ctx context.Context) LookupDataMaskingPolicyResultOutput {
	return o
}

// The list of the application principals. This is a legacy parameter and is no longer used.
func (o LookupDataMaskingPolicyResultOutput) ApplicationPrincipals() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.ApplicationPrincipals }).(pulumi.StringOutput)
}

// The state of the data masking policy.
func (o LookupDataMaskingPolicyResultOutput) DataMaskingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.DataMaskingState }).(pulumi.StringOutput)
}

// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
func (o LookupDataMaskingPolicyResultOutput) ExemptPrincipals() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) *string { return v.ExemptPrincipals }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupDataMaskingPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of Data Masking Policy. Metadata, used for Azure portal.
func (o LookupDataMaskingPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The location of the data masking policy.
func (o LookupDataMaskingPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// The masking level. This is a legacy parameter and is no longer used.
func (o LookupDataMaskingPolicyResultOutput) MaskingLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.MaskingLevel }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupDataMaskingPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupDataMaskingPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataMaskingPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataMaskingPolicyResultOutput{})
}

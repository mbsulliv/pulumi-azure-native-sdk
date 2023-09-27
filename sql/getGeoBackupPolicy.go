// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Geo backup policy for the given database resource.
// Azure REST API version: 2021-11-01.
func LookupGeoBackupPolicy(ctx *pulumi.Context, args *LookupGeoBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGeoBackupPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGeoBackupPolicyResult
	err := ctx.Invoke("azure-native:sql:getGeoBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGeoBackupPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Geo backup policy. This should always be 'Default'.
	GeoBackupPolicyName string `pulumi:"geoBackupPolicyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A Geo backup policy.
type LookupGeoBackupPolicyResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Kind of geo backup policy.  This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Backup policy location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The state of the geo backup policy.
	State string `pulumi:"state"`
	// The storage type of the geo backup policy.
	StorageType string `pulumi:"storageType"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupGeoBackupPolicyOutput(ctx *pulumi.Context, args LookupGeoBackupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGeoBackupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGeoBackupPolicyResult, error) {
			args := v.(LookupGeoBackupPolicyArgs)
			r, err := LookupGeoBackupPolicy(ctx, &args, opts...)
			var s LookupGeoBackupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGeoBackupPolicyResultOutput)
}

type LookupGeoBackupPolicyOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the Geo backup policy. This should always be 'Default'.
	GeoBackupPolicyName pulumi.StringInput `pulumi:"geoBackupPolicyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupGeoBackupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGeoBackupPolicyArgs)(nil)).Elem()
}

// A Geo backup policy.
type LookupGeoBackupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGeoBackupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGeoBackupPolicyResult)(nil)).Elem()
}

func (o LookupGeoBackupPolicyResultOutput) ToLookupGeoBackupPolicyResultOutput() LookupGeoBackupPolicyResultOutput {
	return o
}

func (o LookupGeoBackupPolicyResultOutput) ToLookupGeoBackupPolicyResultOutputWithContext(ctx context.Context) LookupGeoBackupPolicyResultOutput {
	return o
}

func (o LookupGeoBackupPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGeoBackupPolicyResult] {
	return pulumix.Output[LookupGeoBackupPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Resource ID.
func (o LookupGeoBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of geo backup policy.  This is metadata used for the Azure portal experience.
func (o LookupGeoBackupPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Backup policy location.
func (o LookupGeoBackupPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupGeoBackupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The state of the geo backup policy.
func (o LookupGeoBackupPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

// The storage type of the geo backup policy.
func (o LookupGeoBackupPolicyResultOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.StorageType }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupGeoBackupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGeoBackupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGeoBackupPolicyResultOutput{})
}

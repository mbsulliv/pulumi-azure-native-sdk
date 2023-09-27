// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the properties of the specified replication.
// Azure REST API version: 2022-12-01.
func LookupReplication(ctx *pulumi.Context, args *LookupReplicationArgs, opts ...pulumi.InvokeOption) (*LookupReplicationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationResult
	err := ctx.Invoke("azure-native:containerregistry:getReplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupReplicationArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the replication.
	ReplicationName string `pulumi:"replicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a replication for a container registry.
type LookupReplicationResult struct {
	// The resource ID.
	Id string `pulumi:"id"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning state of the replication at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// Specifies whether the replication's regional endpoint is enabled. Requests will not be routed to a replication whose regional endpoint is disabled, however its data will continue to be synced with other replications.
	RegionEndpointEnabled *bool `pulumi:"regionEndpointEnabled"`
	// The status of the replication at the time the operation was called.
	Status StatusResponse `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Whether or not zone redundancy is enabled for this container registry replication
	ZoneRedundancy *string `pulumi:"zoneRedundancy"`
}

// Defaults sets the appropriate defaults for LookupReplicationResult
func (val *LookupReplicationResult) Defaults() *LookupReplicationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.RegionEndpointEnabled == nil {
		regionEndpointEnabled_ := true
		tmp.RegionEndpointEnabled = &regionEndpointEnabled_
	}
	if tmp.ZoneRedundancy == nil {
		zoneRedundancy_ := "Disabled"
		tmp.ZoneRedundancy = &zoneRedundancy_
	}
	return &tmp
}

func LookupReplicationOutput(ctx *pulumi.Context, args LookupReplicationOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationResult, error) {
			args := v.(LookupReplicationArgs)
			r, err := LookupReplication(ctx, &args, opts...)
			var s LookupReplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationResultOutput)
}

type LookupReplicationOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the replication.
	ReplicationName pulumi.StringInput `pulumi:"replicationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationArgs)(nil)).Elem()
}

// An object that represents a replication for a container registry.
type LookupReplicationResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationResult)(nil)).Elem()
}

func (o LookupReplicationResultOutput) ToLookupReplicationResultOutput() LookupReplicationResultOutput {
	return o
}

func (o LookupReplicationResultOutput) ToLookupReplicationResultOutputWithContext(ctx context.Context) LookupReplicationResultOutput {
	return o
}

func (o LookupReplicationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReplicationResult] {
	return pulumix.Output[LookupReplicationResult]{
		OutputState: o.OutputState,
	}
}

// The resource ID.
func (o LookupReplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o LookupReplicationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupReplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the replication at the time the operation was called.
func (o LookupReplicationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies whether the replication's regional endpoint is enabled. Requests will not be routed to a replication whose regional endpoint is disabled, however its data will continue to be synced with other replications.
func (o LookupReplicationResultOutput) RegionEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupReplicationResult) *bool { return v.RegionEndpointEnabled }).(pulumi.BoolPtrOutput)
}

// The status of the replication at the time the operation was called.
func (o LookupReplicationResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupReplicationResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupReplicationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupReplicationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the resource.
func (o LookupReplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupReplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether or not zone redundancy is enabled for this container registry replication
func (o LookupReplicationResultOutput) ZoneRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationResult) *string { return v.ZoneRedundancy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationResultOutput{})
}

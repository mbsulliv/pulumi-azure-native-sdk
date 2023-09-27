// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the properties of the specified replica.
func LookupReplica(ctx *pulumi.Context, args *LookupReplicaArgs, opts ...pulumi.InvokeOption) (*LookupReplicaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicaResult
	err := ctx.Invoke("azure-native:appconfiguration/v20230301:getReplica", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicaArgs struct {
	// The name of the configuration store.
	ConfigStoreName string `pulumi:"configStoreName"`
	// The name of the replica.
	ReplicaName string `pulumi:"replicaName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The replica resource.
type LookupReplicaResult struct {
	// The URI of the replica where the replica API will be available.
	Endpoint string `pulumi:"endpoint"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The location of the replica.
	Location *string `pulumi:"location"`
	// The name of the replica.
	Name string `pulumi:"name"`
	// The provisioning state of the replica.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource system metadata.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupReplicaOutput(ctx *pulumi.Context, args LookupReplicaOutputArgs, opts ...pulumi.InvokeOption) LookupReplicaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicaResult, error) {
			args := v.(LookupReplicaArgs)
			r, err := LookupReplica(ctx, &args, opts...)
			var s LookupReplicaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicaResultOutput)
}

type LookupReplicaOutputArgs struct {
	// The name of the configuration store.
	ConfigStoreName pulumi.StringInput `pulumi:"configStoreName"`
	// The name of the replica.
	ReplicaName pulumi.StringInput `pulumi:"replicaName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupReplicaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicaArgs)(nil)).Elem()
}

// The replica resource.
type LookupReplicaResultOutput struct{ *pulumi.OutputState }

func (LookupReplicaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicaResult)(nil)).Elem()
}

func (o LookupReplicaResultOutput) ToLookupReplicaResultOutput() LookupReplicaResultOutput {
	return o
}

func (o LookupReplicaResultOutput) ToLookupReplicaResultOutputWithContext(ctx context.Context) LookupReplicaResultOutput {
	return o
}

func (o LookupReplicaResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReplicaResult] {
	return pulumix.Output[LookupReplicaResult]{
		OutputState: o.OutputState,
	}
}

// The URI of the replica where the replica API will be available.
func (o LookupReplicaResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// The resource ID.
func (o LookupReplicaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the replica.
func (o LookupReplicaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the replica.
func (o LookupReplicaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the replica.
func (o LookupReplicaResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource system metadata.
func (o LookupReplicaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupReplicaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupReplicaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicaResultOutput{})
}

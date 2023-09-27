// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The resource representation of a service topology.
// Azure REST API version: 2019-11-01-preview.
func LookupServiceTopology(ctx *pulumi.Context, args *LookupServiceTopologyArgs, opts ...pulumi.InvokeOption) (*LookupServiceTopologyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceTopologyResult
	err := ctx.Invoke("azure-native:deploymentmanager:getServiceTopology", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceTopologyArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service topology .
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
}

// The resource representation of a service topology.
type LookupServiceTopologyResult struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupServiceTopologyOutput(ctx *pulumi.Context, args LookupServiceTopologyOutputArgs, opts ...pulumi.InvokeOption) LookupServiceTopologyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceTopologyResult, error) {
			args := v.(LookupServiceTopologyArgs)
			r, err := LookupServiceTopology(ctx, &args, opts...)
			var s LookupServiceTopologyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceTopologyResultOutput)
}

type LookupServiceTopologyOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service topology .
	ServiceTopologyName pulumi.StringInput `pulumi:"serviceTopologyName"`
}

func (LookupServiceTopologyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceTopologyArgs)(nil)).Elem()
}

// The resource representation of a service topology.
type LookupServiceTopologyResultOutput struct{ *pulumi.OutputState }

func (LookupServiceTopologyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceTopologyResult)(nil)).Elem()
}

func (o LookupServiceTopologyResultOutput) ToLookupServiceTopologyResultOutput() LookupServiceTopologyResultOutput {
	return o
}

func (o LookupServiceTopologyResultOutput) ToLookupServiceTopologyResultOutputWithContext(ctx context.Context) LookupServiceTopologyResultOutput {
	return o
}

func (o LookupServiceTopologyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServiceTopologyResult] {
	return pulumix.Output[LookupServiceTopologyResult]{
		OutputState: o.OutputState,
	}
}

// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
func (o LookupServiceTopologyResultOutput) ArtifactSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) *string { return v.ArtifactSourceId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServiceTopologyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupServiceTopologyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupServiceTopologyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupServiceTopologyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServiceTopologyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceTopologyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceTopologyResultOutput{})
}

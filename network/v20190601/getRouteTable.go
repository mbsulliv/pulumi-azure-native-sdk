// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified route table.
func LookupRouteTable(ctx *pulumi.Context, args *LookupRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupRouteTableResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteTableResult
	err := ctx.Invoke("azure-native:network/v20190601:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteTableArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route table.
	RouteTableName string `pulumi:"routeTableName"`
}

// Route table resource.
type LookupRouteTableResult struct {
	// Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation *bool `pulumi:"disableBgpRoutePropagation"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Collection of routes contained within a route table.
	Routes []RouteResponse `pulumi:"routes"`
	// A collection of references to subnets.
	Subnets []SubnetResponse `pulumi:"subnets"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRouteTableOutput(ctx *pulumi.Context, args LookupRouteTableOutputArgs, opts ...pulumi.InvokeOption) LookupRouteTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteTableResult, error) {
			args := v.(LookupRouteTableArgs)
			r, err := LookupRouteTable(ctx, &args, opts...)
			var s LookupRouteTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteTableResultOutput)
}

type LookupRouteTableOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the route table.
	RouteTableName pulumi.StringInput `pulumi:"routeTableName"`
}

func (LookupRouteTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableArgs)(nil)).Elem()
}

// Route table resource.
type LookupRouteTableResultOutput struct{ *pulumi.OutputState }

func (LookupRouteTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableResult)(nil)).Elem()
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutput() LookupRouteTableResultOutput {
	return o
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutputWithContext(ctx context.Context) LookupRouteTableResultOutput {
	return o
}

func (o LookupRouteTableResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRouteTableResult] {
	return pulumix.Output[LookupRouteTableResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
func (o LookupRouteTableResultOutput) DisableBgpRoutePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *bool { return v.DisableBgpRoutePropagation }).(pulumi.BoolPtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o LookupRouteTableResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupRouteTableResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupRouteTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupRouteTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o LookupRouteTableResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Collection of routes contained within a route table.
func (o LookupRouteTableResultOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []RouteResponse { return v.Routes }).(RouteResponseArrayOutput)
}

// A collection of references to subnets.
func (o LookupRouteTableResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o LookupRouteTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupRouteTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteTableResultOutput{})
}

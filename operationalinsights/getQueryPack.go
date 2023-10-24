// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns a Log Analytics QueryPack.
// Azure REST API version: 2019-09-01.
//
// Other available API versions: 2019-09-01-preview.
func LookupQueryPack(ctx *pulumi.Context, args *LookupQueryPackArgs, opts ...pulumi.InvokeOption) (*LookupQueryPackResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupQueryPackResult
	err := ctx.Invoke("azure-native:operationalinsights:getQueryPack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueryPackArgs struct {
	// The name of the Log Analytics QueryPack resource.
	QueryPackName string `pulumi:"queryPackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Log Analytics QueryPack definition.
type LookupQueryPackResult struct {
	// Azure resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Current state of this QueryPack: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState string `pulumi:"provisioningState"`
	// The unique ID of your application. This field cannot be changed.
	QueryPackId string `pulumi:"queryPackId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Creation Date for the Log Analytics QueryPack, in ISO 8601 format.
	TimeCreated string `pulumi:"timeCreated"`
	// Last modified date of the Log Analytics QueryPack, in ISO 8601 format.
	TimeModified string `pulumi:"timeModified"`
	// Azure resource type
	Type string `pulumi:"type"`
}

func LookupQueryPackOutput(ctx *pulumi.Context, args LookupQueryPackOutputArgs, opts ...pulumi.InvokeOption) LookupQueryPackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueryPackResult, error) {
			args := v.(LookupQueryPackArgs)
			r, err := LookupQueryPack(ctx, &args, opts...)
			var s LookupQueryPackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueryPackResultOutput)
}

type LookupQueryPackOutputArgs struct {
	// The name of the Log Analytics QueryPack resource.
	QueryPackName pulumi.StringInput `pulumi:"queryPackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueryPackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueryPackArgs)(nil)).Elem()
}

// An Log Analytics QueryPack definition.
type LookupQueryPackResultOutput struct{ *pulumi.OutputState }

func (LookupQueryPackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueryPackResult)(nil)).Elem()
}

func (o LookupQueryPackResultOutput) ToLookupQueryPackResultOutput() LookupQueryPackResultOutput {
	return o
}

func (o LookupQueryPackResultOutput) ToLookupQueryPackResultOutputWithContext(ctx context.Context) LookupQueryPackResultOutput {
	return o
}

func (o LookupQueryPackResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupQueryPackResult] {
	return pulumix.Output[LookupQueryPackResult]{
		OutputState: o.OutputState,
	}
}

// Azure resource Id
func (o LookupQueryPackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupQueryPackResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o LookupQueryPackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Name }).(pulumi.StringOutput)
}

// Current state of this QueryPack: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
func (o LookupQueryPackResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The unique ID of your application. This field cannot be changed.
func (o LookupQueryPackResultOutput) QueryPackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.QueryPackId }).(pulumi.StringOutput)
}

// Resource tags
func (o LookupQueryPackResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupQueryPackResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Creation Date for the Log Analytics QueryPack, in ISO 8601 format.
func (o LookupQueryPackResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// Last modified date of the Log Analytics QueryPack, in ISO 8601 format.
func (o LookupQueryPackResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o LookupQueryPackResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueryPackResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueryPackResultOutput{})
}

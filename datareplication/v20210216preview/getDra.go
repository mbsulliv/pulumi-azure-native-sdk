// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210216preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the fabric agent.
func LookupDra(ctx *pulumi.Context, args *LookupDraArgs, opts ...pulumi.InvokeOption) (*LookupDraResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDraResult
	err := ctx.Invoke("azure-native:datareplication/v20210216preview:getDra", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDraArgs struct {
	// The fabric agent (Dra) name.
	FabricAgentName string `pulumi:"fabricAgentName"`
	// The fabric name.
	FabricName string `pulumi:"fabricName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Dra model.
type LookupDraResult struct {
	// Gets or sets the Id of the resource.
	Id string `pulumi:"id"`
	// Gets or sets the name of the resource.
	Name string `pulumi:"name"`
	// Dra model properties.
	Properties DraModelPropertiesResponse `pulumi:"properties"`
	SystemData DraModelResponseSystemData `pulumi:"systemData"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
}

func LookupDraOutput(ctx *pulumi.Context, args LookupDraOutputArgs, opts ...pulumi.InvokeOption) LookupDraResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDraResult, error) {
			args := v.(LookupDraArgs)
			r, err := LookupDra(ctx, &args, opts...)
			var s LookupDraResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDraResultOutput)
}

type LookupDraOutputArgs struct {
	// The fabric agent (Dra) name.
	FabricAgentName pulumi.StringInput `pulumi:"fabricAgentName"`
	// The fabric name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDraOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDraArgs)(nil)).Elem()
}

// Dra model.
type LookupDraResultOutput struct{ *pulumi.OutputState }

func (LookupDraResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDraResult)(nil)).Elem()
}

func (o LookupDraResultOutput) ToLookupDraResultOutput() LookupDraResultOutput {
	return o
}

func (o LookupDraResultOutput) ToLookupDraResultOutputWithContext(ctx context.Context) LookupDraResultOutput {
	return o
}

func (o LookupDraResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDraResult] {
	return pulumix.Output[LookupDraResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the Id of the resource.
func (o LookupDraResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDraResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the name of the resource.
func (o LookupDraResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDraResult) string { return v.Name }).(pulumi.StringOutput)
}

// Dra model properties.
func (o LookupDraResultOutput) Properties() DraModelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDraResult) DraModelPropertiesResponse { return v.Properties }).(DraModelPropertiesResponseOutput)
}

func (o LookupDraResultOutput) SystemData() DraModelResponseSystemDataOutput {
	return o.ApplyT(func(v LookupDraResult) DraModelResponseSystemData { return v.SystemData }).(DraModelResponseSystemDataOutput)
}

// Gets or sets the type of the resource.
func (o LookupDraResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDraResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDraResultOutput{})
}

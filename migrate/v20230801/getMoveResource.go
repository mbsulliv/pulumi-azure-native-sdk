// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Move Resource.
func LookupMoveResource(ctx *pulumi.Context, args *LookupMoveResourceArgs, opts ...pulumi.InvokeOption) (*LookupMoveResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMoveResourceResult
	err := ctx.Invoke("azure-native:migrate/v20230801:getMoveResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMoveResourceArgs struct {
	// The Move Collection Name.
	MoveCollectionName string `pulumi:"moveCollectionName"`
	// The Move Resource Name.
	MoveResourceName string `pulumi:"moveResourceName"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the move resource.
type LookupMoveResourceResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the move resource properties.
	Properties MoveResourcePropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupMoveResourceOutput(ctx *pulumi.Context, args LookupMoveResourceOutputArgs, opts ...pulumi.InvokeOption) LookupMoveResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMoveResourceResult, error) {
			args := v.(LookupMoveResourceArgs)
			r, err := LookupMoveResource(ctx, &args, opts...)
			var s LookupMoveResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMoveResourceResultOutput)
}

type LookupMoveResourceOutputArgs struct {
	// The Move Collection Name.
	MoveCollectionName pulumi.StringInput `pulumi:"moveCollectionName"`
	// The Move Resource Name.
	MoveResourceName pulumi.StringInput `pulumi:"moveResourceName"`
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMoveResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMoveResourceArgs)(nil)).Elem()
}

// Defines the move resource.
type LookupMoveResourceResultOutput struct{ *pulumi.OutputState }

func (LookupMoveResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMoveResourceResult)(nil)).Elem()
}

func (o LookupMoveResourceResultOutput) ToLookupMoveResourceResultOutput() LookupMoveResourceResultOutput {
	return o
}

func (o LookupMoveResourceResultOutput) ToLookupMoveResourceResultOutputWithContext(ctx context.Context) LookupMoveResourceResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupMoveResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMoveResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the move resource properties.
func (o LookupMoveResourceResultOutput) Properties() MoveResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) MoveResourcePropertiesResponse { return v.Properties }).(MoveResourcePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupMoveResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupMoveResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMoveResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMoveResourceResultOutput{})
}

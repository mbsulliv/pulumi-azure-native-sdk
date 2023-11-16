// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified static member.
func LookupStaticMember(ctx *pulumi.Context, args *LookupStaticMemberArgs, opts ...pulumi.InvokeOption) (*LookupStaticMemberResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticMemberResult
	err := ctx.Invoke("azure-native:network/v20230201:getStaticMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticMemberArgs struct {
	// The name of the network group.
	NetworkGroupName string `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the static member.
	StaticMemberName string `pulumi:"staticMemberName"`
}

// StaticMember Item.
type LookupStaticMemberResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource region.
	Region string `pulumi:"region"`
	// Resource Id.
	ResourceId *string `pulumi:"resourceId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupStaticMemberOutput(ctx *pulumi.Context, args LookupStaticMemberOutputArgs, opts ...pulumi.InvokeOption) LookupStaticMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticMemberResult, error) {
			args := v.(LookupStaticMemberArgs)
			r, err := LookupStaticMember(ctx, &args, opts...)
			var s LookupStaticMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticMemberResultOutput)
}

type LookupStaticMemberOutputArgs struct {
	// The name of the network group.
	NetworkGroupName pulumi.StringInput `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the static member.
	StaticMemberName pulumi.StringInput `pulumi:"staticMemberName"`
}

func (LookupStaticMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticMemberArgs)(nil)).Elem()
}

// StaticMember Item.
type LookupStaticMemberResultOutput struct{ *pulumi.OutputState }

func (LookupStaticMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticMemberResult)(nil)).Elem()
}

func (o LookupStaticMemberResultOutput) ToLookupStaticMemberResultOutput() LookupStaticMemberResultOutput {
	return o
}

func (o LookupStaticMemberResultOutput) ToLookupStaticMemberResultOutputWithContext(ctx context.Context) LookupStaticMemberResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupStaticMemberResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupStaticMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupStaticMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the scope assignment resource.
func (o LookupStaticMemberResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource region.
func (o LookupStaticMemberResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Region }).(pulumi.StringOutput)
}

// Resource Id.
func (o LookupStaticMemberResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o LookupStaticMemberResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupStaticMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticMemberResultOutput{})
}

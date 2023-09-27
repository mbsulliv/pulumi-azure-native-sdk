// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified network group.
func LookupNetworkGroup(ctx *pulumi.Context, args *LookupNetworkGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkGroupResult
	err := ctx.Invoke("azure-native:network/v20220401preview:getNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkGroupArgs struct {
	// The name of the network group.
	NetworkGroupName string `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The network group resource
type LookupNetworkGroupResult struct {
	// A description of the network group.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Group member type.
	MemberType string `pulumi:"memberType"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNetworkGroupOutput(ctx *pulumi.Context, args LookupNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkGroupResult, error) {
			args := v.(LookupNetworkGroupArgs)
			r, err := LookupNetworkGroup(ctx, &args, opts...)
			var s LookupNetworkGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkGroupResultOutput)
}

type LookupNetworkGroupOutputArgs struct {
	// The name of the network group.
	NetworkGroupName pulumi.StringInput `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkGroupArgs)(nil)).Elem()
}

// The network group resource
type LookupNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkGroupResult)(nil)).Elem()
}

func (o LookupNetworkGroupResultOutput) ToLookupNetworkGroupResultOutput() LookupNetworkGroupResultOutput {
	return o
}

func (o LookupNetworkGroupResultOutput) ToLookupNetworkGroupResultOutputWithContext(ctx context.Context) LookupNetworkGroupResultOutput {
	return o
}

func (o LookupNetworkGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNetworkGroupResult] {
	return pulumix.Output[LookupNetworkGroupResult]{
		OutputState: o.OutputState,
	}
}

// A description of the network group.
func (o LookupNetworkGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNetworkGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Group member type.
func (o LookupNetworkGroupResultOutput) MemberType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.MemberType }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupNetworkGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the scope assignment resource.
func (o LookupNetworkGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupNetworkGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupNetworkGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkGroupResultOutput{})
}

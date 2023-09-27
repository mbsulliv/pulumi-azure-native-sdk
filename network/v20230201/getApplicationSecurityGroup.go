// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about the specified application security group.
func LookupApplicationSecurityGroup(ctx *pulumi.Context, args *LookupApplicationSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSecurityGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationSecurityGroupResult
	err := ctx.Invoke("azure-native:network/v20230201:getApplicationSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationSecurityGroupArgs struct {
	// The name of the application security group.
	ApplicationSecurityGroupName string `pulumi:"applicationSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An application security group in a resource group.
type LookupApplicationSecurityGroupResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the application security group resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupApplicationSecurityGroupOutput(ctx *pulumi.Context, args LookupApplicationSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationSecurityGroupResult, error) {
			args := v.(LookupApplicationSecurityGroupArgs)
			r, err := LookupApplicationSecurityGroup(ctx, &args, opts...)
			var s LookupApplicationSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationSecurityGroupResultOutput)
}

type LookupApplicationSecurityGroupOutputArgs struct {
	// The name of the application security group.
	ApplicationSecurityGroupName pulumi.StringInput `pulumi:"applicationSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSecurityGroupArgs)(nil)).Elem()
}

// An application security group in a resource group.
type LookupApplicationSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSecurityGroupResult)(nil)).Elem()
}

func (o LookupApplicationSecurityGroupResultOutput) ToLookupApplicationSecurityGroupResultOutput() LookupApplicationSecurityGroupResultOutput {
	return o
}

func (o LookupApplicationSecurityGroupResultOutput) ToLookupApplicationSecurityGroupResultOutputWithContext(ctx context.Context) LookupApplicationSecurityGroupResultOutput {
	return o
}

func (o LookupApplicationSecurityGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApplicationSecurityGroupResult] {
	return pulumix.Output[LookupApplicationSecurityGroupResult]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupApplicationSecurityGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupApplicationSecurityGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupApplicationSecurityGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupApplicationSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the application security group resource.
func (o LookupApplicationSecurityGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
func (o LookupApplicationSecurityGroupResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupApplicationSecurityGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupApplicationSecurityGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationSecurityGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationSecurityGroupResultOutput{})
}

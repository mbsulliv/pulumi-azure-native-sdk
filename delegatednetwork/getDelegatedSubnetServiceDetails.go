// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package delegatednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets details about the specified dnc DelegatedSubnet Link.
// Azure REST API version: 2021-03-15.
func LookupDelegatedSubnetServiceDetails(ctx *pulumi.Context, args *LookupDelegatedSubnetServiceDetailsArgs, opts ...pulumi.InvokeOption) (*LookupDelegatedSubnetServiceDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDelegatedSubnetServiceDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork:getDelegatedSubnetServiceDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDelegatedSubnetServiceDetailsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
	ResourceName string `pulumi:"resourceName"`
}

// Represents an instance of a orchestrator.
type LookupDelegatedSubnetServiceDetailsResult struct {
	// Properties of the controller.
	ControllerDetails *ControllerDetailsResponse `pulumi:"controllerDetails"`
	// An identifier that represents the resource.
	Id string `pulumi:"id"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The current state of dnc delegated subnet resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource guid.
	ResourceGuid string `pulumi:"resourceGuid"`
	// subnet details
	SubnetDetails *SubnetDetailsResponse `pulumi:"subnetDetails"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of resource.
	Type string `pulumi:"type"`
}

func LookupDelegatedSubnetServiceDetailsOutput(ctx *pulumi.Context, args LookupDelegatedSubnetServiceDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupDelegatedSubnetServiceDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDelegatedSubnetServiceDetailsResult, error) {
			args := v.(LookupDelegatedSubnetServiceDetailsArgs)
			r, err := LookupDelegatedSubnetServiceDetails(ctx, &args, opts...)
			var s LookupDelegatedSubnetServiceDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDelegatedSubnetServiceDetailsResultOutput)
}

type LookupDelegatedSubnetServiceDetailsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupDelegatedSubnetServiceDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDelegatedSubnetServiceDetailsArgs)(nil)).Elem()
}

// Represents an instance of a orchestrator.
type LookupDelegatedSubnetServiceDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupDelegatedSubnetServiceDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDelegatedSubnetServiceDetailsResult)(nil)).Elem()
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ToLookupDelegatedSubnetServiceDetailsResultOutput() LookupDelegatedSubnetServiceDetailsResultOutput {
	return o
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ToLookupDelegatedSubnetServiceDetailsResultOutputWithContext(ctx context.Context) LookupDelegatedSubnetServiceDetailsResultOutput {
	return o
}

func (o LookupDelegatedSubnetServiceDetailsResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDelegatedSubnetServiceDetailsResult] {
	return pulumix.Output[LookupDelegatedSubnetServiceDetailsResult]{
		OutputState: o.OutputState,
	}
}

// Properties of the controller.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) ControllerDetails() ControllerDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) *ControllerDetailsResponse {
		return v.ControllerDetails
	}).(ControllerDetailsResponsePtrOutput)
}

// An identifier that represents the resource.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current state of dnc delegated subnet resource.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource guid.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// subnet details
func (o LookupDelegatedSubnetServiceDetailsResultOutput) SubnetDetails() SubnetDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) *SubnetDetailsResponse { return v.SubnetDetails }).(SubnetDetailsResponsePtrOutput)
}

// The resource tags.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of resource.
func (o LookupDelegatedSubnetServiceDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegatedSubnetServiceDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDelegatedSubnetServiceDetailsResultOutput{})
}

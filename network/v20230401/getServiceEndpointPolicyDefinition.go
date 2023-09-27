// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the specified service endpoint policy definitions from service endpoint policy.
func LookupServiceEndpointPolicyDefinition(ctx *pulumi.Context, args *LookupServiceEndpointPolicyDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointPolicyDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceEndpointPolicyDefinitionResult
	err := ctx.Invoke("azure-native:network/v20230401:getServiceEndpointPolicyDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceEndpointPolicyDefinitionArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service endpoint policy definition name.
	ServiceEndpointPolicyDefinitionName string `pulumi:"serviceEndpointPolicyDefinitionName"`
	// The name of the service endpoint policy name.
	ServiceEndpointPolicyName string `pulumi:"serviceEndpointPolicyName"`
}

// Service Endpoint policy definitions.
type LookupServiceEndpointPolicyDefinitionResult struct {
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the service endpoint policy definition resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Service endpoint name.
	Service *string `pulumi:"service"`
	// A list of service resources.
	ServiceResources []string `pulumi:"serviceResources"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

func LookupServiceEndpointPolicyDefinitionOutput(ctx *pulumi.Context, args LookupServiceEndpointPolicyDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupServiceEndpointPolicyDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceEndpointPolicyDefinitionResult, error) {
			args := v.(LookupServiceEndpointPolicyDefinitionArgs)
			r, err := LookupServiceEndpointPolicyDefinition(ctx, &args, opts...)
			var s LookupServiceEndpointPolicyDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceEndpointPolicyDefinitionResultOutput)
}

type LookupServiceEndpointPolicyDefinitionOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service endpoint policy definition name.
	ServiceEndpointPolicyDefinitionName pulumi.StringInput `pulumi:"serviceEndpointPolicyDefinitionName"`
	// The name of the service endpoint policy name.
	ServiceEndpointPolicyName pulumi.StringInput `pulumi:"serviceEndpointPolicyName"`
}

func (LookupServiceEndpointPolicyDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyDefinitionArgs)(nil)).Elem()
}

// Service Endpoint policy definitions.
type LookupServiceEndpointPolicyDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupServiceEndpointPolicyDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyDefinitionResult)(nil)).Elem()
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) ToLookupServiceEndpointPolicyDefinitionResultOutput() LookupServiceEndpointPolicyDefinitionResultOutput {
	return o
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) ToLookupServiceEndpointPolicyDefinitionResultOutputWithContext(ctx context.Context) LookupServiceEndpointPolicyDefinitionResultOutput {
	return o
}

func (o LookupServiceEndpointPolicyDefinitionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServiceEndpointPolicyDefinitionResult] {
	return pulumix.Output[LookupServiceEndpointPolicyDefinitionResult]{
		OutputState: o.OutputState,
	}
}

// A description for this rule. Restricted to 140 chars.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the service endpoint policy definition resource.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Service endpoint name.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) Service() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Service }).(pulumi.StringPtrOutput)
}

// A list of service resources.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) ServiceResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) []string { return v.ServiceResources }).(pulumi.StringArrayOutput)
}

// The type of the resource.
func (o LookupServiceEndpointPolicyDefinitionResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyDefinitionResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceEndpointPolicyDefinitionResultOutput{})
}

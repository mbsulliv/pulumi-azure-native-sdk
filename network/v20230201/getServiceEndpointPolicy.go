// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified service Endpoint Policies in a specified resource group.
func LookupServiceEndpointPolicy(ctx *pulumi.Context, args *LookupServiceEndpointPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServiceEndpointPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceEndpointPolicyResult
	err := ctx.Invoke("azure-native:network/v20230201:getServiceEndpointPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceEndpointPolicyArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName string `pulumi:"serviceEndpointPolicyName"`
}

// Service End point policy resource.
type LookupServiceEndpointPolicyResult struct {
	// A collection of contextual service endpoint policy.
	ContextualServiceEndpointPolicies []string `pulumi:"contextualServiceEndpointPolicies"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Kind of service endpoint policy. This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the service endpoint policy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the service endpoint policy resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The alias indicating if the policy belongs to a service
	ServiceAlias *string `pulumi:"serviceAlias"`
	// A collection of service endpoint policy definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionResponse `pulumi:"serviceEndpointPolicyDefinitions"`
	// A collection of references to subnets.
	Subnets []SubnetResponse `pulumi:"subnets"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupServiceEndpointPolicyOutput(ctx *pulumi.Context, args LookupServiceEndpointPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServiceEndpointPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceEndpointPolicyResult, error) {
			args := v.(LookupServiceEndpointPolicyArgs)
			r, err := LookupServiceEndpointPolicy(ctx, &args, opts...)
			var s LookupServiceEndpointPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceEndpointPolicyResultOutput)
}

type LookupServiceEndpointPolicyOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName pulumi.StringInput `pulumi:"serviceEndpointPolicyName"`
}

func (LookupServiceEndpointPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyArgs)(nil)).Elem()
}

// Service End point policy resource.
type LookupServiceEndpointPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServiceEndpointPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceEndpointPolicyResult)(nil)).Elem()
}

func (o LookupServiceEndpointPolicyResultOutput) ToLookupServiceEndpointPolicyResultOutput() LookupServiceEndpointPolicyResultOutput {
	return o
}

func (o LookupServiceEndpointPolicyResultOutput) ToLookupServiceEndpointPolicyResultOutputWithContext(ctx context.Context) LookupServiceEndpointPolicyResultOutput {
	return o
}

// A collection of contextual service endpoint policy.
func (o LookupServiceEndpointPolicyResultOutput) ContextualServiceEndpointPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) []string { return v.ContextualServiceEndpointPolicies }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupServiceEndpointPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupServiceEndpointPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Kind of service endpoint policy. This is metadata used for the Azure portal experience.
func (o LookupServiceEndpointPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupServiceEndpointPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupServiceEndpointPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the service endpoint policy resource.
func (o LookupServiceEndpointPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the service endpoint policy resource.
func (o LookupServiceEndpointPolicyResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The alias indicating if the policy belongs to a service
func (o LookupServiceEndpointPolicyResultOutput) ServiceAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) *string { return v.ServiceAlias }).(pulumi.StringPtrOutput)
}

// A collection of service endpoint policy definitions of the service endpoint policy.
func (o LookupServiceEndpointPolicyResultOutput) ServiceEndpointPolicyDefinitions() ServiceEndpointPolicyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) []ServiceEndpointPolicyDefinitionResponse {
		return v.ServiceEndpointPolicyDefinitions
	}).(ServiceEndpointPolicyDefinitionResponseArrayOutput)
}

// A collection of references to subnets.
func (o LookupServiceEndpointPolicyResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o LookupServiceEndpointPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupServiceEndpointPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceEndpointPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceEndpointPolicyResultOutput{})
}
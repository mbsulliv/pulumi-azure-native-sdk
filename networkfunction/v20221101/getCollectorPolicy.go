// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the collector policy in a specified Traffic Collector
func LookupCollectorPolicy(ctx *pulumi.Context, args *LookupCollectorPolicyArgs, opts ...pulumi.InvokeOption) (*LookupCollectorPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCollectorPolicyResult
	err := ctx.Invoke("azure-native:networkfunction/v20221101:getCollectorPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCollectorPolicyArgs struct {
	// Azure Traffic Collector name
	AzureTrafficCollectorName string `pulumi:"azureTrafficCollectorName"`
	// Collector Policy Name
	CollectorPolicyName string `pulumi:"collectorPolicyName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Collector policy resource.
type LookupCollectorPolicyResult struct {
	// Emission policies.
	EmissionPolicies []EmissionPoliciesPropertiesFormatResponse `pulumi:"emissionPolicies"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Ingestion policies.
	IngestionPolicy *IngestionPolicyPropertiesFormatResponse `pulumi:"ingestionPolicy"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData TrackedResourceResponseSystemData `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupCollectorPolicyOutput(ctx *pulumi.Context, args LookupCollectorPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupCollectorPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCollectorPolicyResult, error) {
			args := v.(LookupCollectorPolicyArgs)
			r, err := LookupCollectorPolicy(ctx, &args, opts...)
			var s LookupCollectorPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCollectorPolicyResultOutput)
}

type LookupCollectorPolicyOutputArgs struct {
	// Azure Traffic Collector name
	AzureTrafficCollectorName pulumi.StringInput `pulumi:"azureTrafficCollectorName"`
	// Collector Policy Name
	CollectorPolicyName pulumi.StringInput `pulumi:"collectorPolicyName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCollectorPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectorPolicyArgs)(nil)).Elem()
}

// Collector policy resource.
type LookupCollectorPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupCollectorPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollectorPolicyResult)(nil)).Elem()
}

func (o LookupCollectorPolicyResultOutput) ToLookupCollectorPolicyResultOutput() LookupCollectorPolicyResultOutput {
	return o
}

func (o LookupCollectorPolicyResultOutput) ToLookupCollectorPolicyResultOutputWithContext(ctx context.Context) LookupCollectorPolicyResultOutput {
	return o
}

// Emission policies.
func (o LookupCollectorPolicyResultOutput) EmissionPolicies() EmissionPoliciesPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) []EmissionPoliciesPropertiesFormatResponse {
		return v.EmissionPolicies
	}).(EmissionPoliciesPropertiesFormatResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupCollectorPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupCollectorPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Ingestion policies.
func (o LookupCollectorPolicyResultOutput) IngestionPolicy() IngestionPolicyPropertiesFormatResponsePtrOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) *IngestionPolicyPropertiesFormatResponse { return v.IngestionPolicy }).(IngestionPolicyPropertiesFormatResponsePtrOutput)
}

// Resource location.
func (o LookupCollectorPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupCollectorPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state.
func (o LookupCollectorPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupCollectorPolicyResultOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) TrackedResourceResponseSystemData { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

// Resource tags.
func (o LookupCollectorPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupCollectorPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCollectorPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCollectorPolicyResultOutput{})
}

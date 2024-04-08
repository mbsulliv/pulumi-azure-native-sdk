// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of generic ARM proxy resource.
// Azure REST API version: 2022-06-01.
//
// Other available API versions: 2023-03-11.
func LookupDataCollectionRuleAssociation(ctx *pulumi.Context, args *LookupDataCollectionRuleAssociationArgs, opts ...pulumi.InvokeOption) (*LookupDataCollectionRuleAssociationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataCollectionRuleAssociationResult
	err := ctx.Invoke("azure-native:insights:getDataCollectionRuleAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataCollectionRuleAssociationArgs struct {
	// The name of the association. The name is case insensitive.
	AssociationName string `pulumi:"associationName"`
	// The identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// Definition of generic ARM proxy resource.
type LookupDataCollectionRuleAssociationResult struct {
	// The resource ID of the data collection endpoint that is to be associated.
	DataCollectionEndpointId *string `pulumi:"dataCollectionEndpointId"`
	// The resource ID of the data collection rule that is to be associated.
	DataCollectionRuleId *string `pulumi:"dataCollectionRuleId"`
	// Description of the association.
	Description *string `pulumi:"description"`
	// Resource entity tag (ETag).
	Etag string `pulumi:"etag"`
	// Fully qualified ID of the resource.
	Id string `pulumi:"id"`
	// Metadata about the resource
	Metadata DataCollectionRuleAssociationResponseMetadata `pulumi:"metadata"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The resource provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupDataCollectionRuleAssociationOutput(ctx *pulumi.Context, args LookupDataCollectionRuleAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupDataCollectionRuleAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataCollectionRuleAssociationResult, error) {
			args := v.(LookupDataCollectionRuleAssociationArgs)
			r, err := LookupDataCollectionRuleAssociation(ctx, &args, opts...)
			var s LookupDataCollectionRuleAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataCollectionRuleAssociationResultOutput)
}

type LookupDataCollectionRuleAssociationOutputArgs struct {
	// The name of the association. The name is case insensitive.
	AssociationName pulumi.StringInput `pulumi:"associationName"`
	// The identifier of the resource.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupDataCollectionRuleAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleAssociationArgs)(nil)).Elem()
}

// Definition of generic ARM proxy resource.
type LookupDataCollectionRuleAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupDataCollectionRuleAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCollectionRuleAssociationResult)(nil)).Elem()
}

func (o LookupDataCollectionRuleAssociationResultOutput) ToLookupDataCollectionRuleAssociationResultOutput() LookupDataCollectionRuleAssociationResultOutput {
	return o
}

func (o LookupDataCollectionRuleAssociationResultOutput) ToLookupDataCollectionRuleAssociationResultOutputWithContext(ctx context.Context) LookupDataCollectionRuleAssociationResultOutput {
	return o
}

// The resource ID of the data collection endpoint that is to be associated.
func (o LookupDataCollectionRuleAssociationResultOutput) DataCollectionEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) *string { return v.DataCollectionEndpointId }).(pulumi.StringPtrOutput)
}

// The resource ID of the data collection rule that is to be associated.
func (o LookupDataCollectionRuleAssociationResultOutput) DataCollectionRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) *string { return v.DataCollectionRuleId }).(pulumi.StringPtrOutput)
}

// Description of the association.
func (o LookupDataCollectionRuleAssociationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource entity tag (ETag).
func (o LookupDataCollectionRuleAssociationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified ID of the resource.
func (o LookupDataCollectionRuleAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Metadata about the resource
func (o LookupDataCollectionRuleAssociationResultOutput) Metadata() DataCollectionRuleAssociationResponseMetadataOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) DataCollectionRuleAssociationResponseMetadata {
		return v.Metadata
	}).(DataCollectionRuleAssociationResponseMetadataOutput)
}

// The name of the resource.
func (o LookupDataCollectionRuleAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource provisioning state.
func (o LookupDataCollectionRuleAssociationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDataCollectionRuleAssociationResultOutput) SystemData() DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) DataCollectionRuleAssociationProxyOnlyResourceResponseSystemData {
		return v.SystemData
	}).(DataCollectionRuleAssociationProxyOnlyResourceResponseSystemDataOutput)
}

// The type of the resource.
func (o LookupDataCollectionRuleAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCollectionRuleAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataCollectionRuleAssociationResultOutput{})
}

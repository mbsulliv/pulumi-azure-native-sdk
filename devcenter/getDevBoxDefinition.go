// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Dev Box definition
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2022-11-11-preview, 2023-08-01-preview, 2023-10-01-preview, 2024-02-01.
func LookupDevBoxDefinition(ctx *pulumi.Context, args *LookupDevBoxDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupDevBoxDefinitionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDevBoxDefinitionResult
	err := ctx.Invoke("azure-native:devcenter:getDevBoxDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDevBoxDefinitionArgs struct {
	// The name of the Dev Box definition.
	DevBoxDefinitionName string `pulumi:"devBoxDefinitionName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a definition for a Developer Machine.
type LookupDevBoxDefinitionResult struct {
	// Image reference information for the currently active image (only populated during updates).
	ActiveImageReference ImageReferenceResponse `pulumi:"activeImageReference"`
	// Indicates whether Dev Boxes created with this definition are capable of hibernation. Not all images are capable of supporting hibernation. To find out more see https://aka.ms/devbox/hibernate
	HibernateSupport *string `pulumi:"hibernateSupport"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Image reference information.
	ImageReference ImageReferenceResponse `pulumi:"imageReference"`
	// Details for image validator error. Populated when the image validation is not successful.
	ImageValidationErrorDetails ImageValidationErrorDetailsResponse `pulumi:"imageValidationErrorDetails"`
	// Validation status of the configured image.
	ImageValidationStatus string `pulumi:"imageValidationStatus"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The storage type used for the Operating System disk of Dev Boxes created using this definition.
	OsStorageType *string `pulumi:"osStorageType"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU for Dev Boxes created using this definition.
	Sku SkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDevBoxDefinitionOutput(ctx *pulumi.Context, args LookupDevBoxDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupDevBoxDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDevBoxDefinitionResult, error) {
			args := v.(LookupDevBoxDefinitionArgs)
			r, err := LookupDevBoxDefinition(ctx, &args, opts...)
			var s LookupDevBoxDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDevBoxDefinitionResultOutput)
}

type LookupDevBoxDefinitionOutputArgs struct {
	// The name of the Dev Box definition.
	DevBoxDefinitionName pulumi.StringInput `pulumi:"devBoxDefinitionName"`
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDevBoxDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevBoxDefinitionArgs)(nil)).Elem()
}

// Represents a definition for a Developer Machine.
type LookupDevBoxDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupDevBoxDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDevBoxDefinitionResult)(nil)).Elem()
}

func (o LookupDevBoxDefinitionResultOutput) ToLookupDevBoxDefinitionResultOutput() LookupDevBoxDefinitionResultOutput {
	return o
}

func (o LookupDevBoxDefinitionResultOutput) ToLookupDevBoxDefinitionResultOutputWithContext(ctx context.Context) LookupDevBoxDefinitionResultOutput {
	return o
}

// Image reference information for the currently active image (only populated during updates).
func (o LookupDevBoxDefinitionResultOutput) ActiveImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) ImageReferenceResponse { return v.ActiveImageReference }).(ImageReferenceResponseOutput)
}

// Indicates whether Dev Boxes created with this definition are capable of hibernation. Not all images are capable of supporting hibernation. To find out more see https://aka.ms/devbox/hibernate
func (o LookupDevBoxDefinitionResultOutput) HibernateSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) *string { return v.HibernateSupport }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDevBoxDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Image reference information.
func (o LookupDevBoxDefinitionResultOutput) ImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponseOutput)
}

// Details for image validator error. Populated when the image validation is not successful.
func (o LookupDevBoxDefinitionResultOutput) ImageValidationErrorDetails() ImageValidationErrorDetailsResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) ImageValidationErrorDetailsResponse {
		return v.ImageValidationErrorDetails
	}).(ImageValidationErrorDetailsResponseOutput)
}

// Validation status of the configured image.
func (o LookupDevBoxDefinitionResultOutput) ImageValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.ImageValidationStatus }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDevBoxDefinitionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDevBoxDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The storage type used for the Operating System disk of Dev Boxes created using this definition.
func (o LookupDevBoxDefinitionResultOutput) OsStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) *string { return v.OsStorageType }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o LookupDevBoxDefinitionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU for Dev Boxes created using this definition.
func (o LookupDevBoxDefinitionResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDevBoxDefinitionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDevBoxDefinitionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDevBoxDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDevBoxDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDevBoxDefinitionResultOutput{})
}

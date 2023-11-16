// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the user ManagementAssociation.
func LookupManagementAssociation(ctx *pulumi.Context, args *LookupManagementAssociationArgs, opts ...pulumi.InvokeOption) (*LookupManagementAssociationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagementAssociationResult
	err := ctx.Invoke("azure-native:operationsmanagement/v20151101preview:getManagementAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementAssociationArgs struct {
	// User ManagementAssociation Name.
	ManagementAssociationName string `pulumi:"managementAssociationName"`
	// Provider name for the parent resource.
	ProviderName string `pulumi:"providerName"`
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parent resource name.
	ResourceName string `pulumi:"resourceName"`
	// Resource type for the parent resource
	ResourceType string `pulumi:"resourceType"`
}

// The container for solution.
type LookupManagementAssociationResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Properties for ManagementAssociation object supported by the OperationsManagement resource provider.
	Properties ManagementAssociationPropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupManagementAssociationOutput(ctx *pulumi.Context, args LookupManagementAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupManagementAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementAssociationResult, error) {
			args := v.(LookupManagementAssociationArgs)
			r, err := LookupManagementAssociation(ctx, &args, opts...)
			var s LookupManagementAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementAssociationResultOutput)
}

type LookupManagementAssociationOutputArgs struct {
	// User ManagementAssociation Name.
	ManagementAssociationName pulumi.StringInput `pulumi:"managementAssociationName"`
	// Provider name for the parent resource.
	ProviderName pulumi.StringInput `pulumi:"providerName"`
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Parent resource name.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// Resource type for the parent resource
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupManagementAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementAssociationArgs)(nil)).Elem()
}

// The container for solution.
type LookupManagementAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupManagementAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementAssociationResult)(nil)).Elem()
}

func (o LookupManagementAssociationResultOutput) ToLookupManagementAssociationResultOutput() LookupManagementAssociationResultOutput {
	return o
}

func (o LookupManagementAssociationResultOutput) ToLookupManagementAssociationResultOutputWithContext(ctx context.Context) LookupManagementAssociationResultOutput {
	return o
}

// Resource ID.
func (o LookupManagementAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupManagementAssociationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupManagementAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties for ManagementAssociation object supported by the OperationsManagement resource provider.
func (o LookupManagementAssociationResultOutput) Properties() ManagementAssociationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) ManagementAssociationPropertiesResponse { return v.Properties }).(ManagementAssociationPropertiesResponseOutput)
}

// Resource type.
func (o LookupManagementAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementAssociationResultOutput{})
}

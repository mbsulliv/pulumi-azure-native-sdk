// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of a partner configuration.
func LookupPartnerConfiguration(ctx *pulumi.Context, args *LookupPartnerConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPartnerConfigurationResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getPartnerConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerConfigurationArgs struct {
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Partner configuration information
type LookupPartnerConfigurationResult struct {
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// The details of authorized partners.
	PartnerAuthorization *PartnerAuthorizationResponse `pulumi:"partnerAuthorization"`
	// Provisioning state of the partner configuration.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The system metadata relating to partner configuration resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupPartnerConfigurationOutput(ctx *pulumi.Context, args LookupPartnerConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerConfigurationResult, error) {
			args := v.(LookupPartnerConfigurationArgs)
			r, err := LookupPartnerConfiguration(ctx, &args, opts...)
			var s LookupPartnerConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerConfigurationResultOutput)
}

type LookupPartnerConfigurationOutputArgs struct {
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerConfigurationArgs)(nil)).Elem()
}

// Partner configuration information
type LookupPartnerConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerConfigurationResult)(nil)).Elem()
}

func (o LookupPartnerConfigurationResultOutput) ToLookupPartnerConfigurationResultOutput() LookupPartnerConfigurationResultOutput {
	return o
}

func (o LookupPartnerConfigurationResultOutput) ToLookupPartnerConfigurationResultOutputWithContext(ctx context.Context) LookupPartnerConfigurationResultOutput {
	return o
}

// Fully qualified identifier of the resource.
func (o LookupPartnerConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource.
func (o LookupPartnerConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupPartnerConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The details of authorized partners.
func (o LookupPartnerConfigurationResultOutput) PartnerAuthorization() PartnerAuthorizationResponsePtrOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) *PartnerAuthorizationResponse { return v.PartnerAuthorization }).(PartnerAuthorizationResponsePtrOutput)
}

// Provisioning state of the partner configuration.
func (o LookupPartnerConfigurationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The system metadata relating to partner configuration resource.
func (o LookupPartnerConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o LookupPartnerConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o LookupPartnerConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerConfigurationResultOutput{})
}

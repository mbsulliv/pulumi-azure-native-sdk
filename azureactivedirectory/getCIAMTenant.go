// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azureactivedirectory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the Azure AD for customers tenant resource.
// Azure REST API version: 2023-05-17-preview.
func LookupCIAMTenant(ctx *pulumi.Context, args *LookupCIAMTenantArgs, opts ...pulumi.InvokeOption) (*LookupCIAMTenantResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCIAMTenantResult
	err := ctx.Invoke("azure-native:azureactivedirectory:getCIAMTenant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCIAMTenantArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The initial sub domain of the tenant.
	ResourceName string `pulumi:"resourceName"`
}

// The Azure AD for customers resource.
type LookupCIAMTenantResult struct {
	// The type of billing. Will be MAU for all new customers. Cannot be changed if value is 'MAU'. Learn more about Azure AD for customers billing at [aka.ms/b2cBilling](https://aka.ms/b2cbilling).
	BillingType string `pulumi:"billingType"`
	// These properties are used to create the Azure AD for customers tenant. These properties are not part of the Azure resource.
	CreateTenantProperties CreateCIAMTenantPropertiesResponse `pulumi:"createTenantProperties"`
	// The domain name of the tenant
	DomainName string `pulumi:"domainName"`
	// The data from which the billing type took effect
	EffectiveStartDateUtc string `pulumi:"effectiveStartDateUtc"`
	// An identifier that represents the Azure AD for customers tenant resource.
	Id string `pulumi:"id"`
	// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/ciam-data-location) for more information.
	Location string `pulumi:"location"`
	// The name of the Azure AD for customers tenant resource.
	Name              string `pulumi:"name"`
	ProvisioningState string `pulumi:"provisioningState"`
	// SKU properties of the Azure AD for customers tenant. Learn more about Azure AD for customers billing at [https://aka.ms/ciambilling](https://aka.ms/ciambilling).
	Sku CIAMResourceSKUResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
	// An identifier of the Azure AD for customers tenant.
	TenantId *string `pulumi:"tenantId"`
	// The type of the Azure AD for customers tenant resource.
	Type string `pulumi:"type"`
}

func LookupCIAMTenantOutput(ctx *pulumi.Context, args LookupCIAMTenantOutputArgs, opts ...pulumi.InvokeOption) LookupCIAMTenantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCIAMTenantResult, error) {
			args := v.(LookupCIAMTenantArgs)
			r, err := LookupCIAMTenant(ctx, &args, opts...)
			var s LookupCIAMTenantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCIAMTenantResultOutput)
}

type LookupCIAMTenantOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The initial sub domain of the tenant.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupCIAMTenantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCIAMTenantArgs)(nil)).Elem()
}

// The Azure AD for customers resource.
type LookupCIAMTenantResultOutput struct{ *pulumi.OutputState }

func (LookupCIAMTenantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCIAMTenantResult)(nil)).Elem()
}

func (o LookupCIAMTenantResultOutput) ToLookupCIAMTenantResultOutput() LookupCIAMTenantResultOutput {
	return o
}

func (o LookupCIAMTenantResultOutput) ToLookupCIAMTenantResultOutputWithContext(ctx context.Context) LookupCIAMTenantResultOutput {
	return o
}

// The type of billing. Will be MAU for all new customers. Cannot be changed if value is 'MAU'. Learn more about Azure AD for customers billing at [aka.ms/b2cBilling](https://aka.ms/b2cbilling).
func (o LookupCIAMTenantResultOutput) BillingType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.BillingType }).(pulumi.StringOutput)
}

// These properties are used to create the Azure AD for customers tenant. These properties are not part of the Azure resource.
func (o LookupCIAMTenantResultOutput) CreateTenantProperties() CreateCIAMTenantPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) CreateCIAMTenantPropertiesResponse { return v.CreateTenantProperties }).(CreateCIAMTenantPropertiesResponseOutput)
}

// The domain name of the tenant
func (o LookupCIAMTenantResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// The data from which the billing type took effect
func (o LookupCIAMTenantResultOutput) EffectiveStartDateUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.EffectiveStartDateUtc }).(pulumi.StringOutput)
}

// An identifier that represents the Azure AD for customers tenant resource.
func (o LookupCIAMTenantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/ciam-data-location) for more information.
func (o LookupCIAMTenantResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the Azure AD for customers tenant resource.
func (o LookupCIAMTenantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCIAMTenantResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SKU properties of the Azure AD for customers tenant. Learn more about Azure AD for customers billing at [https://aka.ms/ciambilling](https://aka.ms/ciambilling).
func (o LookupCIAMTenantResultOutput) Sku() CIAMResourceSKUResponseOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) CIAMResourceSKUResponse { return v.Sku }).(CIAMResourceSKUResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupCIAMTenantResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource Tags
func (o LookupCIAMTenantResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// An identifier of the Azure AD for customers tenant.
func (o LookupCIAMTenantResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the Azure AD for customers tenant resource.
func (o LookupCIAMTenantResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCIAMTenantResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCIAMTenantResultOutput{})
}
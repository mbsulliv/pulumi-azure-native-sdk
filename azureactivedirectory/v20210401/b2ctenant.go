// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type B2CTenant struct {
	pulumi.CustomResourceState

	// The billing configuration for the tenant.
	BillingConfig B2CTenantResourcePropertiesResponseBillingConfigPtrOutput `pulumi:"billingConfig"`
	// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/B2CDataResidency) for more information.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Azure AD B2C tenant resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
	Sku B2CResourceSKUResponseOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource Tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// An identifier of the Azure AD B2C tenant.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the B2C tenant resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewB2CTenant registers a new resource with the given unique name, arguments, and options.
func NewB2CTenant(ctx *pulumi.Context,
	name string, args *B2CTenantArgs, opts ...pulumi.ResourceOption) (*B2CTenant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azureactivedirectory:B2CTenant"),
		},
		{
			Type: pulumi.String("azure-native:azureactivedirectory/v20190101preview:B2CTenant"),
		},
		{
			Type: pulumi.String("azure-native:azureactivedirectory/v20230118preview:B2CTenant"),
		},
		{
			Type: pulumi.String("azure-native:azureactivedirectory/v20230517preview:B2CTenant"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource B2CTenant
	err := ctx.RegisterResource("azure-native:azureactivedirectory/v20210401:B2CTenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetB2CTenant gets an existing B2CTenant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetB2CTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *B2CTenantState, opts ...pulumi.ResourceOption) (*B2CTenant, error) {
	var resource B2CTenant
	err := ctx.ReadResource("azure-native:azureactivedirectory/v20210401:B2CTenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering B2CTenant resources.
type b2ctenantState struct {
}

type B2CTenantState struct {
}

func (B2CTenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*b2ctenantState)(nil)).Elem()
}

type b2ctenantArgs struct {
	// Country code of Azure tenant (e.g. 'US'). Refer to [aka.ms/B2CDataResidency](https://aka.ms/B2CDataResidency) to see valid country codes and corresponding data residency locations. If you do not see a country code in an valid data residency location, choose one from the list.
	CountryCode *string `pulumi:"countryCode"`
	// The display name of the Azure AD B2C tenant.
	DisplayName *string `pulumi:"displayName"`
	// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/B2CDataResidency) for more information.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The initial domain name of the Azure AD B2C tenant.
	ResourceName *string `pulumi:"resourceName"`
	// SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
	Sku B2CResourceSKU `pulumi:"sku"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a B2CTenant resource.
type B2CTenantArgs struct {
	// Country code of Azure tenant (e.g. 'US'). Refer to [aka.ms/B2CDataResidency](https://aka.ms/B2CDataResidency) to see valid country codes and corresponding data residency locations. If you do not see a country code in an valid data residency location, choose one from the list.
	CountryCode pulumi.StringPtrInput
	// The display name of the Azure AD B2C tenant.
	DisplayName pulumi.StringPtrInput
	// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/B2CDataResidency) for more information.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The initial domain name of the Azure AD B2C tenant.
	ResourceName pulumi.StringPtrInput
	// SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
	Sku B2CResourceSKUInput
	// Resource Tags
	Tags pulumi.StringMapInput
}

func (B2CTenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*b2ctenantArgs)(nil)).Elem()
}

type B2CTenantInput interface {
	pulumi.Input

	ToB2CTenantOutput() B2CTenantOutput
	ToB2CTenantOutputWithContext(ctx context.Context) B2CTenantOutput
}

func (*B2CTenant) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CTenant)(nil)).Elem()
}

func (i *B2CTenant) ToB2CTenantOutput() B2CTenantOutput {
	return i.ToB2CTenantOutputWithContext(context.Background())
}

func (i *B2CTenant) ToB2CTenantOutputWithContext(ctx context.Context) B2CTenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CTenantOutput)
}

type B2CTenantOutput struct{ *pulumi.OutputState }

func (B2CTenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CTenant)(nil)).Elem()
}

func (o B2CTenantOutput) ToB2CTenantOutput() B2CTenantOutput {
	return o
}

func (o B2CTenantOutput) ToB2CTenantOutputWithContext(ctx context.Context) B2CTenantOutput {
	return o
}

// The billing configuration for the tenant.
func (o B2CTenantOutput) BillingConfig() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return o.ApplyT(func(v *B2CTenant) B2CTenantResourcePropertiesResponseBillingConfigPtrOutput { return v.BillingConfig }).(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput)
}

// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/B2CDataResidency) for more information.
func (o B2CTenantOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the Azure AD B2C tenant resource.
func (o B2CTenantOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SKU properties of the Azure AD B2C tenant. Learn more about Azure AD B2C billing at [aka.ms/b2cBilling](https://aka.ms/b2cBilling).
func (o B2CTenantOutput) Sku() B2CResourceSKUResponseOutput {
	return o.ApplyT(func(v *B2CTenant) B2CResourceSKUResponseOutput { return v.Sku }).(B2CResourceSKUResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o B2CTenantOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *B2CTenant) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource Tags
func (o B2CTenantOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// An identifier of the Azure AD B2C tenant.
func (o B2CTenantOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the B2C tenant resource.
func (o B2CTenantOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(B2CTenantOutput{})
}

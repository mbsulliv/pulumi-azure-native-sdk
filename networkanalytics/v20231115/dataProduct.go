// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231115

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The data product resource.
type DataProduct struct {
	pulumi.CustomResourceState

	// List of available minor versions of the data product resource.
	AvailableMinorVersions pulumi.StringArrayOutput `pulumi:"availableMinorVersions"`
	// Resource links which exposed to the customer to query the data.
	ConsumptionEndpoints ConsumptionEndpointsPropertiesResponseOutput `pulumi:"consumptionEndpoints"`
	// Current configured minor version of the data product resource.
	CurrentMinorVersion pulumi.StringPtrOutput `pulumi:"currentMinorVersion"`
	// Customer managed encryption key details for data product.
	CustomerEncryptionKey EncryptionKeyDetailsResponsePtrOutput `pulumi:"customerEncryptionKey"`
	// Flag to enable customer managed key encryption for data product.
	CustomerManagedKeyEncryptionEnabled pulumi.StringPtrOutput `pulumi:"customerManagedKeyEncryptionEnabled"`
	// Documentation link for the data product based on definition file.
	Documentation pulumi.StringOutput `pulumi:"documentation"`
	// The managed service identities assigned to this resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Key vault url.
	KeyVaultUrl pulumi.StringOutput `pulumi:"keyVaultUrl"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Major version of data product.
	MajorVersion pulumi.StringOutput `pulumi:"majorVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network rule set for data product.
	Networkacls DataProductNetworkAclsResponsePtrOutput `pulumi:"networkacls"`
	// List of name or email associated with data product resource deployment.
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// Flag to enable or disable private link for data product resource.
	PrivateLinksEnabled pulumi.StringPtrOutput `pulumi:"privateLinksEnabled"`
	// Product name of data product.
	Product pulumi.StringOutput `pulumi:"product"`
	// Latest provisioning state  of data product.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Flag to enable or disable public access of data product resource.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Data product publisher name.
	Publisher pulumi.StringOutput `pulumi:"publisher"`
	// Purview account url for data product to connect to.
	PurviewAccount pulumi.StringPtrOutput `pulumi:"purviewAccount"`
	// Purview collection url for data product to connect to.
	PurviewCollection pulumi.StringPtrOutput `pulumi:"purviewCollection"`
	// Flag to enable or disable redundancy for data product.
	Redundancy pulumi.StringPtrOutput `pulumi:"redundancy"`
	// The resource GUID property of the data product resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataProduct registers a new resource with the given unique name, arguments, and options.
func NewDataProduct(ctx *pulumi.Context,
	name string, args *DataProductArgs, opts ...pulumi.ResourceOption) (*DataProduct, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MajorVersion == nil {
		return nil, errors.New("invalid value for required argument 'MajorVersion'")
	}
	if args.Product == nil {
		return nil, errors.New("invalid value for required argument 'Product'")
	}
	if args.Publisher == nil {
		return nil, errors.New("invalid value for required argument 'Publisher'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkanalytics:DataProduct"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataProduct
	err := ctx.RegisterResource("azure-native:networkanalytics/v20231115:DataProduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataProduct gets an existing DataProduct resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataProductState, opts ...pulumi.ResourceOption) (*DataProduct, error) {
	var resource DataProduct
	err := ctx.ReadResource("azure-native:networkanalytics/v20231115:DataProduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataProduct resources.
type dataProductState struct {
}

type DataProductState struct {
}

func (DataProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataProductState)(nil)).Elem()
}

type dataProductArgs struct {
	// Current configured minor version of the data product resource.
	CurrentMinorVersion *string `pulumi:"currentMinorVersion"`
	// Customer managed encryption key details for data product.
	CustomerEncryptionKey *EncryptionKeyDetails `pulumi:"customerEncryptionKey"`
	// Flag to enable customer managed key encryption for data product.
	CustomerManagedKeyEncryptionEnabled *string `pulumi:"customerManagedKeyEncryptionEnabled"`
	// The data product resource name
	DataProductName *string `pulumi:"dataProductName"`
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Major version of data product.
	MajorVersion string `pulumi:"majorVersion"`
	// Managed resource group configuration.
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration `pulumi:"managedResourceGroupConfiguration"`
	// Network rule set for data product.
	Networkacls *DataProductNetworkAcls `pulumi:"networkacls"`
	// List of name or email associated with data product resource deployment.
	Owners []string `pulumi:"owners"`
	// Flag to enable or disable private link for data product resource.
	PrivateLinksEnabled *string `pulumi:"privateLinksEnabled"`
	// Product name of data product.
	Product string `pulumi:"product"`
	// Flag to enable or disable public access of data product resource.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Data product publisher name.
	Publisher string `pulumi:"publisher"`
	// Purview account url for data product to connect to.
	PurviewAccount *string `pulumi:"purviewAccount"`
	// Purview collection url for data product to connect to.
	PurviewCollection *string `pulumi:"purviewCollection"`
	// Flag to enable or disable redundancy for data product.
	Redundancy *string `pulumi:"redundancy"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DataProduct resource.
type DataProductArgs struct {
	// Current configured minor version of the data product resource.
	CurrentMinorVersion pulumi.StringPtrInput
	// Customer managed encryption key details for data product.
	CustomerEncryptionKey EncryptionKeyDetailsPtrInput
	// Flag to enable customer managed key encryption for data product.
	CustomerManagedKeyEncryptionEnabled pulumi.StringPtrInput
	// The data product resource name
	DataProductName pulumi.StringPtrInput
	// The managed service identities assigned to this resource.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Major version of data product.
	MajorVersion pulumi.StringInput
	// Managed resource group configuration.
	ManagedResourceGroupConfiguration ManagedResourceGroupConfigurationPtrInput
	// Network rule set for data product.
	Networkacls DataProductNetworkAclsPtrInput
	// List of name or email associated with data product resource deployment.
	Owners pulumi.StringArrayInput
	// Flag to enable or disable private link for data product resource.
	PrivateLinksEnabled pulumi.StringPtrInput
	// Product name of data product.
	Product pulumi.StringInput
	// Flag to enable or disable public access of data product resource.
	PublicNetworkAccess pulumi.StringPtrInput
	// Data product publisher name.
	Publisher pulumi.StringInput
	// Purview account url for data product to connect to.
	PurviewAccount pulumi.StringPtrInput
	// Purview collection url for data product to connect to.
	PurviewCollection pulumi.StringPtrInput
	// Flag to enable or disable redundancy for data product.
	Redundancy pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DataProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataProductArgs)(nil)).Elem()
}

type DataProductInput interface {
	pulumi.Input

	ToDataProductOutput() DataProductOutput
	ToDataProductOutputWithContext(ctx context.Context) DataProductOutput
}

func (*DataProduct) ElementType() reflect.Type {
	return reflect.TypeOf((**DataProduct)(nil)).Elem()
}

func (i *DataProduct) ToDataProductOutput() DataProductOutput {
	return i.ToDataProductOutputWithContext(context.Background())
}

func (i *DataProduct) ToDataProductOutputWithContext(ctx context.Context) DataProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataProductOutput)
}

type DataProductOutput struct{ *pulumi.OutputState }

func (DataProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataProduct)(nil)).Elem()
}

func (o DataProductOutput) ToDataProductOutput() DataProductOutput {
	return o
}

func (o DataProductOutput) ToDataProductOutputWithContext(ctx context.Context) DataProductOutput {
	return o
}

// List of available minor versions of the data product resource.
func (o DataProductOutput) AvailableMinorVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringArrayOutput { return v.AvailableMinorVersions }).(pulumi.StringArrayOutput)
}

// Resource links which exposed to the customer to query the data.
func (o DataProductOutput) ConsumptionEndpoints() ConsumptionEndpointsPropertiesResponseOutput {
	return o.ApplyT(func(v *DataProduct) ConsumptionEndpointsPropertiesResponseOutput { return v.ConsumptionEndpoints }).(ConsumptionEndpointsPropertiesResponseOutput)
}

// Current configured minor version of the data product resource.
func (o DataProductOutput) CurrentMinorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringPtrOutput { return v.CurrentMinorVersion }).(pulumi.StringPtrOutput)
}

// Customer managed encryption key details for data product.
func (o DataProductOutput) CustomerEncryptionKey() EncryptionKeyDetailsResponsePtrOutput {
	return o.ApplyT(func(v *DataProduct) EncryptionKeyDetailsResponsePtrOutput { return v.CustomerEncryptionKey }).(EncryptionKeyDetailsResponsePtrOutput)
}

// Flag to enable customer managed key encryption for data product.
func (o DataProductOutput) CustomerManagedKeyEncryptionEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringPtrOutput { return v.CustomerManagedKeyEncryptionEnabled }).(pulumi.StringPtrOutput)
}

// Documentation link for the data product based on definition file.
func (o DataProductOutput) Documentation() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.Documentation }).(pulumi.StringOutput)
}

// The managed service identities assigned to this resource.
func (o DataProductOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DataProduct) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Key vault url.
func (o DataProductOutput) KeyVaultUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.KeyVaultUrl }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DataProductOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Major version of data product.
func (o DataProductOutput) MajorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.MajorVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o DataProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network rule set for data product.
func (o DataProductOutput) Networkacls() DataProductNetworkAclsResponsePtrOutput {
	return o.ApplyT(func(v *DataProduct) DataProductNetworkAclsResponsePtrOutput { return v.Networkacls }).(DataProductNetworkAclsResponsePtrOutput)
}

// List of name or email associated with data product resource deployment.
func (o DataProductOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringArrayOutput { return v.Owners }).(pulumi.StringArrayOutput)
}

// Flag to enable or disable private link for data product resource.
func (o DataProductOutput) PrivateLinksEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringPtrOutput { return v.PrivateLinksEnabled }).(pulumi.StringPtrOutput)
}

// Product name of data product.
func (o DataProductOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.Product }).(pulumi.StringOutput)
}

// Latest provisioning state  of data product.
func (o DataProductOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Flag to enable or disable public access of data product resource.
func (o DataProductOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Data product publisher name.
func (o DataProductOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.Publisher }).(pulumi.StringOutput)
}

// Purview account url for data product to connect to.
func (o DataProductOutput) PurviewAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringPtrOutput { return v.PurviewAccount }).(pulumi.StringPtrOutput)
}

// Purview collection url for data product to connect to.
func (o DataProductOutput) PurviewCollection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringPtrOutput { return v.PurviewCollection }).(pulumi.StringPtrOutput)
}

// Flag to enable or disable redundancy for data product.
func (o DataProductOutput) Redundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringPtrOutput { return v.Redundancy }).(pulumi.StringPtrOutput)
}

// The resource GUID property of the data product resource.
func (o DataProductOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DataProductOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DataProduct) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DataProductOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataProductOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataProduct) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataProductOutput{})
}

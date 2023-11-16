// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package purview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Account resource
// Azure REST API version: 2021-12-01. Prior API version in Azure Native 1.x: 2020-12-01-preview.
//
// Other available API versions: 2020-12-01-preview, 2021-07-01.
type Account struct {
	pulumi.CustomResourceState

	// Gets or sets the status of the account.
	AccountStatus AccountPropertiesResponseAccountStatusOutput `pulumi:"accountStatus"`
	// Cloud connectors.
	// External cloud identifier used as part of scanning configuration.
	CloudConnectors CloudConnectorsResponsePtrOutput `pulumi:"cloudConnectors"`
	// Gets the time at which the entity was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Gets the creator of the entity.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// Gets the creators of the entity's object id.
	CreatedByObjectId pulumi.StringOutput `pulumi:"createdByObjectId"`
	// The URIs that are the public endpoints of the account.
	Endpoints AccountPropertiesResponseEndpointsOutput `pulumi:"endpoints"`
	// Gets or sets the friendly name.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// Identity Info on the tracked resource
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Gets or sets the location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	//  Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
	ManagedEventHubState pulumi.StringPtrOutput `pulumi:"managedEventHubState"`
	// Gets or sets the managed resource group name
	ManagedResourceGroupName pulumi.StringPtrOutput `pulumi:"managedResourceGroupName"`
	// Gets the resource identifiers of the managed resources.
	ManagedResources AccountPropertiesResponseManagedResourcesOutput `pulumi:"managedResources"`
	// Gets or sets the public network access for managed resources.
	ManagedResourcesPublicNetworkAccess pulumi.StringPtrOutput `pulumi:"managedResourcesPublicNetworkAccess"`
	// Gets or sets the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the private endpoint connections information.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Gets or sets the state of the provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the public network access.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Gets or sets the Sku.
	Sku AccountResponseSkuOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData TrackedResourceResponseSystemDataOutput `pulumi:"systemData"`
	// Tags on the azure resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets the type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ManagedEventHubState == nil {
		args.ManagedEventHubState = pulumi.StringPtr("NotSpecified")
	}
	if args.ManagedResourcesPublicNetworkAccess == nil {
		args.ManagedResourcesPublicNetworkAccess = pulumi.StringPtr("NotSpecified")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:purview/v20201201preview:Account"),
		},
		{
			Type: pulumi.String("azure-native:purview/v20210701:Account"),
		},
		{
			Type: pulumi.String("azure-native:purview/v20211201:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("azure-native:purview:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-native:purview:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
}

type AccountState struct {
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the account.
	AccountName *string `pulumi:"accountName"`
	// Identity Info on the tracked resource
	Identity *Identity `pulumi:"identity"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	//  Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
	ManagedEventHubState *string `pulumi:"managedEventHubState"`
	// Gets or sets the managed resource group name
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Gets or sets the public network access for managed resources.
	ManagedResourcesPublicNetworkAccess *string `pulumi:"managedResourcesPublicNetworkAccess"`
	// Gets or sets the public network access.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags on the azure resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the account.
	AccountName pulumi.StringPtrInput
	// Identity Info on the tracked resource
	Identity IdentityPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	//  Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
	ManagedEventHubState pulumi.StringPtrInput
	// Gets or sets the managed resource group name
	ManagedResourceGroupName pulumi.StringPtrInput
	// Gets or sets the public network access for managed resources.
	ManagedResourcesPublicNetworkAccess pulumi.StringPtrInput
	// Gets or sets the public network access.
	PublicNetworkAccess pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Tags on the azure resource.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

// Gets or sets the status of the account.
func (o AccountOutput) AccountStatus() AccountPropertiesResponseAccountStatusOutput {
	return o.ApplyT(func(v *Account) AccountPropertiesResponseAccountStatusOutput { return v.AccountStatus }).(AccountPropertiesResponseAccountStatusOutput)
}

// Cloud connectors.
// External cloud identifier used as part of scanning configuration.
func (o AccountOutput) CloudConnectors() CloudConnectorsResponsePtrOutput {
	return o.ApplyT(func(v *Account) CloudConnectorsResponsePtrOutput { return v.CloudConnectors }).(CloudConnectorsResponsePtrOutput)
}

// Gets the time at which the entity was created.
func (o AccountOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Gets the creator of the entity.
func (o AccountOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// Gets the creators of the entity's object id.
func (o AccountOutput) CreatedByObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.CreatedByObjectId }).(pulumi.StringOutput)
}

// The URIs that are the public endpoints of the account.
func (o AccountOutput) Endpoints() AccountPropertiesResponseEndpointsOutput {
	return o.ApplyT(func(v *Account) AccountPropertiesResponseEndpointsOutput { return v.Endpoints }).(AccountPropertiesResponseEndpointsOutput)
}

// Gets or sets the friendly name.
func (o AccountOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

// Identity Info on the tracked resource
func (o AccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Account) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Gets or sets the location.
func (o AccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the state of managed eventhub. If enabled managed eventhub will be created, if disabled the managed eventhub will be removed.
func (o AccountOutput) ManagedEventHubState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.ManagedEventHubState }).(pulumi.StringPtrOutput)
}

// Gets or sets the managed resource group name
func (o AccountOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

// Gets the resource identifiers of the managed resources.
func (o AccountOutput) ManagedResources() AccountPropertiesResponseManagedResourcesOutput {
	return o.ApplyT(func(v *Account) AccountPropertiesResponseManagedResourcesOutput { return v.ManagedResources }).(AccountPropertiesResponseManagedResourcesOutput)
}

// Gets or sets the public network access for managed resources.
func (o AccountOutput) ManagedResourcesPublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.ManagedResourcesPublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Gets or sets the name.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the private endpoint connections information.
func (o AccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Account) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Gets or sets the state of the provisioning.
func (o AccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the public network access.
func (o AccountOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Gets or sets the Sku.
func (o AccountOutput) Sku() AccountResponseSkuOutput {
	return o.ApplyT(func(v *Account) AccountResponseSkuOutput { return v.Sku }).(AccountResponseSkuOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AccountOutput) SystemData() TrackedResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *Account) TrackedResourceResponseSystemDataOutput { return v.SystemData }).(TrackedResourceResponseSystemDataOutput)
}

// Tags on the azure resource.
func (o AccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Account) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type.
func (o AccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}

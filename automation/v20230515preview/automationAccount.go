// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230515preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the automation account type.
type AutomationAccount struct {
	pulumi.CustomResourceState

	// URL of automation hybrid service which is used for hybrid worker on-boarding.
	AutomationHybridServiceUrl pulumi.StringPtrOutput `pulumi:"automationHybridServiceUrl"`
	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether requests using non-AAD authentication are blocked
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// Encryption properties for the automation account
	Encryption EncryptionPropertiesResponsePtrOutput `pulumi:"encryption"`
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Gets or sets the last modified by.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// Gets the last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of Automation operations supported by the Automation resource provider.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Indicates whether traffic on the non-ARM endpoint (Webhook/Agent) is allowed from the public internet
	PublicNetworkAccess pulumi.BoolPtrOutput `pulumi:"publicNetworkAccess"`
	// Gets or sets the SKU of account.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Gets status of account.
	State pulumi.StringOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAutomationAccount registers a new resource with the given unique name, arguments, and options.
func NewAutomationAccount(ctx *pulumi.Context,
	name string, args *AutomationAccountArgs, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20210622:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:AutomationAccount"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20231101:AutomationAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AutomationAccount
	err := ctx.RegisterResource("azure-native:automation/v20230515preview:AutomationAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomationAccount gets an existing AutomationAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomationAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationAccountState, opts ...pulumi.ResourceOption) (*AutomationAccount, error) {
	var resource AutomationAccount
	err := ctx.ReadResource("azure-native:automation/v20230515preview:AutomationAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutomationAccount resources.
type automationAccountState struct {
}

type AutomationAccountState struct {
}

func (AutomationAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountState)(nil)).Elem()
}

type automationAccountArgs struct {
	// The name of the automation account.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// Indicates whether requests using non-AAD authentication are blocked
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Set the encryption properties for the automation account
	Encryption *EncryptionProperties `pulumi:"encryption"`
	// Sets the identity property for automation account
	Identity *Identity `pulumi:"identity"`
	// Gets or sets the location of the resource.
	Location *string `pulumi:"location"`
	// Gets or sets name of the resource.
	Name *string `pulumi:"name"`
	// Indicates whether traffic on the non-ARM endpoint (Webhook/Agent) is allowed from the public internet
	PublicNetworkAccess *bool `pulumi:"publicNetworkAccess"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets account SKU.
	Sku *Sku `pulumi:"sku"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AutomationAccount resource.
type AutomationAccountArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringPtrInput
	// Indicates whether requests using non-AAD authentication are blocked
	DisableLocalAuth pulumi.BoolPtrInput
	// Set the encryption properties for the automation account
	Encryption EncryptionPropertiesPtrInput
	// Sets the identity property for automation account
	Identity IdentityPtrInput
	// Gets or sets the location of the resource.
	Location pulumi.StringPtrInput
	// Gets or sets name of the resource.
	Name pulumi.StringPtrInput
	// Indicates whether traffic on the non-ARM endpoint (Webhook/Agent) is allowed from the public internet
	PublicNetworkAccess pulumi.BoolPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets account SKU.
	Sku SkuPtrInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (AutomationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationAccountArgs)(nil)).Elem()
}

type AutomationAccountInput interface {
	pulumi.Input

	ToAutomationAccountOutput() AutomationAccountOutput
	ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput
}

func (*AutomationAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationAccount)(nil)).Elem()
}

func (i *AutomationAccount) ToAutomationAccountOutput() AutomationAccountOutput {
	return i.ToAutomationAccountOutputWithContext(context.Background())
}

func (i *AutomationAccount) ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationAccountOutput)
}

type AutomationAccountOutput struct{ *pulumi.OutputState }

func (AutomationAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationAccount)(nil)).Elem()
}

func (o AutomationAccountOutput) ToAutomationAccountOutput() AutomationAccountOutput {
	return o
}

func (o AutomationAccountOutput) ToAutomationAccountOutputWithContext(ctx context.Context) AutomationAccountOutput {
	return o
}

// URL of automation hybrid service which is used for hybrid worker on-boarding.
func (o AutomationAccountOutput) AutomationHybridServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.AutomationHybridServiceUrl }).(pulumi.StringPtrOutput)
}

// Gets the creation time.
func (o AutomationAccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o AutomationAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether requests using non-AAD authentication are blocked
func (o AutomationAccountOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Encryption properties for the automation account
func (o AutomationAccountOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AutomationAccount) EncryptionPropertiesResponsePtrOutput { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

// Gets or sets the etag of the resource.
func (o AutomationAccountOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Identity for the resource.
func (o AutomationAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *AutomationAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Gets or sets the last modified by.
func (o AutomationAccountOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// Gets the last modified time.
func (o AutomationAccountOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o AutomationAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AutomationAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of Automation operations supported by the Automation resource provider.
func (o AutomationAccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *AutomationAccount) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Indicates whether traffic on the non-ARM endpoint (Webhook/Agent) is allowed from the public internet
func (o AutomationAccountOutput) PublicNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.BoolPtrOutput { return v.PublicNetworkAccess }).(pulumi.BoolPtrOutput)
}

// Gets or sets the SKU of account.
func (o AutomationAccountOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *AutomationAccount) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Gets status of account.
func (o AutomationAccountOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AutomationAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AutomationAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AutomationAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AutomationAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AutomationAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AutomationAccountOutput{})
}

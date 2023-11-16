// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AzureStorageInfo dictionary resource.
type WebAppAzureStorageAccountsSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure storage accounts.
	Properties AzureStorageInfoValueResponseMapOutput `pulumi:"properties"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppAzureStorageAccountsSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppAzureStorageAccountsSlot(ctx *pulumi.Context,
	name string, args *WebAppAzureStorageAccountsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccountsSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppAzureStorageAccountsSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppAzureStorageAccountsSlot
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppAzureStorageAccountsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppAzureStorageAccountsSlot gets an existing WebAppAzureStorageAccountsSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppAzureStorageAccountsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAzureStorageAccountsSlotState, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccountsSlot, error) {
	var resource WebAppAzureStorageAccountsSlot
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppAzureStorageAccountsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppAzureStorageAccountsSlot resources.
type webAppAzureStorageAccountsSlotState struct {
}

type WebAppAzureStorageAccountsSlotState struct {
}

func (WebAppAzureStorageAccountsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsSlotState)(nil)).Elem()
}

type webAppAzureStorageAccountsSlotArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Azure storage accounts.
	Properties map[string]AzureStorageInfoValue `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppAzureStorageAccountsSlot resource.
type WebAppAzureStorageAccountsSlotArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Azure storage accounts.
	Properties AzureStorageInfoValueMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will update the Azure storage account configurations for the production slot.
	Slot pulumi.StringInput
}

func (WebAppAzureStorageAccountsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsSlotArgs)(nil)).Elem()
}

type WebAppAzureStorageAccountsSlotInput interface {
	pulumi.Input

	ToWebAppAzureStorageAccountsSlotOutput() WebAppAzureStorageAccountsSlotOutput
	ToWebAppAzureStorageAccountsSlotOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsSlotOutput
}

func (*WebAppAzureStorageAccountsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccountsSlot)(nil)).Elem()
}

func (i *WebAppAzureStorageAccountsSlot) ToWebAppAzureStorageAccountsSlotOutput() WebAppAzureStorageAccountsSlotOutput {
	return i.ToWebAppAzureStorageAccountsSlotOutputWithContext(context.Background())
}

func (i *WebAppAzureStorageAccountsSlot) ToWebAppAzureStorageAccountsSlotOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAzureStorageAccountsSlotOutput)
}

type WebAppAzureStorageAccountsSlotOutput struct{ *pulumi.OutputState }

func (WebAppAzureStorageAccountsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccountsSlot)(nil)).Elem()
}

func (o WebAppAzureStorageAccountsSlotOutput) ToWebAppAzureStorageAccountsSlotOutput() WebAppAzureStorageAccountsSlotOutput {
	return o
}

func (o WebAppAzureStorageAccountsSlotOutput) ToWebAppAzureStorageAccountsSlotOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsSlotOutput {
	return o
}

// Kind of resource.
func (o WebAppAzureStorageAccountsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppAzureStorageAccountsSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure storage accounts.
func (o WebAppAzureStorageAccountsSlotOutput) Properties() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) AzureStorageInfoValueResponseMapOutput { return v.Properties }).(AzureStorageInfoValueResponseMapOutput)
}

// The system metadata relating to this resource.
func (o WebAppAzureStorageAccountsSlotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o WebAppAzureStorageAccountsSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAzureStorageAccountsSlotOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Slot Config names azure resource.
type WebAppSlotConfigurationNames struct {
	pulumi.CustomResourceState

	// List of application settings names.
	AppSettingNames pulumi.StringArrayOutput `pulumi:"appSettingNames"`
	// List of external Azure storage account identifiers.
	AzureStorageConfigNames pulumi.StringArrayOutput `pulumi:"azureStorageConfigNames"`
	// List of connection string names.
	ConnectionStringNames pulumi.StringArrayOutput `pulumi:"connectionStringNames"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSlotConfigurationNames registers a new resource with the given unique name, arguments, and options.
func NewWebAppSlotConfigurationNames(ctx *pulumi.Context,
	name string, args *WebAppSlotConfigurationNamesArgs, opts ...pulumi.ResourceOption) (*WebAppSlotConfigurationNames, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSlotConfigurationNames"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppSlotConfigurationNames
	err := ctx.RegisterResource("azure-native:web/v20220901:WebAppSlotConfigurationNames", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSlotConfigurationNames gets an existing WebAppSlotConfigurationNames resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSlotConfigurationNames(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSlotConfigurationNamesState, opts ...pulumi.ResourceOption) (*WebAppSlotConfigurationNames, error) {
	var resource WebAppSlotConfigurationNames
	err := ctx.ReadResource("azure-native:web/v20220901:WebAppSlotConfigurationNames", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSlotConfigurationNames resources.
type webAppSlotConfigurationNamesState struct {
}

type WebAppSlotConfigurationNamesState struct {
}

func (WebAppSlotConfigurationNamesState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotConfigurationNamesState)(nil)).Elem()
}

type webAppSlotConfigurationNamesArgs struct {
	// List of application settings names.
	AppSettingNames []string `pulumi:"appSettingNames"`
	// List of external Azure storage account identifiers.
	AzureStorageConfigNames []string `pulumi:"azureStorageConfigNames"`
	// List of connection string names.
	ConnectionStringNames []string `pulumi:"connectionStringNames"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppSlotConfigurationNames resource.
type WebAppSlotConfigurationNamesArgs struct {
	// List of application settings names.
	AppSettingNames pulumi.StringArrayInput
	// List of external Azure storage account identifiers.
	AzureStorageConfigNames pulumi.StringArrayInput
	// List of connection string names.
	ConnectionStringNames pulumi.StringArrayInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppSlotConfigurationNamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotConfigurationNamesArgs)(nil)).Elem()
}

type WebAppSlotConfigurationNamesInput interface {
	pulumi.Input

	ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput
	ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput
}

func (*WebAppSlotConfigurationNames) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSlotConfigurationNames)(nil)).Elem()
}

func (i *WebAppSlotConfigurationNames) ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput {
	return i.ToWebAppSlotConfigurationNamesOutputWithContext(context.Background())
}

func (i *WebAppSlotConfigurationNames) ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSlotConfigurationNamesOutput)
}

func (i *WebAppSlotConfigurationNames) ToOutput(ctx context.Context) pulumix.Output[*WebAppSlotConfigurationNames] {
	return pulumix.Output[*WebAppSlotConfigurationNames]{
		OutputState: i.ToWebAppSlotConfigurationNamesOutputWithContext(ctx).OutputState,
	}
}

type WebAppSlotConfigurationNamesOutput struct{ *pulumi.OutputState }

func (WebAppSlotConfigurationNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSlotConfigurationNames)(nil)).Elem()
}

func (o WebAppSlotConfigurationNamesOutput) ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput {
	return o
}

func (o WebAppSlotConfigurationNamesOutput) ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput {
	return o
}

func (o WebAppSlotConfigurationNamesOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppSlotConfigurationNames] {
	return pulumix.Output[*WebAppSlotConfigurationNames]{
		OutputState: o.OutputState,
	}
}

// List of application settings names.
func (o WebAppSlotConfigurationNamesOutput) AppSettingNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringArrayOutput { return v.AppSettingNames }).(pulumi.StringArrayOutput)
}

// List of external Azure storage account identifiers.
func (o WebAppSlotConfigurationNamesOutput) AzureStorageConfigNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringArrayOutput { return v.AzureStorageConfigNames }).(pulumi.StringArrayOutput)
}

// List of connection string names.
func (o WebAppSlotConfigurationNamesOutput) ConnectionStringNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringArrayOutput { return v.ConnectionStringNames }).(pulumi.StringArrayOutput)
}

// Kind of resource.
func (o WebAppSlotConfigurationNamesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppSlotConfigurationNamesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppSlotConfigurationNamesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSlotConfigurationNamesOutput{})
}

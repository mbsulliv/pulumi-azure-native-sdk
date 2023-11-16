// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The IoT Central application.
type App struct {
	pulumi.CustomResourceState

	// The ID of the application.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The display name of the application.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The managed identities for the IoT Central application.
	Identity SystemAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ARM resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A valid instance SKU.
	Sku AppSkuInfoResponseOutput `pulumi:"sku"`
	// The current state of the application.
	State pulumi.StringOutput `pulumi:"state"`
	// The subdomain of the application.
	Subdomain pulumi.StringPtrOutput `pulumi:"subdomain"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
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
			Type: pulumi.String("azure-native:iotcentral:App"),
		},
		{
			Type: pulumi.String("azure-native:iotcentral/v20180901:App"),
		},
		{
			Type: pulumi.String("azure-native:iotcentral/v20211101preview:App"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource App
	err := ctx.RegisterResource("azure-native:iotcentral/v20210601:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("azure-native:iotcentral/v20210601:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
}

type AppState struct {
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// The display name of the application.
	DisplayName *string `pulumi:"displayName"`
	// The managed identities for the IoT Central application.
	Identity *SystemAssignedServiceIdentity `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the IoT Central application.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ARM resource name of the IoT Central application.
	ResourceName *string `pulumi:"resourceName"`
	// A valid instance SKU.
	Sku AppSkuInfo `pulumi:"sku"`
	// The subdomain of the application.
	Subdomain *string `pulumi:"subdomain"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
	Template *string `pulumi:"template"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// The display name of the application.
	DisplayName pulumi.StringPtrInput
	// The managed identities for the IoT Central application.
	Identity SystemAssignedServiceIdentityPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the IoT Central application.
	ResourceGroupName pulumi.StringInput
	// The ARM resource name of the IoT Central application.
	ResourceName pulumi.StringPtrInput
	// A valid instance SKU.
	Sku AppSkuInfoInput
	// The subdomain of the application.
	Subdomain pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
	Template pulumi.StringPtrInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

// The ID of the application.
func (o AppOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// The display name of the application.
func (o AppOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The managed identities for the IoT Central application.
func (o AppOutput) Identity() SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *App) SystemAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(SystemAssignedServiceIdentityResponsePtrOutput)
}

// The resource location.
func (o AppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The ARM resource name.
func (o AppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A valid instance SKU.
func (o AppOutput) Sku() AppSkuInfoResponseOutput {
	return o.ApplyT(func(v *App) AppSkuInfoResponseOutput { return v.Sku }).(AppSkuInfoResponseOutput)
}

// The current state of the application.
func (o AppOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The subdomain of the application.
func (o AppOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Subdomain }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o AppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *App) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
func (o AppOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Template }).(pulumi.StringPtrOutput)
}

// The resource type.
func (o AppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}

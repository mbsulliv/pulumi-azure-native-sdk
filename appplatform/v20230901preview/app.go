// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// App resource payload
type App struct {
	pulumi.CustomResourceState

	// The Managed Identity type of the app resource
	Identity ManagedIdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	// The GEO location of the application, always the same with its parent resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the App resource
	Properties AppResourcePropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
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
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToAppResourcePropertiesPtrOutput().ApplyT(func(v *AppResourceProperties) *AppResourceProperties { return v.Defaults() }).(AppResourcePropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:App"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:App"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource App
	err := ctx.RegisterResource("azure-native:appplatform/v20230901preview:App", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:appplatform/v20230901preview:App", name, id, state, &resource, opts...)
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
	// The name of the App resource.
	AppName *string `pulumi:"appName"`
	// The Managed Identity type of the app resource
	Identity *ManagedIdentityProperties `pulumi:"identity"`
	// The GEO location of the application, always the same with its parent resource
	Location *string `pulumi:"location"`
	// Properties of the App resource
	Properties *AppResourceProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// The name of the App resource.
	AppName pulumi.StringPtrInput
	// The Managed Identity type of the app resource
	Identity ManagedIdentityPropertiesPtrInput
	// The GEO location of the application, always the same with its parent resource
	Location pulumi.StringPtrInput
	// Properties of the App resource
	Properties AppResourcePropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
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

// The Managed Identity type of the app resource
func (o AppOutput) Identity() ManagedIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *App) ManagedIdentityPropertiesResponsePtrOutput { return v.Identity }).(ManagedIdentityPropertiesResponsePtrOutput)
}

// The GEO location of the application, always the same with its parent resource
func (o AppOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o AppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the App resource
func (o AppOutput) Properties() AppResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *App) AppResourcePropertiesResponseOutput { return v.Properties }).(AppResourcePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AppOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *App) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o AppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}

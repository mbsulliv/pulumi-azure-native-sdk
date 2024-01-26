// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for Application properties.
type Application struct {
	pulumi.CustomResourceState

	// Resource Type of Application.
	ApplicationType pulumi.StringPtrOutput `pulumi:"applicationType"`
	// Command Line Arguments for Application.
	CommandLineArguments pulumi.StringPtrOutput `pulumi:"commandLineArguments"`
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting pulumi.StringOutput `pulumi:"commandLineSetting"`
	// Description of Application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies a path for the executable file for the application.
	FilePath pulumi.StringPtrOutput `pulumi:"filePath"`
	// Friendly name of Application.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// the icon a 64 bit string as a byte array.
	IconContent pulumi.StringOutput `pulumi:"iconContent"`
	// Hash of the icon.
	IconHash pulumi.StringOutput `pulumi:"iconHash"`
	// Index of the icon.
	IconIndex pulumi.IntPtrOutput `pulumi:"iconIndex"`
	// Path to icon.
	IconPath pulumi.StringPtrOutput `pulumi:"iconPath"`
	// Specifies the package application Id for MSIX applications
	MsixPackageApplicationId pulumi.StringPtrOutput `pulumi:"msixPackageApplicationId"`
	// Specifies the package family name for MSIX applications
	MsixPackageFamilyName pulumi.StringPtrOutput `pulumi:"msixPackageFamilyName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ObjectId of Application. (internal use)
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal pulumi.BoolPtrOutput `pulumi:"showInPortal"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGroupName'")
	}
	if args.CommandLineSetting == nil {
		return nil, errors.New("invalid value for required argument 'CommandLineSetting'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220909:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230707preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230905:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20231004preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20240116preview:Application"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20231101preview:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20231101preview:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The name of the application group
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	// The name of the application within the specified application group
	ApplicationName *string `pulumi:"applicationName"`
	// Resource Type of Application.
	ApplicationType *string `pulumi:"applicationType"`
	// Command Line Arguments for Application.
	CommandLineArguments *string `pulumi:"commandLineArguments"`
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting string `pulumi:"commandLineSetting"`
	// Description of Application.
	Description *string `pulumi:"description"`
	// Specifies a path for the executable file for the application.
	FilePath *string `pulumi:"filePath"`
	// Friendly name of Application.
	FriendlyName *string `pulumi:"friendlyName"`
	// Index of the icon.
	IconIndex *int `pulumi:"iconIndex"`
	// Path to icon.
	IconPath *string `pulumi:"iconPath"`
	// Specifies the package application Id for MSIX applications
	MsixPackageApplicationId *string `pulumi:"msixPackageApplicationId"`
	// Specifies the package family name for MSIX applications
	MsixPackageFamilyName *string `pulumi:"msixPackageFamilyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal *bool `pulumi:"showInPortal"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The name of the application group
	ApplicationGroupName pulumi.StringInput
	// The name of the application within the specified application group
	ApplicationName pulumi.StringPtrInput
	// Resource Type of Application.
	ApplicationType pulumi.StringPtrInput
	// Command Line Arguments for Application.
	CommandLineArguments pulumi.StringPtrInput
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting pulumi.StringInput
	// Description of Application.
	Description pulumi.StringPtrInput
	// Specifies a path for the executable file for the application.
	FilePath pulumi.StringPtrInput
	// Friendly name of Application.
	FriendlyName pulumi.StringPtrInput
	// Index of the icon.
	IconIndex pulumi.IntPtrInput
	// Path to icon.
	IconPath pulumi.StringPtrInput
	// Specifies the package application Id for MSIX applications
	MsixPackageApplicationId pulumi.StringPtrInput
	// Specifies the package family name for MSIX applications
	MsixPackageFamilyName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal pulumi.BoolPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// Resource Type of Application.
func (o ApplicationOutput) ApplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.ApplicationType }).(pulumi.StringPtrOutput)
}

// Command Line Arguments for Application.
func (o ApplicationOutput) CommandLineArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.CommandLineArguments }).(pulumi.StringPtrOutput)
}

// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
func (o ApplicationOutput) CommandLineSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.CommandLineSetting }).(pulumi.StringOutput)
}

// Description of Application.
func (o ApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies a path for the executable file for the application.
func (o ApplicationOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.FilePath }).(pulumi.StringPtrOutput)
}

// Friendly name of Application.
func (o ApplicationOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// the icon a 64 bit string as a byte array.
func (o ApplicationOutput) IconContent() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.IconContent }).(pulumi.StringOutput)
}

// Hash of the icon.
func (o ApplicationOutput) IconHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.IconHash }).(pulumi.StringOutput)
}

// Index of the icon.
func (o ApplicationOutput) IconIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.IntPtrOutput { return v.IconIndex }).(pulumi.IntPtrOutput)
}

// Path to icon.
func (o ApplicationOutput) IconPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.IconPath }).(pulumi.StringPtrOutput)
}

// Specifies the package application Id for MSIX applications
func (o ApplicationOutput) MsixPackageApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.MsixPackageApplicationId }).(pulumi.StringPtrOutput)
}

// Specifies the package family name for MSIX applications
func (o ApplicationOutput) MsixPackageFamilyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.MsixPackageFamilyName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of Application. (internal use)
func (o ApplicationOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

// Specifies whether to show the RemoteApp program in the RD Web Access server.
func (o ApplicationOutput) ShowInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.BoolPtrOutput { return v.ShowInPortal }).(pulumi.BoolPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApplicationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Application) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}

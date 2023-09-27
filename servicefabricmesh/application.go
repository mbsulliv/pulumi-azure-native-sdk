// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabricmesh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This type describes an application resource.
// Azure REST API version: 2018-09-01-preview. Prior API version in Azure Native 1.x: 2018-09-01-preview
type Application struct {
	pulumi.CustomResourceState

	// Internal - used by Visual Studio to setup the debugging session on the local development environment.
	DebugParams pulumi.StringPtrOutput `pulumi:"debugParams"`
	// User readable description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics DiagnosticsDescriptionResponsePtrOutput `pulumi:"diagnostics"`
	// Describes the health state of an application resource.
	HealthState pulumi.StringOutput `pulumi:"healthState"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Names of the services in the application.
	ServiceNames pulumi.StringArrayOutput `pulumi:"serviceNames"`
	// Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
	Services ServiceResourceDescriptionResponseArrayOutput `pulumi:"services"`
	// Status of the application.
	Status pulumi.StringOutput `pulumi:"status"`
	// Gives additional information about the current status of the application.
	StatusDetails pulumi.StringOutput `pulumi:"statusDetails"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
	UnhealthyEvaluation pulumi.StringOutput `pulumi:"unhealthyEvaluation"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180701preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:Application"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("azure-native:servicefabricmesh:Application", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:servicefabricmesh:Application", name, id, state, &resource, opts...)
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
	// The identity of the application.
	ApplicationResourceName *string `pulumi:"applicationResourceName"`
	// Internal - used by Visual Studio to setup the debugging session on the local development environment.
	DebugParams *string `pulumi:"debugParams"`
	// User readable description of the application.
	Description *string `pulumi:"description"`
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics *DiagnosticsDescription `pulumi:"diagnostics"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
	Services []ServiceResourceDescription `pulumi:"services"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The identity of the application.
	ApplicationResourceName pulumi.StringPtrInput
	// Internal - used by Visual Studio to setup the debugging session on the local development environment.
	DebugParams pulumi.StringPtrInput
	// User readable description of the application.
	Description pulumi.StringPtrInput
	// Describes the diagnostics definition and usage for an application resource.
	Diagnostics DiagnosticsDescriptionPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Azure resource group name
	ResourceGroupName pulumi.StringInput
	// Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
	Services ServiceResourceDescriptionArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
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

func (i *Application) ToOutput(ctx context.Context) pulumix.Output[*Application] {
	return pulumix.Output[*Application]{
		OutputState: i.ToApplicationOutputWithContext(ctx).OutputState,
	}
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

func (o ApplicationOutput) ToOutput(ctx context.Context) pulumix.Output[*Application] {
	return pulumix.Output[*Application]{
		OutputState: o.OutputState,
	}
}

// Internal - used by Visual Studio to setup the debugging session on the local development environment.
func (o ApplicationOutput) DebugParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.DebugParams }).(pulumi.StringPtrOutput)
}

// User readable description of the application.
func (o ApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Describes the diagnostics definition and usage for an application resource.
func (o ApplicationOutput) Diagnostics() DiagnosticsDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *Application) DiagnosticsDescriptionResponsePtrOutput { return v.Diagnostics }).(DiagnosticsDescriptionResponsePtrOutput)
}

// Describes the health state of an application resource.
func (o ApplicationOutput) HealthState() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.HealthState }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ApplicationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the resource.
func (o ApplicationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Names of the services in the application.
func (o ApplicationOutput) ServiceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.StringArrayOutput { return v.ServiceNames }).(pulumi.StringArrayOutput)
}

// Describes the services in the application. This property is used to create or modify services of the application. On get only the name of the service is returned. The service description can be obtained by querying for the service resource.
func (o ApplicationOutput) Services() ServiceResourceDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Application) ServiceResourceDescriptionResponseArrayOutput { return v.Services }).(ServiceResourceDescriptionResponseArrayOutput)
}

// Status of the application.
func (o ApplicationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Gives additional information about the current status of the application.
func (o ApplicationOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.StatusDetails }).(pulumi.StringOutput)
}

// Resource tags.
func (o ApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// When the application's health state is not 'Ok', this additional details from service fabric Health Manager for the user to know why the application is marked unhealthy.
func (o ApplicationOutput) UnhealthyEvaluation() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.UnhealthyEvaluation }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}

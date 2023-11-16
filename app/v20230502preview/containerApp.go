// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230502preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App.
type ContainerApp struct {
	pulumi.CustomResourceState

	// Non versioned Container App configuration properties.
	Configuration ConfigurationResponsePtrOutput `pulumi:"configuration"`
	// Id used to verify domain name ownership
	CustomDomainVerificationId pulumi.StringOutput `pulumi:"customDomainVerificationId"`
	// Resource ID of environment.
	EnvironmentId pulumi.StringPtrOutput `pulumi:"environmentId"`
	// The endpoint of the eventstream of the container app.
	EventStreamEndpoint pulumi.StringOutput `pulumi:"eventStreamEndpoint"`
	// The complex type of the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Name of the latest ready revision of the Container App.
	LatestReadyRevisionName pulumi.StringOutput `pulumi:"latestReadyRevisionName"`
	// Fully Qualified Domain Name of the latest revision of the Container App.
	LatestRevisionFqdn pulumi.StringOutput `pulumi:"latestRevisionFqdn"`
	// Name of the latest revision of the Container App.
	LatestRevisionName pulumi.StringOutput `pulumi:"latestRevisionName"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// Deprecated. Resource ID of the Container App's environment.
	ManagedEnvironmentId pulumi.StringPtrOutput `pulumi:"managedEnvironmentId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Outbound IP Addresses for container app.
	OutboundIpAddresses pulumi.StringArrayOutput `pulumi:"outboundIpAddresses"`
	// Provisioning state of the Container App.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Container App versioned application definition.
	Template TemplateResponsePtrOutput `pulumi:"template"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Workload profile name to pin for container app execution.
	WorkloadProfileName pulumi.StringPtrOutput `pulumi:"workloadProfileName"`
}

// NewContainerApp registers a new resource with the given unique name, arguments, and options.
func NewContainerApp(ctx *pulumi.Context,
	name string, args *ContainerAppArgs, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Configuration != nil {
		args.Configuration = args.Configuration.ToConfigurationPtrOutput().ApplyT(func(v *Configuration) *Configuration { return v.Defaults() }).(ConfigurationPtrOutput)
	}
	if args.Template != nil {
		args.Template = args.Template.ToTemplatePtrOutput().ApplyT(func(v *Template) *Template { return v.Defaults() }).(TemplatePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220101preview:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220301:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230401preview:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230501:ContainerApp"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230801preview:ContainerApp"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ContainerApp
	err := ctx.RegisterResource("azure-native:app/v20230502preview:ContainerApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerApp gets an existing ContainerApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAppState, opts ...pulumi.ResourceOption) (*ContainerApp, error) {
	var resource ContainerApp
	err := ctx.ReadResource("azure-native:app/v20230502preview:ContainerApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerApp resources.
type containerAppState struct {
}

type ContainerAppState struct {
}

func (ContainerAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppState)(nil)).Elem()
}

type containerAppArgs struct {
	// Non versioned Container App configuration properties.
	Configuration *Configuration `pulumi:"configuration"`
	// Name of the Container App.
	ContainerAppName *string `pulumi:"containerAppName"`
	// Resource ID of environment.
	EnvironmentId *string `pulumi:"environmentId"`
	// The complex type of the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string `pulumi:"managedBy"`
	// Deprecated. Resource ID of the Container App's environment.
	ManagedEnvironmentId *string `pulumi:"managedEnvironmentId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Container App versioned application definition.
	Template *Template `pulumi:"template"`
	// Workload profile name to pin for container app execution.
	WorkloadProfileName *string `pulumi:"workloadProfileName"`
}

// The set of arguments for constructing a ContainerApp resource.
type ContainerAppArgs struct {
	// Non versioned Container App configuration properties.
	Configuration ConfigurationPtrInput
	// Name of the Container App.
	ContainerAppName pulumi.StringPtrInput
	// Resource ID of environment.
	EnvironmentId pulumi.StringPtrInput
	// The complex type of the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrInput
	// Deprecated. Resource ID of the Container App's environment.
	ManagedEnvironmentId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Container App versioned application definition.
	Template TemplatePtrInput
	// Workload profile name to pin for container app execution.
	WorkloadProfileName pulumi.StringPtrInput
}

func (ContainerAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAppArgs)(nil)).Elem()
}

type ContainerAppInput interface {
	pulumi.Input

	ToContainerAppOutput() ContainerAppOutput
	ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput
}

func (*ContainerApp) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (i *ContainerApp) ToContainerAppOutput() ContainerAppOutput {
	return i.ToContainerAppOutputWithContext(context.Background())
}

func (i *ContainerApp) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppOutput)
}

type ContainerAppOutput struct{ *pulumi.OutputState }

func (ContainerAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerApp)(nil)).Elem()
}

func (o ContainerAppOutput) ToContainerAppOutput() ContainerAppOutput {
	return o
}

func (o ContainerAppOutput) ToContainerAppOutputWithContext(ctx context.Context) ContainerAppOutput {
	return o
}

// Non versioned Container App configuration properties.
func (o ContainerAppOutput) Configuration() ConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) ConfigurationResponsePtrOutput { return v.Configuration }).(ConfigurationResponsePtrOutput)
}

// Id used to verify domain name ownership
func (o ContainerAppOutput) CustomDomainVerificationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.CustomDomainVerificationId }).(pulumi.StringOutput)
}

// Resource ID of environment.
func (o ContainerAppOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// The endpoint of the eventstream of the container app.
func (o ContainerAppOutput) EventStreamEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.EventStreamEndpoint }).(pulumi.StringOutput)
}

// The complex type of the extended location.
func (o ContainerAppOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// managed identities for the Container App to interact with other Azure services without maintaining any secrets or credentials in code.
func (o ContainerAppOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Name of the latest ready revision of the Container App.
func (o ContainerAppOutput) LatestReadyRevisionName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.LatestReadyRevisionName }).(pulumi.StringOutput)
}

// Fully Qualified Domain Name of the latest revision of the Container App.
func (o ContainerAppOutput) LatestRevisionFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.LatestRevisionFqdn }).(pulumi.StringOutput)
}

// Name of the latest revision of the Container App.
func (o ContainerAppOutput) LatestRevisionName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.LatestRevisionName }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ContainerAppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o ContainerAppOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// Deprecated. Resource ID of the Container App's environment.
func (o ContainerAppOutput) ManagedEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringPtrOutput { return v.ManagedEnvironmentId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ContainerAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Outbound IP Addresses for container app.
func (o ContainerAppOutput) OutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringArrayOutput { return v.OutboundIpAddresses }).(pulumi.StringArrayOutput)
}

// Provisioning state of the Container App.
func (o ContainerAppOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ContainerAppOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContainerApp) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ContainerAppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Container App versioned application definition.
func (o ContainerAppOutput) Template() TemplateResponsePtrOutput {
	return o.ApplyT(func(v *ContainerApp) TemplateResponsePtrOutput { return v.Template }).(TemplateResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContainerAppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Workload profile name to pin for container app execution.
func (o ContainerAppOutput) WorkloadProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerApp) pulumi.StringPtrOutput { return v.WorkloadProfileName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerAppOutput{})
}

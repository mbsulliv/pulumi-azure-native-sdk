// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the update summaries for the cluster
type UpdateSummary struct {
	pulumi.CustomResourceState

	// Current Solution Bundle version of the stamp.
	CurrentVersion pulumi.StringPtrOutput `pulumi:"currentVersion"`
	// Name of the hardware model.
	HardwareModel pulumi.StringPtrOutput `pulumi:"hardwareModel"`
	// Last time the package-specific checks were run.
	HealthCheckDate pulumi.StringPtrOutput `pulumi:"healthCheckDate"`
	// Last time the update service successfully checked for updates
	LastChecked pulumi.StringPtrOutput `pulumi:"lastChecked"`
	// Last time an update installation completed successfully.
	LastUpdated pulumi.StringPtrOutput `pulumi:"lastUpdated"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// OEM family name.
	OemFamily pulumi.StringPtrOutput `pulumi:"oemFamily"`
	// Provisioning state of the UpdateSummaries proxy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Overall update state of the stamp.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUpdateSummary registers a new resource with the given unique name, arguments, and options.
func NewUpdateSummary(ctx *pulumi.Context,
	name string, args *UpdateSummaryArgs, opts ...pulumi.ResourceOption) (*UpdateSummary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230201:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230601:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801preview:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20231101preview:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:UpdateSummary"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240215preview:UpdateSummary"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource UpdateSummary
	err := ctx.RegisterResource("azure-native:azurestackhci/v20230301:UpdateSummary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpdateSummary gets an existing UpdateSummary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpdateSummary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdateSummaryState, opts ...pulumi.ResourceOption) (*UpdateSummary, error) {
	var resource UpdateSummary
	err := ctx.ReadResource("azure-native:azurestackhci/v20230301:UpdateSummary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpdateSummary resources.
type updateSummaryState struct {
}

type UpdateSummaryState struct {
}

func (UpdateSummaryState) ElementType() reflect.Type {
	return reflect.TypeOf((*updateSummaryState)(nil)).Elem()
}

type updateSummaryArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// Current Solution Bundle version of the stamp.
	CurrentVersion *string `pulumi:"currentVersion"`
	// Name of the hardware model.
	HardwareModel *string `pulumi:"hardwareModel"`
	// Last time the package-specific checks were run.
	HealthCheckDate *string `pulumi:"healthCheckDate"`
	// Last time the update service successfully checked for updates
	LastChecked *string `pulumi:"lastChecked"`
	// Last time an update installation completed successfully.
	LastUpdated *string `pulumi:"lastUpdated"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// OEM family name.
	OemFamily *string `pulumi:"oemFamily"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Overall update state of the stamp.
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a UpdateSummary resource.
type UpdateSummaryArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// Current Solution Bundle version of the stamp.
	CurrentVersion pulumi.StringPtrInput
	// Name of the hardware model.
	HardwareModel pulumi.StringPtrInput
	// Last time the package-specific checks were run.
	HealthCheckDate pulumi.StringPtrInput
	// Last time the update service successfully checked for updates
	LastChecked pulumi.StringPtrInput
	// Last time an update installation completed successfully.
	LastUpdated pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// OEM family name.
	OemFamily pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Overall update state of the stamp.
	State pulumi.StringPtrInput
}

func (UpdateSummaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updateSummaryArgs)(nil)).Elem()
}

type UpdateSummaryInput interface {
	pulumi.Input

	ToUpdateSummaryOutput() UpdateSummaryOutput
	ToUpdateSummaryOutputWithContext(ctx context.Context) UpdateSummaryOutput
}

func (*UpdateSummary) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateSummary)(nil)).Elem()
}

func (i *UpdateSummary) ToUpdateSummaryOutput() UpdateSummaryOutput {
	return i.ToUpdateSummaryOutputWithContext(context.Background())
}

func (i *UpdateSummary) ToUpdateSummaryOutputWithContext(ctx context.Context) UpdateSummaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateSummaryOutput)
}

type UpdateSummaryOutput struct{ *pulumi.OutputState }

func (UpdateSummaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateSummary)(nil)).Elem()
}

func (o UpdateSummaryOutput) ToUpdateSummaryOutput() UpdateSummaryOutput {
	return o
}

func (o UpdateSummaryOutput) ToUpdateSummaryOutputWithContext(ctx context.Context) UpdateSummaryOutput {
	return o
}

// Current Solution Bundle version of the stamp.
func (o UpdateSummaryOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

// Name of the hardware model.
func (o UpdateSummaryOutput) HardwareModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.HardwareModel }).(pulumi.StringPtrOutput)
}

// Last time the package-specific checks were run.
func (o UpdateSummaryOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

// Last time the update service successfully checked for updates
func (o UpdateSummaryOutput) LastChecked() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.LastChecked }).(pulumi.StringPtrOutput)
}

// Last time an update installation completed successfully.
func (o UpdateSummaryOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.LastUpdated }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o UpdateSummaryOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o UpdateSummaryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OEM family name.
func (o UpdateSummaryOutput) OemFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.OemFamily }).(pulumi.StringPtrOutput)
}

// Provisioning state of the UpdateSummaries proxy resource.
func (o UpdateSummaryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Overall update state of the stamp.
func (o UpdateSummaryOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o UpdateSummaryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UpdateSummary) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o UpdateSummaryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UpdateSummaryOutput{})
}

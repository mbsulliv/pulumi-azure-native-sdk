// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Edge device resource
type DeploymentSetting struct {
	pulumi.CustomResourceState

	// Azure resource ids of Arc machines to be part of cluster.
	ArcNodeResourceIds pulumi.StringArrayOutput `pulumi:"arcNodeResourceIds"`
	// Scale units will contains list of deployment data
	DeploymentConfiguration DeploymentConfigurationResponseOutput `pulumi:"deploymentConfiguration"`
	// The deployment mode for cluster deployment.
	DeploymentMode pulumi.StringOutput `pulumi:"deploymentMode"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// DeploymentSetting provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Deployment Status reported from cluster.
	ReportedProperties ReportedPropertiesResponseOutput `pulumi:"reportedProperties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeploymentSetting registers a new resource with the given unique name, arguments, and options.
func NewDeploymentSetting(ctx *pulumi.Context,
	name string, args *DeploymentSettingArgs, opts ...pulumi.ResourceOption) (*DeploymentSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ArcNodeResourceIds == nil {
		return nil, errors.New("invalid value for required argument 'ArcNodeResourceIds'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DeploymentConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentConfiguration'")
	}
	if args.DeploymentMode == nil {
		return nil, errors.New("invalid value for required argument 'DeploymentMode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:DeploymentSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801preview:DeploymentSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20231101preview:DeploymentSetting"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DeploymentSetting
	err := ctx.RegisterResource("azure-native:azurestackhci/v20240101:DeploymentSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentSetting gets an existing DeploymentSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentSettingState, opts ...pulumi.ResourceOption) (*DeploymentSetting, error) {
	var resource DeploymentSetting
	err := ctx.ReadResource("azure-native:azurestackhci/v20240101:DeploymentSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentSetting resources.
type deploymentSettingState struct {
}

type DeploymentSettingState struct {
}

func (DeploymentSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentSettingState)(nil)).Elem()
}

type deploymentSettingArgs struct {
	// Azure resource ids of Arc machines to be part of cluster.
	ArcNodeResourceIds []string `pulumi:"arcNodeResourceIds"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// Scale units will contains list of deployment data
	DeploymentConfiguration DeploymentConfiguration `pulumi:"deploymentConfiguration"`
	// The deployment mode for cluster deployment.
	DeploymentMode string `pulumi:"deploymentMode"`
	// Name of Deployment Setting
	DeploymentSettingsName *string `pulumi:"deploymentSettingsName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DeploymentSetting resource.
type DeploymentSettingArgs struct {
	// Azure resource ids of Arc machines to be part of cluster.
	ArcNodeResourceIds pulumi.StringArrayInput
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// Scale units will contains list of deployment data
	DeploymentConfiguration DeploymentConfigurationInput
	// The deployment mode for cluster deployment.
	DeploymentMode pulumi.StringInput
	// Name of Deployment Setting
	DeploymentSettingsName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (DeploymentSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentSettingArgs)(nil)).Elem()
}

type DeploymentSettingInput interface {
	pulumi.Input

	ToDeploymentSettingOutput() DeploymentSettingOutput
	ToDeploymentSettingOutputWithContext(ctx context.Context) DeploymentSettingOutput
}

func (*DeploymentSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSetting)(nil)).Elem()
}

func (i *DeploymentSetting) ToDeploymentSettingOutput() DeploymentSettingOutput {
	return i.ToDeploymentSettingOutputWithContext(context.Background())
}

func (i *DeploymentSetting) ToDeploymentSettingOutputWithContext(ctx context.Context) DeploymentSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingOutput)
}

type DeploymentSettingOutput struct{ *pulumi.OutputState }

func (DeploymentSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSetting)(nil)).Elem()
}

func (o DeploymentSettingOutput) ToDeploymentSettingOutput() DeploymentSettingOutput {
	return o
}

func (o DeploymentSettingOutput) ToDeploymentSettingOutputWithContext(ctx context.Context) DeploymentSettingOutput {
	return o
}

// Azure resource ids of Arc machines to be part of cluster.
func (o DeploymentSettingOutput) ArcNodeResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeploymentSetting) pulumi.StringArrayOutput { return v.ArcNodeResourceIds }).(pulumi.StringArrayOutput)
}

// Scale units will contains list of deployment data
func (o DeploymentSettingOutput) DeploymentConfiguration() DeploymentConfigurationResponseOutput {
	return o.ApplyT(func(v *DeploymentSetting) DeploymentConfigurationResponseOutput { return v.DeploymentConfiguration }).(DeploymentConfigurationResponseOutput)
}

// The deployment mode for cluster deployment.
func (o DeploymentSettingOutput) DeploymentMode() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentSetting) pulumi.StringOutput { return v.DeploymentMode }).(pulumi.StringOutput)
}

// The name of the resource
func (o DeploymentSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// DeploymentSetting provisioning state
func (o DeploymentSettingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentSetting) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Deployment Status reported from cluster.
func (o DeploymentSettingOutput) ReportedProperties() ReportedPropertiesResponseOutput {
	return o.ApplyT(func(v *DeploymentSetting) ReportedPropertiesResponseOutput { return v.ReportedProperties }).(ReportedPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DeploymentSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DeploymentSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DeploymentSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentSettingOutput{})
}

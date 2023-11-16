// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the dsc node configuration.
// Azure REST API version: 2022-08-08. Prior API version in Azure Native 1.x: 2019-06-01.
//
// Other available API versions: 2023-05-15-preview, 2023-11-01.
type DscNodeConfiguration struct {
	pulumi.CustomResourceState

	// Gets or sets the configuration of the node.
	Configuration DscConfigurationAssociationPropertyResponsePtrOutput `pulumi:"configuration"`
	// Gets or sets creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// If a new build version of NodeConfiguration is required.
	IncrementNodeConfigurationBuild pulumi.BoolPtrOutput `pulumi:"incrementNodeConfigurationBuild"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of nodes with this node configuration assigned
	NodeCount pulumi.Float64PtrOutput `pulumi:"nodeCount"`
	// Source of node configuration.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDscNodeConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDscNodeConfiguration(ctx *pulumi.Context,
	name string, args *DscNodeConfigurationArgs, opts ...pulumi.ResourceOption) (*DscNodeConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20151031:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20180115:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20231101:DscNodeConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DscNodeConfiguration
	err := ctx.RegisterResource("azure-native:automation:DscNodeConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDscNodeConfiguration gets an existing DscNodeConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDscNodeConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscNodeConfigurationState, opts ...pulumi.ResourceOption) (*DscNodeConfiguration, error) {
	var resource DscNodeConfiguration
	err := ctx.ReadResource("azure-native:automation:DscNodeConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DscNodeConfiguration resources.
type dscNodeConfigurationState struct {
}

type DscNodeConfigurationState struct {
}

func (DscNodeConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscNodeConfigurationState)(nil)).Elem()
}

type dscNodeConfigurationArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the configuration of the node.
	Configuration DscConfigurationAssociationProperty `pulumi:"configuration"`
	// If a new build version of NodeConfiguration is required.
	IncrementNodeConfigurationBuild *bool `pulumi:"incrementNodeConfigurationBuild"`
	// Name of the node configuration.
	Name *string `pulumi:"name"`
	// The Dsc node configuration name.
	NodeConfigurationName *string `pulumi:"nodeConfigurationName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the source.
	Source ContentSource `pulumi:"source"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DscNodeConfiguration resource.
type DscNodeConfigurationArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the configuration of the node.
	Configuration DscConfigurationAssociationPropertyInput
	// If a new build version of NodeConfiguration is required.
	IncrementNodeConfigurationBuild pulumi.BoolPtrInput
	// Name of the node configuration.
	Name pulumi.StringPtrInput
	// The Dsc node configuration name.
	NodeConfigurationName pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the source.
	Source ContentSourceInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (DscNodeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscNodeConfigurationArgs)(nil)).Elem()
}

type DscNodeConfigurationInput interface {
	pulumi.Input

	ToDscNodeConfigurationOutput() DscNodeConfigurationOutput
	ToDscNodeConfigurationOutputWithContext(ctx context.Context) DscNodeConfigurationOutput
}

func (*DscNodeConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DscNodeConfiguration)(nil)).Elem()
}

func (i *DscNodeConfiguration) ToDscNodeConfigurationOutput() DscNodeConfigurationOutput {
	return i.ToDscNodeConfigurationOutputWithContext(context.Background())
}

func (i *DscNodeConfiguration) ToDscNodeConfigurationOutputWithContext(ctx context.Context) DscNodeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscNodeConfigurationOutput)
}

type DscNodeConfigurationOutput struct{ *pulumi.OutputState }

func (DscNodeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscNodeConfiguration)(nil)).Elem()
}

func (o DscNodeConfigurationOutput) ToDscNodeConfigurationOutput() DscNodeConfigurationOutput {
	return o
}

func (o DscNodeConfigurationOutput) ToDscNodeConfigurationOutputWithContext(ctx context.Context) DscNodeConfigurationOutput {
	return o
}

// Gets or sets the configuration of the node.
func (o DscNodeConfigurationOutput) Configuration() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) DscConfigurationAssociationPropertyResponsePtrOutput {
		return v.Configuration
	}).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

// Gets or sets creation time.
func (o DscNodeConfigurationOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// If a new build version of NodeConfiguration is required.
func (o DscNodeConfigurationOutput) IncrementNodeConfigurationBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) pulumi.BoolPtrOutput { return v.IncrementNodeConfigurationBuild }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o DscNodeConfigurationOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o DscNodeConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of nodes with this node configuration assigned
func (o DscNodeConfigurationOutput) NodeCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) pulumi.Float64PtrOutput { return v.NodeCount }).(pulumi.Float64PtrOutput)
}

// Source of node configuration.
func (o DscNodeConfigurationOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o DscNodeConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DscNodeConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DscNodeConfigurationOutput{})
}

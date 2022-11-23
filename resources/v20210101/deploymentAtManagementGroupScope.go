// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment information.
type DeploymentAtManagementGroupScope struct {
	pulumi.CustomResourceState

	// the location of the deployment.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the deployment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Deployment properties.
	Properties DeploymentPropertiesExtendedResponseOutput `pulumi:"properties"`
	// Deployment tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the deployment.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDeploymentAtManagementGroupScope registers a new resource with the given unique name, arguments, and options.
func NewDeploymentAtManagementGroupScope(ctx *pulumi.Context,
	name string, args *DeploymentAtManagementGroupScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtManagementGroupScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtManagementGroupScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:DeploymentAtManagementGroupScope"),
		},
	})
	opts = append(opts, aliases)
	var resource DeploymentAtManagementGroupScope
	err := ctx.RegisterResource("azure-native:resources/v20210101:DeploymentAtManagementGroupScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentAtManagementGroupScope gets an existing DeploymentAtManagementGroupScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentAtManagementGroupScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtManagementGroupScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtManagementGroupScope, error) {
	var resource DeploymentAtManagementGroupScope
	err := ctx.ReadResource("azure-native:resources/v20210101:DeploymentAtManagementGroupScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentAtManagementGroupScope resources.
type deploymentAtManagementGroupScopeState struct {
}

type DeploymentAtManagementGroupScopeState struct {
}

func (DeploymentAtManagementGroupScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtManagementGroupScopeState)(nil)).Elem()
}

type deploymentAtManagementGroupScopeArgs struct {
	// The name of the deployment.
	DeploymentName *string `pulumi:"deploymentName"`
	// The management group ID.
	GroupId string `pulumi:"groupId"`
	// The location to store the deployment data.
	Location *string `pulumi:"location"`
	// The deployment properties.
	Properties DeploymentProperties `pulumi:"properties"`
	// Deployment tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DeploymentAtManagementGroupScope resource.
type DeploymentAtManagementGroupScopeArgs struct {
	// The name of the deployment.
	DeploymentName pulumi.StringPtrInput
	// The management group ID.
	GroupId pulumi.StringInput
	// The location to store the deployment data.
	Location pulumi.StringPtrInput
	// The deployment properties.
	Properties DeploymentPropertiesInput
	// Deployment tags
	Tags pulumi.StringMapInput
}

func (DeploymentAtManagementGroupScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtManagementGroupScopeArgs)(nil)).Elem()
}

type DeploymentAtManagementGroupScopeInput interface {
	pulumi.Input

	ToDeploymentAtManagementGroupScopeOutput() DeploymentAtManagementGroupScopeOutput
	ToDeploymentAtManagementGroupScopeOutputWithContext(ctx context.Context) DeploymentAtManagementGroupScopeOutput
}

func (*DeploymentAtManagementGroupScope) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtManagementGroupScope)(nil)).Elem()
}

func (i *DeploymentAtManagementGroupScope) ToDeploymentAtManagementGroupScopeOutput() DeploymentAtManagementGroupScopeOutput {
	return i.ToDeploymentAtManagementGroupScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtManagementGroupScope) ToDeploymentAtManagementGroupScopeOutputWithContext(ctx context.Context) DeploymentAtManagementGroupScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtManagementGroupScopeOutput)
}

type DeploymentAtManagementGroupScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtManagementGroupScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtManagementGroupScope)(nil)).Elem()
}

func (o DeploymentAtManagementGroupScopeOutput) ToDeploymentAtManagementGroupScopeOutput() DeploymentAtManagementGroupScopeOutput {
	return o
}

func (o DeploymentAtManagementGroupScopeOutput) ToDeploymentAtManagementGroupScopeOutputWithContext(ctx context.Context) DeploymentAtManagementGroupScopeOutput {
	return o
}

// the location of the deployment.
func (o DeploymentAtManagementGroupScopeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the deployment.
func (o DeploymentAtManagementGroupScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Deployment properties.
func (o DeploymentAtManagementGroupScopeOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) DeploymentPropertiesExtendedResponseOutput {
		return v.Properties
	}).(DeploymentPropertiesExtendedResponseOutput)
}

// Deployment tags
func (o DeploymentAtManagementGroupScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the deployment.
func (o DeploymentAtManagementGroupScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtManagementGroupScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentAtManagementGroupScopeOutput{})
}

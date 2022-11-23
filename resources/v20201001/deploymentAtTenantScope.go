// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment information.
type DeploymentAtTenantScope struct {
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

// NewDeploymentAtTenantScope registers a new resource with the given unique name, arguments, and options.
func NewDeploymentAtTenantScope(ctx *pulumi.Context,
	name string, args *DeploymentAtTenantScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtTenantScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtTenantScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:DeploymentAtTenantScope"),
		},
	})
	opts = append(opts, aliases)
	var resource DeploymentAtTenantScope
	err := ctx.RegisterResource("azure-native:resources/v20201001:DeploymentAtTenantScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentAtTenantScope gets an existing DeploymentAtTenantScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentAtTenantScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtTenantScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtTenantScope, error) {
	var resource DeploymentAtTenantScope
	err := ctx.ReadResource("azure-native:resources/v20201001:DeploymentAtTenantScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentAtTenantScope resources.
type deploymentAtTenantScopeState struct {
}

type DeploymentAtTenantScopeState struct {
}

func (DeploymentAtTenantScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtTenantScopeState)(nil)).Elem()
}

type deploymentAtTenantScopeArgs struct {
	// The name of the deployment.
	DeploymentName *string `pulumi:"deploymentName"`
	// The location to store the deployment data.
	Location *string `pulumi:"location"`
	// The deployment properties.
	Properties DeploymentProperties `pulumi:"properties"`
	// Deployment tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DeploymentAtTenantScope resource.
type DeploymentAtTenantScopeArgs struct {
	// The name of the deployment.
	DeploymentName pulumi.StringPtrInput
	// The location to store the deployment data.
	Location pulumi.StringPtrInput
	// The deployment properties.
	Properties DeploymentPropertiesInput
	// Deployment tags
	Tags pulumi.StringMapInput
}

func (DeploymentAtTenantScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtTenantScopeArgs)(nil)).Elem()
}

type DeploymentAtTenantScopeInput interface {
	pulumi.Input

	ToDeploymentAtTenantScopeOutput() DeploymentAtTenantScopeOutput
	ToDeploymentAtTenantScopeOutputWithContext(ctx context.Context) DeploymentAtTenantScopeOutput
}

func (*DeploymentAtTenantScope) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtTenantScope)(nil)).Elem()
}

func (i *DeploymentAtTenantScope) ToDeploymentAtTenantScopeOutput() DeploymentAtTenantScopeOutput {
	return i.ToDeploymentAtTenantScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtTenantScope) ToDeploymentAtTenantScopeOutputWithContext(ctx context.Context) DeploymentAtTenantScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtTenantScopeOutput)
}

type DeploymentAtTenantScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtTenantScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtTenantScope)(nil)).Elem()
}

func (o DeploymentAtTenantScopeOutput) ToDeploymentAtTenantScopeOutput() DeploymentAtTenantScopeOutput {
	return o
}

func (o DeploymentAtTenantScopeOutput) ToDeploymentAtTenantScopeOutputWithContext(ctx context.Context) DeploymentAtTenantScopeOutput {
	return o
}

// the location of the deployment.
func (o DeploymentAtTenantScopeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentAtTenantScope) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the deployment.
func (o DeploymentAtTenantScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtTenantScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Deployment properties.
func (o DeploymentAtTenantScopeOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentAtTenantScope) DeploymentPropertiesExtendedResponseOutput { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

// Deployment tags
func (o DeploymentAtTenantScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentAtTenantScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the deployment.
func (o DeploymentAtTenantScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtTenantScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentAtTenantScopeOutput{})
}

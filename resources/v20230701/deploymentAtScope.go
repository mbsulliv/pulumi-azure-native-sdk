// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment information.
type DeploymentAtScope struct {
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

// NewDeploymentAtScope registers a new resource with the given unique name, arguments, and options.
func NewDeploymentAtScope(ctx *pulumi.Context,
	name string, args *DeploymentAtScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:DeploymentAtScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DeploymentAtScope
	err := ctx.RegisterResource("azure-native:resources/v20230701:DeploymentAtScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentAtScope gets an existing DeploymentAtScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentAtScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtScope, error) {
	var resource DeploymentAtScope
	err := ctx.ReadResource("azure-native:resources/v20230701:DeploymentAtScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentAtScope resources.
type deploymentAtScopeState struct {
}

type DeploymentAtScopeState struct {
}

func (DeploymentAtScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtScopeState)(nil)).Elem()
}

type deploymentAtScopeArgs struct {
	// The name of the deployment.
	DeploymentName *string `pulumi:"deploymentName"`
	// The location to store the deployment data.
	Location *string `pulumi:"location"`
	// The deployment properties.
	Properties DeploymentProperties `pulumi:"properties"`
	// The resource scope.
	Scope string `pulumi:"scope"`
	// Deployment tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DeploymentAtScope resource.
type DeploymentAtScopeArgs struct {
	// The name of the deployment.
	DeploymentName pulumi.StringPtrInput
	// The location to store the deployment data.
	Location pulumi.StringPtrInput
	// The deployment properties.
	Properties DeploymentPropertiesInput
	// The resource scope.
	Scope pulumi.StringInput
	// Deployment tags
	Tags pulumi.StringMapInput
}

func (DeploymentAtScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtScopeArgs)(nil)).Elem()
}

type DeploymentAtScopeInput interface {
	pulumi.Input

	ToDeploymentAtScopeOutput() DeploymentAtScopeOutput
	ToDeploymentAtScopeOutputWithContext(ctx context.Context) DeploymentAtScopeOutput
}

func (*DeploymentAtScope) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtScope)(nil)).Elem()
}

func (i *DeploymentAtScope) ToDeploymentAtScopeOutput() DeploymentAtScopeOutput {
	return i.ToDeploymentAtScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtScope) ToDeploymentAtScopeOutputWithContext(ctx context.Context) DeploymentAtScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtScopeOutput)
}

type DeploymentAtScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtScope)(nil)).Elem()
}

func (o DeploymentAtScopeOutput) ToDeploymentAtScopeOutput() DeploymentAtScopeOutput {
	return o
}

func (o DeploymentAtScopeOutput) ToDeploymentAtScopeOutputWithContext(ctx context.Context) DeploymentAtScopeOutput {
	return o
}

// the location of the deployment.
func (o DeploymentAtScopeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the deployment.
func (o DeploymentAtScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Deployment properties.
func (o DeploymentAtScopeOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentAtScope) DeploymentPropertiesExtendedResponseOutput { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

// Deployment tags
func (o DeploymentAtScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the deployment.
func (o DeploymentAtScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentAtScopeOutput{})
}

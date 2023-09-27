// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Deployment information.
type DeploymentAtSubscriptionScope struct {
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

// NewDeploymentAtSubscriptionScope registers a new resource with the given unique name, arguments, and options.
func NewDeploymentAtSubscriptionScope(ctx *pulumi.Context,
	name string, args *DeploymentAtSubscriptionScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtSubscriptionScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180501:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:DeploymentAtSubscriptionScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DeploymentAtSubscriptionScope
	err := ctx.RegisterResource("azure-native:resources/v20230701:DeploymentAtSubscriptionScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentAtSubscriptionScope gets an existing DeploymentAtSubscriptionScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentAtSubscriptionScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtSubscriptionScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtSubscriptionScope, error) {
	var resource DeploymentAtSubscriptionScope
	err := ctx.ReadResource("azure-native:resources/v20230701:DeploymentAtSubscriptionScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentAtSubscriptionScope resources.
type deploymentAtSubscriptionScopeState struct {
}

type DeploymentAtSubscriptionScopeState struct {
}

func (DeploymentAtSubscriptionScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtSubscriptionScopeState)(nil)).Elem()
}

type deploymentAtSubscriptionScopeArgs struct {
	// The name of the deployment.
	DeploymentName *string `pulumi:"deploymentName"`
	// The location to store the deployment data.
	Location *string `pulumi:"location"`
	// The deployment properties.
	Properties DeploymentProperties `pulumi:"properties"`
	// Deployment tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DeploymentAtSubscriptionScope resource.
type DeploymentAtSubscriptionScopeArgs struct {
	// The name of the deployment.
	DeploymentName pulumi.StringPtrInput
	// The location to store the deployment data.
	Location pulumi.StringPtrInput
	// The deployment properties.
	Properties DeploymentPropertiesInput
	// Deployment tags
	Tags pulumi.StringMapInput
}

func (DeploymentAtSubscriptionScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtSubscriptionScopeArgs)(nil)).Elem()
}

type DeploymentAtSubscriptionScopeInput interface {
	pulumi.Input

	ToDeploymentAtSubscriptionScopeOutput() DeploymentAtSubscriptionScopeOutput
	ToDeploymentAtSubscriptionScopeOutputWithContext(ctx context.Context) DeploymentAtSubscriptionScopeOutput
}

func (*DeploymentAtSubscriptionScope) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtSubscriptionScope)(nil)).Elem()
}

func (i *DeploymentAtSubscriptionScope) ToDeploymentAtSubscriptionScopeOutput() DeploymentAtSubscriptionScopeOutput {
	return i.ToDeploymentAtSubscriptionScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtSubscriptionScope) ToDeploymentAtSubscriptionScopeOutputWithContext(ctx context.Context) DeploymentAtSubscriptionScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtSubscriptionScopeOutput)
}

func (i *DeploymentAtSubscriptionScope) ToOutput(ctx context.Context) pulumix.Output[*DeploymentAtSubscriptionScope] {
	return pulumix.Output[*DeploymentAtSubscriptionScope]{
		OutputState: i.ToDeploymentAtSubscriptionScopeOutputWithContext(ctx).OutputState,
	}
}

type DeploymentAtSubscriptionScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtSubscriptionScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtSubscriptionScope)(nil)).Elem()
}

func (o DeploymentAtSubscriptionScopeOutput) ToDeploymentAtSubscriptionScopeOutput() DeploymentAtSubscriptionScopeOutput {
	return o
}

func (o DeploymentAtSubscriptionScopeOutput) ToDeploymentAtSubscriptionScopeOutputWithContext(ctx context.Context) DeploymentAtSubscriptionScopeOutput {
	return o
}

func (o DeploymentAtSubscriptionScopeOutput) ToOutput(ctx context.Context) pulumix.Output[*DeploymentAtSubscriptionScope] {
	return pulumix.Output[*DeploymentAtSubscriptionScope]{
		OutputState: o.OutputState,
	}
}

// the location of the deployment.
func (o DeploymentAtSubscriptionScopeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the deployment.
func (o DeploymentAtSubscriptionScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Deployment properties.
func (o DeploymentAtSubscriptionScopeOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) DeploymentPropertiesExtendedResponseOutput { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

// Deployment tags
func (o DeploymentAtSubscriptionScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the deployment.
func (o DeploymentAtSubscriptionScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentAtSubscriptionScopeOutput{})
}

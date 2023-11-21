// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Application Group object
type ApplicationGroup struct {
	pulumi.CustomResourceState

	// The Unique identifier for application group.Supports SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
	ClientAppGroupIdentifier pulumi.StringOutput `pulumi:"clientAppGroupIdentifier"`
	// Determines if Application Group is allowed to create connection with namespace or not. Once the isEnabled is set to false, all the existing connections of application group gets dropped and no new connections will be allowed
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of group policies that define the behavior of application group. The policies can support resource governance scenarios such as limiting ingress or egress traffic.
	Policies ThrottlingPolicyResponseArrayOutput `pulumi:"policies"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationGroup registers a new resource with the given unique name, arguments, and options.
func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientAppGroupIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ClientAppGroupIdentifier'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20230101preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20240101:ApplicationGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationGroup
	err := ctx.RegisterResource("azure-native:eventhub/v20221001preview:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGroup gets an existing ApplicationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("azure-native:eventhub/v20221001preview:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGroup resources.
type applicationGroupState struct {
}

type ApplicationGroupState struct {
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	// The Application Group name
	ApplicationGroupName *string `pulumi:"applicationGroupName"`
	// The Unique identifier for application group.Supports SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
	ClientAppGroupIdentifier string `pulumi:"clientAppGroupIdentifier"`
	// Determines if Application Group is allowed to create connection with namespace or not. Once the isEnabled is set to false, all the existing connections of application group gets dropped and no new connections will be allowed
	IsEnabled *bool `pulumi:"isEnabled"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// List of group policies that define the behavior of application group. The policies can support resource governance scenarios such as limiting ingress or egress traffic.
	Policies []ThrottlingPolicy `pulumi:"policies"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ApplicationGroup resource.
type ApplicationGroupArgs struct {
	// The Application Group name
	ApplicationGroupName pulumi.StringPtrInput
	// The Unique identifier for application group.Supports SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
	ClientAppGroupIdentifier pulumi.StringInput
	// Determines if Application Group is allowed to create connection with namespace or not. Once the isEnabled is set to false, all the existing connections of application group gets dropped and no new connections will be allowed
	IsEnabled pulumi.BoolPtrInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// List of group policies that define the behavior of application group. The policies can support resource governance scenarios such as limiting ingress or egress traffic.
	Policies ThrottlingPolicyArrayInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

// The Unique identifier for application group.Supports SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
func (o ApplicationGroupOutput) ClientAppGroupIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ClientAppGroupIdentifier }).(pulumi.StringOutput)
}

// Determines if Application Group is allowed to create connection with namespace or not. Once the isEnabled is set to false, all the existing connections of application group gets dropped and no new connections will be allowed
func (o ApplicationGroupOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o ApplicationGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApplicationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of group policies that define the behavior of application group. The policies can support resource governance scenarios such as limiting ingress or egress traffic.
func (o ApplicationGroupOutput) Policies() ThrottlingPolicyResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGroup) ThrottlingPolicyResponseArrayOutput { return v.Policies }).(ThrottlingPolicyResponseArrayOutput)
}

// The system meta data relating to this resource.
func (o ApplicationGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o ApplicationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
}

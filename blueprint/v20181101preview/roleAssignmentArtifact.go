// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Blueprint artifact that applies a Role assignment.
type RoleAssignmentArtifact struct {
	pulumi.CustomResourceState

	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayOutput `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Specifies the kind of blueprint artifact.
	// Expected value is 'roleAssignment'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Array of user or group identities in Azure Active Directory. The roleDefinition will apply to each identity.
	PrincipalIds pulumi.AnyOutput `pulumi:"principalIds"`
	// RoleAssignment will be scope to this resourceGroup. If empty, it scopes to the subscription.
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// Azure resource ID of the RoleDefinition.
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoleAssignmentArtifact registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignmentArtifact(ctx *pulumi.Context,
	name string, args *RoleAssignmentArtifactArgs, opts ...pulumi.ResourceOption) (*RoleAssignmentArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.PrincipalIds == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalIds'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	if args.RoleDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'RoleDefinitionId'")
	}
	args.Kind = pulumi.String("roleAssignment")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blueprint:RoleAssignmentArtifact"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RoleAssignmentArtifact
	err := ctx.RegisterResource("azure-native:blueprint/v20181101preview:RoleAssignmentArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssignmentArtifact gets an existing RoleAssignmentArtifact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignmentArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentArtifactState, opts ...pulumi.ResourceOption) (*RoleAssignmentArtifact, error) {
	var resource RoleAssignmentArtifact
	err := ctx.ReadResource("azure-native:blueprint/v20181101preview:RoleAssignmentArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignmentArtifact resources.
type roleAssignmentArtifactState struct {
}

type RoleAssignmentArtifactState struct {
}

func (RoleAssignmentArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArtifactState)(nil)).Elem()
}

type roleAssignmentArtifactArgs struct {
	// Name of the blueprint artifact.
	ArtifactName *string `pulumi:"artifactName"`
	// Name of the blueprint definition.
	BlueprintName string `pulumi:"blueprintName"`
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn []string `pulumi:"dependsOn"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Specifies the kind of blueprint artifact.
	// Expected value is 'roleAssignment'.
	Kind string `pulumi:"kind"`
	// Array of user or group identities in Azure Active Directory. The roleDefinition will apply to each identity.
	PrincipalIds interface{} `pulumi:"principalIds"`
	// RoleAssignment will be scope to this resourceGroup. If empty, it scopes to the subscription.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope string `pulumi:"resourceScope"`
	// Azure resource ID of the RoleDefinition.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}

// The set of arguments for constructing a RoleAssignmentArtifact resource.
type RoleAssignmentArtifactArgs struct {
	// Name of the blueprint artifact.
	ArtifactName pulumi.StringPtrInput
	// Name of the blueprint definition.
	BlueprintName pulumi.StringInput
	// Artifacts which need to be deployed before the specified artifact.
	DependsOn pulumi.StringArrayInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Specifies the kind of blueprint artifact.
	// Expected value is 'roleAssignment'.
	Kind pulumi.StringInput
	// Array of user or group identities in Azure Active Directory. The roleDefinition will apply to each identity.
	PrincipalIds pulumi.Input
	// RoleAssignment will be scope to this resourceGroup. If empty, it scopes to the subscription.
	ResourceGroup pulumi.StringPtrInput
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope pulumi.StringInput
	// Azure resource ID of the RoleDefinition.
	RoleDefinitionId pulumi.StringInput
}

func (RoleAssignmentArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArtifactArgs)(nil)).Elem()
}

type RoleAssignmentArtifactInput interface {
	pulumi.Input

	ToRoleAssignmentArtifactOutput() RoleAssignmentArtifactOutput
	ToRoleAssignmentArtifactOutputWithContext(ctx context.Context) RoleAssignmentArtifactOutput
}

func (*RoleAssignmentArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignmentArtifact)(nil)).Elem()
}

func (i *RoleAssignmentArtifact) ToRoleAssignmentArtifactOutput() RoleAssignmentArtifactOutput {
	return i.ToRoleAssignmentArtifactOutputWithContext(context.Background())
}

func (i *RoleAssignmentArtifact) ToRoleAssignmentArtifactOutputWithContext(ctx context.Context) RoleAssignmentArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentArtifactOutput)
}

func (i *RoleAssignmentArtifact) ToOutput(ctx context.Context) pulumix.Output[*RoleAssignmentArtifact] {
	return pulumix.Output[*RoleAssignmentArtifact]{
		OutputState: i.ToRoleAssignmentArtifactOutputWithContext(ctx).OutputState,
	}
}

type RoleAssignmentArtifactOutput struct{ *pulumi.OutputState }

func (RoleAssignmentArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignmentArtifact)(nil)).Elem()
}

func (o RoleAssignmentArtifactOutput) ToRoleAssignmentArtifactOutput() RoleAssignmentArtifactOutput {
	return o
}

func (o RoleAssignmentArtifactOutput) ToRoleAssignmentArtifactOutputWithContext(ctx context.Context) RoleAssignmentArtifactOutput {
	return o
}

func (o RoleAssignmentArtifactOutput) ToOutput(ctx context.Context) pulumix.Output[*RoleAssignmentArtifact] {
	return pulumix.Output[*RoleAssignmentArtifact]{
		OutputState: o.OutputState,
	}
}

// Artifacts which need to be deployed before the specified artifact.
func (o RoleAssignmentArtifactOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringArrayOutput { return v.DependsOn }).(pulumi.StringArrayOutput)
}

// Multi-line explain this resource.
func (o RoleAssignmentArtifactOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o RoleAssignmentArtifactOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Specifies the kind of blueprint artifact.
// Expected value is 'roleAssignment'.
func (o RoleAssignmentArtifactOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of this resource.
func (o RoleAssignmentArtifactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Array of user or group identities in Azure Active Directory. The roleDefinition will apply to each identity.
func (o RoleAssignmentArtifactOutput) PrincipalIds() pulumi.AnyOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.AnyOutput { return v.PrincipalIds }).(pulumi.AnyOutput)
}

// RoleAssignment will be scope to this resourceGroup. If empty, it scopes to the subscription.
func (o RoleAssignmentArtifactOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Azure resource ID of the RoleDefinition.
func (o RoleAssignmentArtifactOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringOutput { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

// Type of this resource.
func (o RoleAssignmentArtifactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignmentArtifact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleAssignmentArtifactOutput{})
}

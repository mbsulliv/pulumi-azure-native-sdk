// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Hunt in Azure Security Insights.
type Hunt struct {
	pulumi.CustomResourceState

	// A list of mitre attack tactics the hunt is associated with
	AttackTactics pulumi.StringArrayOutput `pulumi:"attackTactics"`
	// A list of a mitre attack techniques the hunt is associated with
	AttackTechniques pulumi.StringArrayOutput `pulumi:"attackTechniques"`
	// The description of the hunt
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the hunt
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The hypothesis status of the hunt.
	HypothesisStatus pulumi.StringPtrOutput `pulumi:"hypothesisStatus"`
	// List of labels relevant to this hunt
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes a user that the hunt is assigned to
	Owner HuntOwnerResponsePtrOutput `pulumi:"owner"`
	// The status of the hunt.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHunt registers a new resource with the given unique name, arguments, and options.
func NewHunt(ctx *pulumi.Context,
	name string, args *HuntArgs, opts ...pulumi.ResourceOption) (*Hunt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	if args.HypothesisStatus == nil {
		args.HypothesisStatus = pulumi.StringPtr("Unknown")
	}
	if args.Status == nil {
		args.Status = pulumi.StringPtr("New")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Hunt"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:Hunt"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:Hunt"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:Hunt"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:Hunt"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:Hunt"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:Hunt"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:Hunt"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Hunt
	err := ctx.RegisterResource("azure-native:securityinsights/v20230701preview:Hunt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHunt gets an existing Hunt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHunt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HuntState, opts ...pulumi.ResourceOption) (*Hunt, error) {
	var resource Hunt
	err := ctx.ReadResource("azure-native:securityinsights/v20230701preview:Hunt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hunt resources.
type huntState struct {
}

type HuntState struct {
}

func (HuntState) ElementType() reflect.Type {
	return reflect.TypeOf((*huntState)(nil)).Elem()
}

type huntArgs struct {
	// A list of mitre attack tactics the hunt is associated with
	AttackTactics []string `pulumi:"attackTactics"`
	// A list of a mitre attack techniques the hunt is associated with
	AttackTechniques []string `pulumi:"attackTechniques"`
	// The description of the hunt
	Description string `pulumi:"description"`
	// The display name of the hunt
	DisplayName string `pulumi:"displayName"`
	// The hunt id (GUID)
	HuntId *string `pulumi:"huntId"`
	// The hypothesis status of the hunt.
	HypothesisStatus *string `pulumi:"hypothesisStatus"`
	// List of labels relevant to this hunt
	Labels []string `pulumi:"labels"`
	// Describes a user that the hunt is assigned to
	Owner *HuntOwner `pulumi:"owner"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The status of the hunt.
	Status *string `pulumi:"status"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Hunt resource.
type HuntArgs struct {
	// A list of mitre attack tactics the hunt is associated with
	AttackTactics pulumi.StringArrayInput
	// A list of a mitre attack techniques the hunt is associated with
	AttackTechniques pulumi.StringArrayInput
	// The description of the hunt
	Description pulumi.StringInput
	// The display name of the hunt
	DisplayName pulumi.StringInput
	// The hunt id (GUID)
	HuntId pulumi.StringPtrInput
	// The hypothesis status of the hunt.
	HypothesisStatus pulumi.StringPtrInput
	// List of labels relevant to this hunt
	Labels pulumi.StringArrayInput
	// Describes a user that the hunt is assigned to
	Owner HuntOwnerPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The status of the hunt.
	Status pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (HuntArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*huntArgs)(nil)).Elem()
}

type HuntInput interface {
	pulumi.Input

	ToHuntOutput() HuntOutput
	ToHuntOutputWithContext(ctx context.Context) HuntOutput
}

func (*Hunt) ElementType() reflect.Type {
	return reflect.TypeOf((**Hunt)(nil)).Elem()
}

func (i *Hunt) ToHuntOutput() HuntOutput {
	return i.ToHuntOutputWithContext(context.Background())
}

func (i *Hunt) ToHuntOutputWithContext(ctx context.Context) HuntOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HuntOutput)
}

type HuntOutput struct{ *pulumi.OutputState }

func (HuntOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hunt)(nil)).Elem()
}

func (o HuntOutput) ToHuntOutput() HuntOutput {
	return o
}

func (o HuntOutput) ToHuntOutputWithContext(ctx context.Context) HuntOutput {
	return o
}

// A list of mitre attack tactics the hunt is associated with
func (o HuntOutput) AttackTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringArrayOutput { return v.AttackTactics }).(pulumi.StringArrayOutput)
}

// A list of a mitre attack techniques the hunt is associated with
func (o HuntOutput) AttackTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringArrayOutput { return v.AttackTechniques }).(pulumi.StringArrayOutput)
}

// The description of the hunt
func (o HuntOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The display name of the hunt
func (o HuntOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Etag of the azure resource
func (o HuntOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The hypothesis status of the hunt.
func (o HuntOutput) HypothesisStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringPtrOutput { return v.HypothesisStatus }).(pulumi.StringPtrOutput)
}

// List of labels relevant to this hunt
func (o HuntOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o HuntOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Describes a user that the hunt is assigned to
func (o HuntOutput) Owner() HuntOwnerResponsePtrOutput {
	return o.ApplyT(func(v *Hunt) HuntOwnerResponsePtrOutput { return v.Owner }).(HuntOwnerResponsePtrOutput)
}

// The status of the hunt.
func (o HuntOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o HuntOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Hunt) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o HuntOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Hunt) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HuntOutput{})
}

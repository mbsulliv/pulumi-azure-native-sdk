// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Cognitive Services RaiPolicy.
type RaiPolicy struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of Cognitive Services RaiPolicy.
	Properties RaiPolicyPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRaiPolicy registers a new resource with the given unique name, arguments, and options.
func NewRaiPolicy(ctx *pulumi.Context,
	name string, args *RaiPolicyArgs, opts ...pulumi.ResourceOption) (*RaiPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cognitiveservices:RaiPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RaiPolicy
	err := ctx.RegisterResource("azure-native:cognitiveservices/v20231001preview:RaiPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRaiPolicy gets an existing RaiPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRaiPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RaiPolicyState, opts ...pulumi.ResourceOption) (*RaiPolicy, error) {
	var resource RaiPolicy
	err := ctx.ReadResource("azure-native:cognitiveservices/v20231001preview:RaiPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RaiPolicy resources.
type raiPolicyState struct {
}

type RaiPolicyState struct {
}

func (RaiPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*raiPolicyState)(nil)).Elem()
}

type raiPolicyArgs struct {
	// The name of Cognitive Services account.
	AccountName string `pulumi:"accountName"`
	// Properties of Cognitive Services RaiPolicy.
	Properties *RaiPolicyProperties `pulumi:"properties"`
	// The name of the RaiPolicy associated with the Cognitive Services Account
	RaiPolicyName *string `pulumi:"raiPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RaiPolicy resource.
type RaiPolicyArgs struct {
	// The name of Cognitive Services account.
	AccountName pulumi.StringInput
	// Properties of Cognitive Services RaiPolicy.
	Properties RaiPolicyPropertiesPtrInput
	// The name of the RaiPolicy associated with the Cognitive Services Account
	RaiPolicyName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (RaiPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*raiPolicyArgs)(nil)).Elem()
}

type RaiPolicyInput interface {
	pulumi.Input

	ToRaiPolicyOutput() RaiPolicyOutput
	ToRaiPolicyOutputWithContext(ctx context.Context) RaiPolicyOutput
}

func (*RaiPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RaiPolicy)(nil)).Elem()
}

func (i *RaiPolicy) ToRaiPolicyOutput() RaiPolicyOutput {
	return i.ToRaiPolicyOutputWithContext(context.Background())
}

func (i *RaiPolicy) ToRaiPolicyOutputWithContext(ctx context.Context) RaiPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RaiPolicyOutput)
}

func (i *RaiPolicy) ToOutput(ctx context.Context) pulumix.Output[*RaiPolicy] {
	return pulumix.Output[*RaiPolicy]{
		OutputState: i.ToRaiPolicyOutputWithContext(ctx).OutputState,
	}
}

type RaiPolicyOutput struct{ *pulumi.OutputState }

func (RaiPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RaiPolicy)(nil)).Elem()
}

func (o RaiPolicyOutput) ToRaiPolicyOutput() RaiPolicyOutput {
	return o
}

func (o RaiPolicyOutput) ToRaiPolicyOutputWithContext(ctx context.Context) RaiPolicyOutput {
	return o
}

func (o RaiPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*RaiPolicy] {
	return pulumix.Output[*RaiPolicy]{
		OutputState: o.OutputState,
	}
}

// Resource Etag.
func (o RaiPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RaiPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource
func (o RaiPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RaiPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of Cognitive Services RaiPolicy.
func (o RaiPolicyOutput) Properties() RaiPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *RaiPolicy) RaiPolicyPropertiesResponseOutput { return v.Properties }).(RaiPolicyPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o RaiPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RaiPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o RaiPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RaiPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RaiPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RaiPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RaiPolicyOutput{})
}
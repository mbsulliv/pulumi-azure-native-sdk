// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The integration account partner.
// Azure REST API version: 2019-05-01. Prior API version in Azure Native 1.x: 2019-05-01.
//
// Other available API versions: 2015-08-01-preview.
type IntegrationAccountPartner struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The partner content.
	Content PartnerContentResponseOutput `pulumi:"content"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The partner type.
	PartnerType pulumi.StringOutput `pulumi:"partnerType"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountPartner registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountPartner(ctx *pulumi.Context,
	name string, args *IntegrationAccountPartnerArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountPartner, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.PartnerType == nil {
		return nil, errors.New("invalid value for required argument 'PartnerType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountPartner"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountPartner"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationAccountPartner
	err := ctx.RegisterResource("azure-native:logic:IntegrationAccountPartner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountPartner gets an existing IntegrationAccountPartner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountPartner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountPartnerState, opts ...pulumi.ResourceOption) (*IntegrationAccountPartner, error) {
	var resource IntegrationAccountPartner
	err := ctx.ReadResource("azure-native:logic:IntegrationAccountPartner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountPartner resources.
type integrationAccountPartnerState struct {
}

type IntegrationAccountPartnerState struct {
}

func (IntegrationAccountPartnerState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountPartnerState)(nil)).Elem()
}

type integrationAccountPartnerArgs struct {
	// The partner content.
	Content PartnerContent `pulumi:"content"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The integration account partner name.
	PartnerName *string `pulumi:"partnerName"`
	// The partner type.
	PartnerType string `pulumi:"partnerType"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccountPartner resource.
type IntegrationAccountPartnerArgs struct {
	// The partner content.
	Content PartnerContentInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.Input
	// The integration account partner name.
	PartnerName pulumi.StringPtrInput
	// The partner type.
	PartnerType pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountPartnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountPartnerArgs)(nil)).Elem()
}

type IntegrationAccountPartnerInput interface {
	pulumi.Input

	ToIntegrationAccountPartnerOutput() IntegrationAccountPartnerOutput
	ToIntegrationAccountPartnerOutputWithContext(ctx context.Context) IntegrationAccountPartnerOutput
}

func (*IntegrationAccountPartner) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountPartner)(nil)).Elem()
}

func (i *IntegrationAccountPartner) ToIntegrationAccountPartnerOutput() IntegrationAccountPartnerOutput {
	return i.ToIntegrationAccountPartnerOutputWithContext(context.Background())
}

func (i *IntegrationAccountPartner) ToIntegrationAccountPartnerOutputWithContext(ctx context.Context) IntegrationAccountPartnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountPartnerOutput)
}

func (i *IntegrationAccountPartner) ToOutput(ctx context.Context) pulumix.Output[*IntegrationAccountPartner] {
	return pulumix.Output[*IntegrationAccountPartner]{
		OutputState: i.ToIntegrationAccountPartnerOutputWithContext(ctx).OutputState,
	}
}

type IntegrationAccountPartnerOutput struct{ *pulumi.OutputState }

func (IntegrationAccountPartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountPartner)(nil)).Elem()
}

func (o IntegrationAccountPartnerOutput) ToIntegrationAccountPartnerOutput() IntegrationAccountPartnerOutput {
	return o
}

func (o IntegrationAccountPartnerOutput) ToIntegrationAccountPartnerOutputWithContext(ctx context.Context) IntegrationAccountPartnerOutput {
	return o
}

func (o IntegrationAccountPartnerOutput) ToOutput(ctx context.Context) pulumix.Output[*IntegrationAccountPartner] {
	return pulumix.Output[*IntegrationAccountPartner]{
		OutputState: o.OutputState,
	}
}

// The changed time.
func (o IntegrationAccountPartnerOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

// The partner content.
func (o IntegrationAccountPartnerOutput) Content() PartnerContentResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) PartnerContentResponseOutput { return v.Content }).(PartnerContentResponseOutput)
}

// The created time.
func (o IntegrationAccountPartnerOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The resource location.
func (o IntegrationAccountPartnerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The metadata.
func (o IntegrationAccountPartnerOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// Gets the resource name.
func (o IntegrationAccountPartnerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The partner type.
func (o IntegrationAccountPartnerOutput) PartnerType() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.StringOutput { return v.PartnerType }).(pulumi.StringOutput)
}

// The resource tags.
func (o IntegrationAccountPartnerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o IntegrationAccountPartnerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountPartner) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountPartnerOutput{})
}

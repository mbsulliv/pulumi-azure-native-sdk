// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A domain specific resource identifier.
type WebAppDomainOwnershipIdentifier struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppDomainOwnershipIdentifier registers a new resource with the given unique name, arguments, and options.
func NewWebAppDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, args *WebAppDomainOwnershipIdentifierArgs, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppDomainOwnershipIdentifier"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppDomainOwnershipIdentifier"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppDomainOwnershipIdentifier
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppDomainOwnershipIdentifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppDomainOwnershipIdentifier gets an existing WebAppDomainOwnershipIdentifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDomainOwnershipIdentifierState, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifier, error) {
	var resource WebAppDomainOwnershipIdentifier
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppDomainOwnershipIdentifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppDomainOwnershipIdentifier resources.
type webAppDomainOwnershipIdentifierState struct {
}

type WebAppDomainOwnershipIdentifierState struct {
}

func (WebAppDomainOwnershipIdentifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierState)(nil)).Elem()
}

type webAppDomainOwnershipIdentifierArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName *string `pulumi:"domainOwnershipIdentifierName"`
	// String representation of the identity.
	Id *string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppDomainOwnershipIdentifier resource.
type WebAppDomainOwnershipIdentifierArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName pulumi.StringPtrInput
	// String representation of the identity.
	Id pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppDomainOwnershipIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierArgs)(nil)).Elem()
}

type WebAppDomainOwnershipIdentifierInput interface {
	pulumi.Input

	ToWebAppDomainOwnershipIdentifierOutput() WebAppDomainOwnershipIdentifierOutput
	ToWebAppDomainOwnershipIdentifierOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierOutput
}

func (*WebAppDomainOwnershipIdentifier) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifier)(nil)).Elem()
}

func (i *WebAppDomainOwnershipIdentifier) ToWebAppDomainOwnershipIdentifierOutput() WebAppDomainOwnershipIdentifierOutput {
	return i.ToWebAppDomainOwnershipIdentifierOutputWithContext(context.Background())
}

func (i *WebAppDomainOwnershipIdentifier) ToWebAppDomainOwnershipIdentifierOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDomainOwnershipIdentifierOutput)
}

func (i *WebAppDomainOwnershipIdentifier) ToOutput(ctx context.Context) pulumix.Output[*WebAppDomainOwnershipIdentifier] {
	return pulumix.Output[*WebAppDomainOwnershipIdentifier]{
		OutputState: i.ToWebAppDomainOwnershipIdentifierOutputWithContext(ctx).OutputState,
	}
}

type WebAppDomainOwnershipIdentifierOutput struct{ *pulumi.OutputState }

func (WebAppDomainOwnershipIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppDomainOwnershipIdentifier)(nil)).Elem()
}

func (o WebAppDomainOwnershipIdentifierOutput) ToWebAppDomainOwnershipIdentifierOutput() WebAppDomainOwnershipIdentifierOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierOutput) ToWebAppDomainOwnershipIdentifierOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppDomainOwnershipIdentifier] {
	return pulumix.Output[*WebAppDomainOwnershipIdentifier]{
		OutputState: o.OutputState,
	}
}

// Kind of resource.
func (o WebAppDomainOwnershipIdentifierOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifier) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppDomainOwnershipIdentifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppDomainOwnershipIdentifierOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppDomainOwnershipIdentifier) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppDomainOwnershipIdentifierOutput{})
}

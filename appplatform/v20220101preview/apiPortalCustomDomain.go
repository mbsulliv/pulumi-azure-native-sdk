// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Custom domain of the API portal
type ApiPortalCustomDomain struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of custom domain for API portal
	Properties ApiPortalCustomDomainPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiPortalCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewApiPortalCustomDomain(ctx *pulumi.Context,
	name string, args *ApiPortalCustomDomainArgs, opts ...pulumi.ResourceOption) (*ApiPortalCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiPortalName == nil {
		return nil, errors.New("invalid value for required argument 'ApiPortalName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:ApiPortalCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiPortalCustomDomain
	err := ctx.RegisterResource("azure-native:appplatform/v20220101preview:ApiPortalCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiPortalCustomDomain gets an existing ApiPortalCustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiPortalCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiPortalCustomDomainState, opts ...pulumi.ResourceOption) (*ApiPortalCustomDomain, error) {
	var resource ApiPortalCustomDomain
	err := ctx.ReadResource("azure-native:appplatform/v20220101preview:ApiPortalCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiPortalCustomDomain resources.
type apiPortalCustomDomainState struct {
}

type ApiPortalCustomDomainState struct {
}

func (ApiPortalCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalCustomDomainState)(nil)).Elem()
}

type apiPortalCustomDomainArgs struct {
	// The name of API portal.
	ApiPortalName string `pulumi:"apiPortalName"`
	// The name of the API portal custom domain.
	DomainName *string `pulumi:"domainName"`
	// The properties of custom domain for API portal
	Properties *ApiPortalCustomDomainProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ApiPortalCustomDomain resource.
type ApiPortalCustomDomainArgs struct {
	// The name of API portal.
	ApiPortalName pulumi.StringInput
	// The name of the API portal custom domain.
	DomainName pulumi.StringPtrInput
	// The properties of custom domain for API portal
	Properties ApiPortalCustomDomainPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (ApiPortalCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalCustomDomainArgs)(nil)).Elem()
}

type ApiPortalCustomDomainInput interface {
	pulumi.Input

	ToApiPortalCustomDomainOutput() ApiPortalCustomDomainOutput
	ToApiPortalCustomDomainOutputWithContext(ctx context.Context) ApiPortalCustomDomainOutput
}

func (*ApiPortalCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalCustomDomain)(nil)).Elem()
}

func (i *ApiPortalCustomDomain) ToApiPortalCustomDomainOutput() ApiPortalCustomDomainOutput {
	return i.ToApiPortalCustomDomainOutputWithContext(context.Background())
}

func (i *ApiPortalCustomDomain) ToApiPortalCustomDomainOutputWithContext(ctx context.Context) ApiPortalCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalCustomDomainOutput)
}

type ApiPortalCustomDomainOutput struct{ *pulumi.OutputState }

func (ApiPortalCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalCustomDomain)(nil)).Elem()
}

func (o ApiPortalCustomDomainOutput) ToApiPortalCustomDomainOutput() ApiPortalCustomDomainOutput {
	return o
}

func (o ApiPortalCustomDomainOutput) ToApiPortalCustomDomainOutputWithContext(ctx context.Context) ApiPortalCustomDomainOutput {
	return o
}

// The name of the resource.
func (o ApiPortalCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of custom domain for API portal
func (o ApiPortalCustomDomainOutput) Properties() ApiPortalCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) ApiPortalCustomDomainPropertiesResponseOutput { return v.Properties }).(ApiPortalCustomDomainPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApiPortalCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ApiPortalCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiPortalCustomDomainOutput{})
}

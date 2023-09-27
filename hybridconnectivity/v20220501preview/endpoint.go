// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The endpoint for the target resource.
type Endpoint struct {
	pulumi.CustomResourceState

	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrOutput `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrOutput `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrOutput `pulumi:"lastModifiedByType"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource Id of the connectivity endpoint (optional).
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridconnectivity:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:hybridconnectivity/v20211006preview:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:hybridconnectivity/v20230315:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("azure-native:hybridconnectivity/v20220501preview:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure-native:hybridconnectivity/v20220501preview:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The endpoint name.
	EndpointName *string `pulumi:"endpointName"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// The resource Id of the connectivity endpoint (optional).
	ResourceId *string `pulumi:"resourceId"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri string `pulumi:"resourceUri"`
	// The type of endpoint.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrInput
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrInput
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrInput
	// The endpoint name.
	EndpointName pulumi.StringPtrInput
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrInput
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrInput
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrInput
	// The resource Id of the connectivity endpoint (optional).
	ResourceId pulumi.StringPtrInput
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri pulumi.StringInput
	// The type of endpoint.
	Type pulumi.StringInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

func (i *Endpoint) ToOutput(ctx context.Context) pulumix.Output[*Endpoint] {
	return pulumix.Output[*Endpoint]{
		OutputState: i.ToEndpointOutputWithContext(ctx).OutputState,
	}
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*Endpoint] {
	return pulumix.Output[*Endpoint]{
		OutputState: o.OutputState,
	}
}

// The timestamp of resource creation (UTC).
func (o EndpointOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o EndpointOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o EndpointOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o EndpointOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o EndpointOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o EndpointOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o EndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource provisioning state.
func (o EndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource Id of the connectivity endpoint (optional).
func (o EndpointOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}

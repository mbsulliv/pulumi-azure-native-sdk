// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The entity query kind
type EntityTimelineKind string

const (
	// activity
	EntityTimelineKindActivity = EntityTimelineKind("Activity")
	// bookmarks
	EntityTimelineKindBookmark = EntityTimelineKind("Bookmark")
	// security alerts
	EntityTimelineKindSecurityAlert = EntityTimelineKind("SecurityAlert")
	// anomaly
	EntityTimelineKindAnomaly = EntityTimelineKind("Anomaly")
)

func (EntityTimelineKind) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityTimelineKind)(nil)).Elem()
}

func (e EntityTimelineKind) ToEntityTimelineKindOutput() EntityTimelineKindOutput {
	return pulumi.ToOutput(e).(EntityTimelineKindOutput)
}

func (e EntityTimelineKind) ToEntityTimelineKindOutputWithContext(ctx context.Context) EntityTimelineKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EntityTimelineKindOutput)
}

func (e EntityTimelineKind) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return e.ToEntityTimelineKindPtrOutputWithContext(context.Background())
}

func (e EntityTimelineKind) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return EntityTimelineKind(e).ToEntityTimelineKindOutputWithContext(ctx).ToEntityTimelineKindPtrOutputWithContext(ctx)
}

func (e EntityTimelineKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityTimelineKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityTimelineKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntityTimelineKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EntityTimelineKindOutput struct{ *pulumi.OutputState }

func (EntityTimelineKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityTimelineKind)(nil)).Elem()
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindOutput() EntityTimelineKindOutput {
	return o
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindOutputWithContext(ctx context.Context) EntityTimelineKindOutput {
	return o
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return o.ToEntityTimelineKindPtrOutputWithContext(context.Background())
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntityTimelineKind) *EntityTimelineKind {
		return &v
	}).(EntityTimelineKindPtrOutput)
}

func (o EntityTimelineKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EntityTimelineKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityTimelineKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EntityTimelineKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityTimelineKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityTimelineKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EntityTimelineKindPtrOutput struct{ *pulumi.OutputState }

func (EntityTimelineKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityTimelineKind)(nil)).Elem()
}

func (o EntityTimelineKindPtrOutput) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return o
}

func (o EntityTimelineKindPtrOutput) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return o
}

func (o EntityTimelineKindPtrOutput) Elem() EntityTimelineKindOutput {
	return o.ApplyT(func(v *EntityTimelineKind) EntityTimelineKind {
		if v != nil {
			return *v
		}
		var ret EntityTimelineKind
		return ret
	}).(EntityTimelineKindOutput)
}

func (o EntityTimelineKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityTimelineKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EntityTimelineKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EntityTimelineKindInput is an input type that accepts values of the EntityTimelineKind enum
// A concrete instance of `EntityTimelineKindInput` can be one of the following:
//
//	EntityTimelineKindActivity
//	EntityTimelineKindBookmark
//	EntityTimelineKindSecurityAlert
//	EntityTimelineKindAnomaly
type EntityTimelineKindInput interface {
	pulumi.Input

	ToEntityTimelineKindOutput() EntityTimelineKindOutput
	ToEntityTimelineKindOutputWithContext(context.Context) EntityTimelineKindOutput
}

var entityTimelineKindPtrType = reflect.TypeOf((**EntityTimelineKind)(nil)).Elem()

type EntityTimelineKindPtrInput interface {
	pulumi.Input

	ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput
	ToEntityTimelineKindPtrOutputWithContext(context.Context) EntityTimelineKindPtrOutput
}

type entityTimelineKindPtr string

func EntityTimelineKindPtr(v string) EntityTimelineKindPtrInput {
	return (*entityTimelineKindPtr)(&v)
}

func (*entityTimelineKindPtr) ElementType() reflect.Type {
	return entityTimelineKindPtrType
}

func (in *entityTimelineKindPtr) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return pulumi.ToOutput(in).(EntityTimelineKindPtrOutput)
}

func (in *entityTimelineKindPtr) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EntityTimelineKindPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EntityTimelineKindOutput{})
	pulumi.RegisterOutputType(EntityTimelineKindPtrOutput{})
}

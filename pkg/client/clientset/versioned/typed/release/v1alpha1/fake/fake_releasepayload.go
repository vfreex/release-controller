// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openshift/release-controller/pkg/apis/release/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReleasePayloads implements ReleasePayloadInterface
type FakeReleasePayloads struct {
	Fake *FakeReleaseV1alpha1
	ns   string
}

var releasepayloadsResource = schema.GroupVersionResource{Group: "release.openshift.io", Version: "v1alpha1", Resource: "releasepayloads"}

var releasepayloadsKind = schema.GroupVersionKind{Group: "release.openshift.io", Version: "v1alpha1", Kind: "ReleasePayload"}

// Get takes name of the releasePayload, and returns the corresponding releasePayload object, and an error if there is any.
func (c *FakeReleasePayloads) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ReleasePayload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(releasepayloadsResource, c.ns, name), &v1alpha1.ReleasePayload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReleasePayload), err
}

// List takes label and field selectors, and returns the list of ReleasePayloads that match those selectors.
func (c *FakeReleasePayloads) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReleasePayloadList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(releasepayloadsResource, releasepayloadsKind, c.ns, opts), &v1alpha1.ReleasePayloadList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReleasePayloadList{ListMeta: obj.(*v1alpha1.ReleasePayloadList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReleasePayloadList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested releasePayloads.
func (c *FakeReleasePayloads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(releasepayloadsResource, c.ns, opts))

}

// Create takes the representation of a releasePayload and creates it.  Returns the server's representation of the releasePayload, and an error, if there is any.
func (c *FakeReleasePayloads) Create(ctx context.Context, releasePayload *v1alpha1.ReleasePayload, opts v1.CreateOptions) (result *v1alpha1.ReleasePayload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(releasepayloadsResource, c.ns, releasePayload), &v1alpha1.ReleasePayload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReleasePayload), err
}

// Update takes the representation of a releasePayload and updates it. Returns the server's representation of the releasePayload, and an error, if there is any.
func (c *FakeReleasePayloads) Update(ctx context.Context, releasePayload *v1alpha1.ReleasePayload, opts v1.UpdateOptions) (result *v1alpha1.ReleasePayload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(releasepayloadsResource, c.ns, releasePayload), &v1alpha1.ReleasePayload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReleasePayload), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReleasePayloads) UpdateStatus(ctx context.Context, releasePayload *v1alpha1.ReleasePayload, opts v1.UpdateOptions) (*v1alpha1.ReleasePayload, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(releasepayloadsResource, "status", c.ns, releasePayload), &v1alpha1.ReleasePayload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReleasePayload), err
}

// Delete takes name of the releasePayload and deletes it. Returns an error if one occurs.
func (c *FakeReleasePayloads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(releasepayloadsResource, c.ns, name), &v1alpha1.ReleasePayload{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReleasePayloads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(releasepayloadsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReleasePayloadList{})
	return err
}

// Patch applies the patch and returns the patched releasePayload.
func (c *FakeReleasePayloads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReleasePayload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(releasepayloadsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReleasePayload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReleasePayload), err
}

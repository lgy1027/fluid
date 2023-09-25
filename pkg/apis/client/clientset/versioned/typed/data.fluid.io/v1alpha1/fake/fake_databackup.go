/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataBackups implements DataBackupInterface
type FakeDataBackups struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var databackupsResource = schema.GroupVersionResource{Group: "data.fluid.io", Version: "v1alpha1", Resource: "databackups"}

var databackupsKind = schema.GroupVersionKind{Group: "data.fluid.io", Version: "v1alpha1", Kind: "DataBackup"}

// Get takes name of the dataBackup, and returns the corresponding dataBackup object, and an error if there is any.
func (c *FakeDataBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(databackupsResource, c.ns, name), &v1alpha1.DataBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataBackup), err
}

// List takes label and field selectors, and returns the list of DataBackups that match those selectors.
func (c *FakeDataBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(databackupsResource, databackupsKind, c.ns, opts), &v1alpha1.DataBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataBackupList{ListMeta: obj.(*v1alpha1.DataBackupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataBackups.
func (c *FakeDataBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(databackupsResource, c.ns, opts))

}

// Create takes the representation of a dataBackup and creates it.  Returns the server's representation of the dataBackup, and an error, if there is any.
func (c *FakeDataBackups) Create(ctx context.Context, dataBackup *v1alpha1.DataBackup, opts v1.CreateOptions) (result *v1alpha1.DataBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(databackupsResource, c.ns, dataBackup), &v1alpha1.DataBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataBackup), err
}

// Update takes the representation of a dataBackup and updates it. Returns the server's representation of the dataBackup, and an error, if there is any.
func (c *FakeDataBackups) Update(ctx context.Context, dataBackup *v1alpha1.DataBackup, opts v1.UpdateOptions) (result *v1alpha1.DataBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(databackupsResource, c.ns, dataBackup), &v1alpha1.DataBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataBackups) UpdateStatus(ctx context.Context, dataBackup *v1alpha1.DataBackup, opts v1.UpdateOptions) (*v1alpha1.DataBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(databackupsResource, "status", c.ns, dataBackup), &v1alpha1.DataBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataBackup), err
}

// Delete takes name of the dataBackup and deletes it. Returns an error if one occurs.
func (c *FakeDataBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(databackupsResource, c.ns, name, opts), &v1alpha1.DataBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(databackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataBackupList{})
	return err
}

// Patch applies the patch and returns the patched dataBackup.
func (c *FakeDataBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(databackupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataBackup), err
}

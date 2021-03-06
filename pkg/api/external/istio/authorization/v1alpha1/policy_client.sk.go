// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type PolicyWatcher interface {
	// watch namespace-scoped Policies
	Watch(namespace string, opts clients.WatchOpts) (<-chan PolicyList, <-chan error, error)
}

type PolicyClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Policy, error)
	Write(resource *Policy, opts clients.WriteOpts) (*Policy, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (PolicyList, error)
	PolicyWatcher
}

type policyClient struct {
	rc clients.ResourceClient
}

func NewPolicyClient(rcFactory factory.ResourceClientFactory) (PolicyClient, error) {
	return NewPolicyClientWithToken(rcFactory, "")
}

func NewPolicyClientWithToken(rcFactory factory.ResourceClientFactory, token string) (PolicyClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Policy{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Policy resource client")
	}
	return NewPolicyClientWithBase(rc), nil
}

func NewPolicyClientWithBase(rc clients.ResourceClient) PolicyClient {
	return &policyClient{
		rc: rc,
	}
}

func (client *policyClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *policyClient) Register() error {
	return client.rc.Register()
}

func (client *policyClient) Read(namespace, name string, opts clients.ReadOpts) (*Policy, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Policy), nil
}

func (client *policyClient) Write(policy *Policy, opts clients.WriteOpts) (*Policy, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(policy, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Policy), nil
}

func (client *policyClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *policyClient) List(namespace string, opts clients.ListOpts) (PolicyList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToPolicy(resourceList), nil
}

func (client *policyClient) Watch(namespace string, opts clients.WatchOpts) (<-chan PolicyList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	policiesChan := make(chan PolicyList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				policiesChan <- convertToPolicy(resourceList)
			case <-opts.Ctx.Done():
				close(policiesChan)
				return
			}
		}
	}()
	return policiesChan, errs, nil
}

func convertToPolicy(resources resources.ResourceList) PolicyList {
	var policyList PolicyList
	for _, resource := range resources {
		policyList = append(policyList, resource.(*Policy))
	}
	return policyList
}

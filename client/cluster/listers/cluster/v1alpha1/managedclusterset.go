// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/open-cluster-management/api/cluster/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagedClusterSetLister helps list ManagedClusterSets.
type ManagedClusterSetLister interface {
	// List lists all ManagedClusterSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedClusterSet, err error)
	// Get retrieves the ManagedClusterSet from the index for a given name.
	Get(name string) (*v1alpha1.ManagedClusterSet, error)
	ManagedClusterSetListerExpansion
}

// managedClusterSetLister implements the ManagedClusterSetLister interface.
type managedClusterSetLister struct {
	indexer cache.Indexer
}

// NewManagedClusterSetLister returns a new ManagedClusterSetLister.
func NewManagedClusterSetLister(indexer cache.Indexer) ManagedClusterSetLister {
	return &managedClusterSetLister{indexer: indexer}
}

// List lists all ManagedClusterSets in the indexer.
func (s *managedClusterSetLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedClusterSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedClusterSet))
	})
	return ret, err
}

// Get retrieves the ManagedClusterSet from the index for a given name.
func (s *managedClusterSetLister) Get(name string) (*v1alpha1.ManagedClusterSet, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedclusterset"), name)
	}
	return obj.(*v1alpha1.ManagedClusterSet), nil
}

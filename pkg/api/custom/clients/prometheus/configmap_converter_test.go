package prometheus_test

import (
	"context"

	supergloov1 "github.com/solo-io/supergloo/pkg/api/v1"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kubev1 "k8s.io/api/core/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/configmap"
	"github.com/solo-io/supergloo/test/inputs"

	. "github.com/solo-io/supergloo/pkg/api/custom/clients/prometheus"
	v1 "github.com/solo-io/supergloo/pkg/api/external/prometheus/v1"
)

var _ = Describe("ConfigmapConverter", func() {
	c := NewPrometheusConfigmapConverter()
	fakeResourceClient, _ := configmap.NewResourceClient(nil, &v1.PrometheusConfig{}, nil, false)
	It("converts a prometheus configmap to a v1.PrometheusConfig", func() {
		configMap := inputs.PrometheusConfigMap("myname", "mynamespace")
		promCfg, err := c.FromKubeConfigMap(context.TODO(), fakeResourceClient, configMap)
		Expect(err).NotTo(HaveOccurred())
		Expect(promCfg).To(Equal(inputs.PrometheusConfig(promCfg.GetMetadata().Name, promCfg.GetMetadata().Namespace)))
	})
	It("returns nil/nil for a non-prometheus configmap", func() {
		configMap := &kubev1.ConfigMap{}
		promCfg, err := c.FromKubeConfigMap(context.TODO(), fakeResourceClient, configMap)
		Expect(err).NotTo(HaveOccurred())
		Expect(promCfg).To(BeNil())
	})
	It("returns err for anything other than a v1.PrometheusConfig", func() {
		_, err := c.ToKubeConfigMap(context.TODO(), fakeResourceClient, &supergloov1.Mesh{})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("*v1.Mesh not type *v1.PrometheusConfig cannot convert"))
	})
	It("converts a prometheus config resource to a v1.PrometheusConfig", func() {
		promCfg := inputs.PrometheusConfig("foo", "bar")
		promConfigMap, err := c.ToKubeConfigMap(context.TODO(), fakeResourceClient, promCfg)
		Expect(err).NotTo(HaveOccurred())
		yam, err := yaml.Marshal(promCfg.Config)
		Expect(err).NotTo(HaveOccurred())
		Expect(promConfigMap).To(Equal(&kubev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar", Annotations: map[string]string{"cfg_key_name": "prometheus.yml"}},
			Data:       map[string]string{"prometheus.yml": string(yam)},
		}))
	})
})

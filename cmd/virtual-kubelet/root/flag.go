// Copyright © 2017 The virtual-kubelet authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package root

import (
	"flag"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

func InstallFlags(flags *pflag.FlagSet, c *Opts) {
	flags.StringVar(&c.HomeKubeconfig, "home-kubeconfig", c.HomeKubeconfig, "kube config file to use for connecting to the Kubernetes API server")
	flags.StringVar(&c.NodeName, "nodename", c.NodeName, "kubernetes node name")
	flags.StringVar(&c.MetricsAddr, "metrics-addr", c.MetricsAddr, "address to listen for metrics/stats requests")

	flags.UintVar(&c.PodWorkers, "pod-reflection-workers", c.PodWorkers, "the number of pod reflection workers")
	flags.UintVar(&c.ServiceWorkers, "service-reflection-workers", c.ServiceWorkers, "the number of service reflection workers")
	flags.UintVar(&c.EndpointSliceWorkers, "endpointslice-reflection-workers", c.EndpointSliceWorkers,
		"the number of endpointslice reflection workers")
	flags.UintVar(&c.ConfigMapWorkers, "configmap-reflection-workers", c.ConfigMapWorkers, "the number of configmap reflection workers")
	flags.UintVar(&c.SecretWorkers, "secret-reflection-workers", c.SecretWorkers, "the number of secret reflection workers")
	flags.UintVar(&c.PersistenVolumeClaimWorkers, "persistentvolumeclaim-reflection-workers", c.PersistenVolumeClaimWorkers,
		"the number of persistentvolumeclaim reflection workers")

	flags.DurationVar(&c.InformerResyncPeriod, "full-resync-period", c.InformerResyncPeriod,
		"how often to perform a full resync of pods between kubernetes and the provider")
	flags.DurationVar(&c.LiqoInformerResyncPeriod, "liqo-resync-period", c.LiqoInformerResyncPeriod,
		"how often to perform a full resync of Liqo resources informers")
	flags.DurationVar(&c.StartupTimeout, "startup-timeout", c.StartupTimeout, "How long to wait for the virtual-kubelet to start")

	flags.StringVar(&c.ForeignClusterID, "foreign-cluster-id", c.ForeignClusterID, "The Id of the foreign cluster")
	flags.StringVar(&c.KubeletNamespace, "kubelet-namespace", c.KubeletNamespace, "The namespace of the virtual kubelet")
	flags.StringVar(&c.HomeClusterID, "home-cluster-id", c.HomeClusterID, "The Id of the home cluster")
	flags.StringVar(&c.LiqoIpamServer, "ipam-server", c.LiqoIpamServer, "The server the Virtual Kubelet should "+
		"connect to in order to contact the IPAM module")
	flags.BoolVar(&c.Profiling, "enable-profiling", c.Profiling, "Enable pprof profiling")

	flags.Var(&c.NodeExtraAnnotations, "node-extra-annotations", "Extra annotations to add to the Virtual Node")
	flags.Var(&c.NodeExtraLabels, "node-extra-labels", "Extra labels to add to the Virtual Node")

	flags.BoolVar(&c.EnableStorage, "enable-storage", false, "Enable the Liqo storage reflection")
	flags.StringVar(&c.VirtualStorageClassName, "virtual-storage-class-name", "liqo", "Name of the virtual storage class")
	flags.StringVar(&c.RemoteRealStorageClassName, "remote-real-storage-class-name", "", "Name of the real storage class to use for the actual volumes")

	flagset := flag.NewFlagSet("klog", flag.PanicOnError)
	klog.InitFlags(flagset)
	flagset.VisitAll(func(f *flag.Flag) {
		f.Name = "klog." + f.Name
		flags.AddGoFlag(f)
	})
}

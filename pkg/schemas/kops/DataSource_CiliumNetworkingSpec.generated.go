package schemas

import (
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func DataSourceCiliumNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version":                           ComputedString(),
			"access_log":                        ComputedString(),
			"agent_labels":                      ComputedList(String()),
			"agent_prometheus_port":             ComputedInt(),
			"allow_localhost":                   ComputedString(),
			"auto_ipv_6node_routes":             ComputedBool(),
			"bpf_root":                          ComputedString(),
			"container_runtime":                 ComputedList(String()),
			"container_runtime_endpoint":        ComputedMap(String()),
			"debug":                             ComputedBool(),
			"debug_verbose":                     ComputedList(String()),
			"device":                            ComputedString(),
			"disable_conntrack":                 ComputedBool(),
			"disable_ipv_4":                     ComputedBool(),
			"disable_k8s_services":              ComputedBool(),
			"enable_policy":                     ComputedString(),
			"enable_tracing":                    ComputedBool(),
			"enable_prometheus_metrics":         ComputedBool(),
			"envoy_log":                         ComputedString(),
			"ipv_4cluster_cidr_mask_size":       ComputedInt(),
			"ipv_4node":                         ComputedString(),
			"ipv_4range":                        ComputedString(),
			"ipv_4service_range":                ComputedString(),
			"ipv_6cluster_alloc_cidr":           ComputedString(),
			"ipv_6node":                         ComputedString(),
			"ipv_6range":                        ComputedString(),
			"ipv_6service_range":                ComputedString(),
			"k8s_api_server":                    ComputedString(),
			"k8s_kubeconfig_path":               ComputedString(),
			"keep_bpf_templates":                ComputedBool(),
			"keep_config":                       ComputedBool(),
			"label_prefix_file":                 ComputedString(),
			"labels":                            ComputedList(String()),
			"lb":                                ComputedString(),
			"lib_dir":                           ComputedString(),
			"log_drivers":                       ComputedList(String()),
			"log_opt":                           ComputedMap(String()),
			"logstash":                          ComputedBool(),
			"logstash_agent":                    ComputedString(),
			"logstash_probe_timer":              ComputedInt(),
			"disable_masquerade":                ComputedBool(),
			"nat_46range":                       ComputedString(),
			"pprof":                             ComputedBool(),
			"prefilter_device":                  ComputedString(),
			"prometheus_serve_addr":             ComputedString(),
			"restore":                           ComputedBool(),
			"single_cluster_route":              ComputedBool(),
			"socket_path":                       ComputedString(),
			"state_dir":                         ComputedString(),
			"trace_payload_len":                 ComputedInt(),
			"tunnel":                            ComputedString(),
			"enable_ipv_6":                      ComputedBool(),
			"enable_ipv_4":                      ComputedBool(),
			"monitor_aggregation":               ComputedString(),
			"bpfct_global_tcp_max":              ComputedInt(),
			"bpfct_global_any_max":              ComputedInt(),
			"preallocate_bpf_maps":              ComputedBool(),
			"sidecar_istio_proxy_image":         ComputedString(),
			"cluster_name":                      ComputedString(),
			"to_fqdns_dns_reject_response_code": ComputedString(),
			"to_fqdns_enable_poller":            ComputedBool(),
			"container_runtime_labels":          ComputedString(),
			"ipam":                              ComputedString(),
			"ip_tables_rules_noinstall":         ComputedBool(),
			"auto_direct_node_routes":           ComputedBool(),
			"enable_node_port":                  ComputedBool(),
			"etcd_managed":                      ComputedBool(),
			"enable_remote_node_identity":       ComputedBool(),
			"remove_cbr_bridge":                 ComputedBool(),
			"restart_pods":                      ComputedBool(),
			"reconfigure_kubelet":               ComputedBool(),
			"node_init_bootstrap_file":          ComputedString(),
			"cni_bin_path":                      ComputedString(),
		},
	}
}

func ExpandDataSourceCiliumNetworkingSpec(in map[string]interface{}) kops.CiliumNetworkingSpec {
	if in == nil {
		panic("expand CiliumNetworkingSpec failure, in is nil")
	}
	return kops.CiliumNetworkingSpec{
		Version: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["version"]),
		AccessLog: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["access_log"]),
		AgentLabels: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["agent_labels"]),
		AgentPrometheusPort: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["agent_prometheus_port"]),
		AllowLocalhost: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["allow_localhost"]),
		AutoIpv6NodeRoutes: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["auto_ipv_6node_routes"]),
		BPFRoot: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["bpf_root"]),
		ContainerRuntime: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["container_runtime"]),
		ContainerRuntimeEndpoint: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["container_runtime_endpoint"]),
		Debug: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["debug"]),
		DebugVerbose: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["debug_verbose"]),
		Device: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["device"]),
		DisableConntrack: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_conntrack"]),
		DisableIpv4: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_ipv_4"]),
		DisableK8sServices: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_k8s_services"]),
		EnablePolicy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["enable_policy"]),
		EnableTracing: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_tracing"]),
		EnablePrometheusMetrics: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_prometheus_metrics"]),
		EnvoyLog: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["envoy_log"]),
		Ipv4ClusterCIDRMaskSize: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["ipv_4cluster_cidr_mask_size"]),
		Ipv4Node: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_4node"]),
		Ipv4Range: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_4range"]),
		Ipv4ServiceRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_4service_range"]),
		Ipv6ClusterAllocCidr: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6cluster_alloc_cidr"]),
		Ipv6Node: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6node"]),
		Ipv6Range: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6range"]),
		Ipv6ServiceRange: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipv_6service_range"]),
		K8sAPIServer: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["k8s_api_server"]),
		K8sKubeconfigPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["k8s_kubeconfig_path"]),
		KeepBPFTemplates: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["keep_bpf_templates"]),
		KeepConfig: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["keep_config"]),
		LabelPrefixFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["label_prefix_file"]),
		Labels: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["labels"]),
		LB: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["lb"]),
		LibDir: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["lib_dir"]),
		LogDrivers: func(in interface{}) []string {
			return func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["log_drivers"]),
		LogOpt: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
		}(in["log_opt"]),
		Logstash: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["logstash"]),
		LogstashAgent: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["logstash_agent"]),
		LogstashProbeTimer: func(in interface{}) uint32 {
			return uint32(ExpandInt(in))
		}(in["logstash_probe_timer"]),
		DisableMasquerade: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["disable_masquerade"]),
		Nat46Range: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["nat_46range"]),
		Pprof: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["pprof"]),
		PrefilterDevice: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["prefilter_device"]),
		PrometheusServeAddr: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["prometheus_serve_addr"]),
		Restore: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["restore"]),
		SingleClusterRoute: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["single_cluster_route"]),
		SocketPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["socket_path"]),
		StateDir: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["state_dir"]),
		TracePayloadLen: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["trace_payload_len"]),
		Tunnel: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tunnel"]),
		EnableIpv6: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_ipv_6"]),
		EnableIpv4: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_ipv_4"]),
		MonitorAggregation: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["monitor_aggregation"]),
		BPFCTGlobalTCPMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpfct_global_tcp_max"]),
		BPFCTGlobalAnyMax: func(in interface{}) int {
			return int(ExpandInt(in))
		}(in["bpfct_global_any_max"]),
		PreallocateBPFMaps: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["preallocate_bpf_maps"]),
		SidecarIstioProxyImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["sidecar_istio_proxy_image"]),
		ClusterName: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cluster_name"]),
		ToFqdnsDNSRejectResponseCode: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["to_fqdns_dns_reject_response_code"]),
		ToFqdnsEnablePoller: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["to_fqdns_enable_poller"]),
		ContainerRuntimeLabels: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["container_runtime_labels"]),
		Ipam: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["ipam"]),
		IPTablesRulesNoinstall: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["ip_tables_rules_noinstall"]),
		AutoDirectNodeRoutes: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["auto_direct_node_routes"]),
		EnableNodePort: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_node_port"]),
		EtcdManaged: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["etcd_managed"]),
		EnableRemoteNodeIdentity: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["enable_remote_node_identity"]),
		RemoveCbrBridge: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["remove_cbr_bridge"]),
		RestartPods: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["restart_pods"]),
		ReconfigureKubelet: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["reconfigure_kubelet"]),
		NodeInitBootstrapFile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["node_init_bootstrap_file"]),
		CniBinPath: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["cni_bin_path"]),
	}
}

func FlattenDataSourceCiliumNetworkingSpecInto(in kops.CiliumNetworkingSpec, out map[string]interface{}) {
	out["version"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Version)
	out["access_log"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AccessLog)
	out["agent_labels"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AgentLabels)
	out["agent_prometheus_port"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.AgentPrometheusPort)
	out["allow_localhost"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AllowLocalhost)
	out["auto_ipv_6node_routes"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AutoIpv6NodeRoutes)
	out["bpf_root"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.BPFRoot)
	out["container_runtime"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.ContainerRuntime)
	out["container_runtime_endpoint"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.ContainerRuntimeEndpoint)
	out["debug"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Debug)
	out["debug_verbose"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.DebugVerbose)
	out["device"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Device)
	out["disable_conntrack"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableConntrack)
	out["disable_ipv_4"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableIpv4)
	out["disable_k8s_services"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableK8sServices)
	out["enable_policy"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EnablePolicy)
	out["enable_tracing"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableTracing)
	out["enable_prometheus_metrics"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnablePrometheusMetrics)
	out["envoy_log"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.EnvoyLog)
	out["ipv_4cluster_cidr_mask_size"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.Ipv4ClusterCIDRMaskSize)
	out["ipv_4node"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipv4Node)
	out["ipv_4range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipv4Range)
	out["ipv_4service_range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipv4ServiceRange)
	out["ipv_6cluster_alloc_cidr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipv6ClusterAllocCidr)
	out["ipv_6node"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipv6Node)
	out["ipv_6range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipv6Range)
	out["ipv_6service_range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipv6ServiceRange)
	out["k8s_api_server"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.K8sAPIServer)
	out["k8s_kubeconfig_path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.K8sKubeconfigPath)
	out["keep_bpf_templates"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.KeepBPFTemplates)
	out["keep_config"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.KeepConfig)
	out["label_prefix_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LabelPrefixFile)
	out["labels"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Labels)
	out["lb"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LB)
	out["lib_dir"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LibDir)
	out["log_drivers"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.LogDrivers)
	out["log_opt"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.LogOpt)
	out["logstash"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Logstash)
	out["logstash_agent"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.LogstashAgent)
	out["logstash_probe_timer"] = func(in uint32) interface{} {
		return FlattenInt(int(in))
	}(in.LogstashProbeTimer)
	out["disable_masquerade"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.DisableMasquerade)
	out["nat_46range"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Nat46Range)
	out["pprof"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Pprof)
	out["prefilter_device"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PrefilterDevice)
	out["prometheus_serve_addr"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.PrometheusServeAddr)
	out["restore"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Restore)
	out["single_cluster_route"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.SingleClusterRoute)
	out["socket_path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SocketPath)
	out["state_dir"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.StateDir)
	out["trace_payload_len"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.TracePayloadLen)
	out["tunnel"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Tunnel)
	out["enable_ipv_6"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableIpv6)
	out["enable_ipv_4"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableIpv4)
	out["monitor_aggregation"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MonitorAggregation)
	out["bpfct_global_tcp_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFCTGlobalTCPMax)
	out["bpfct_global_any_max"] = func(in int) interface{} {
		return FlattenInt(int(in))
	}(in.BPFCTGlobalAnyMax)
	out["preallocate_bpf_maps"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.PreallocateBPFMaps)
	out["sidecar_istio_proxy_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SidecarIstioProxyImage)
	out["cluster_name"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ClusterName)
	out["to_fqdns_dns_reject_response_code"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ToFqdnsDNSRejectResponseCode)
	out["to_fqdns_enable_poller"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.ToFqdnsEnablePoller)
	out["container_runtime_labels"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ContainerRuntimeLabels)
	out["ipam"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Ipam)
	out["ip_tables_rules_noinstall"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.IPTablesRulesNoinstall)
	out["auto_direct_node_routes"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.AutoDirectNodeRoutes)
	out["enable_node_port"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableNodePort)
	out["etcd_managed"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EtcdManaged)
	out["enable_remote_node_identity"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.EnableRemoteNodeIdentity)
	out["remove_cbr_bridge"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.RemoveCbrBridge)
	out["restart_pods"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.RestartPods)
	out["reconfigure_kubelet"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.ReconfigureKubelet)
	out["node_init_bootstrap_file"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.NodeInitBootstrapFile)
	out["cni_bin_path"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.CniBinPath)
}

func FlattenDataSourceCiliumNetworkingSpec(in kops.CiliumNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceCiliumNetworkingSpecInto(in, out)
	return out
}

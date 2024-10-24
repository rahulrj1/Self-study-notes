CKA Course Notes

Installing k8s means installing the following things: API server, etcd, kubelet, Container runtime, scheduler & kube-controller-manager.

Kubelet and container runtime are there on the worker nodes & rest all are there on the master node

Earlier k8s supported only docker. Then, k8s introduced an interface CRI (Container Runtime Interface) which allowed any vendor to work as a container runtime for Kubernetes, as long as they adhere to the OCI standards. OCI stands for Open Container Initiative & it consists of an imagespec (how an image should be built) and a runtimespec (how any container runtime should be developed).  When you run a kubectl command, the kubectl utility is reching to the kube-apiserver. The kube-api server first authenticates the request and validates it. It then retrieves data from the etcd cluster and responds back with the requested information.
You don’t really need to use the kubectl command line. Instead you could also invoke the APIs directly by sending a POST request like this.  
// What happens when you create a pod When we create a pod, the request is first authenticated and then validated. The API server creates a pod object without assigning it to a node, updates the information in the etcd cluster, updates the user that the pod has been created. The scheduler continuously monitors the API server and notices that there is a new pod with no nodes assigned. The scheduler then identifies the right node to place the new pod on and communicates that back to the kube-apiserver. The api server then updates the information in the etcd cluster. The api server then passes the information to the Kubelet in the appropriate worker node. The Kubelet then creates the pod on the node and instructs the container runtime engine to deploy the application image. Once done, the Kubelet updates the information back to the api server and the api server then updates the data back in the etcd cluster.  A controller is a process that monitors the state of various components within the system and works towards bringing the whole system to the desired functioning state. Node controller, replication controller are two examples of controllers. 

Scheduler is responsible for scheduling pods on nodes. It is only responsible for deciding which pod goes on which node. It doesn’t actually place the pod on the node. That’s the job of the Kubelet.

// kube-proxy
Within a Kubernetes cluster, every pod can reach every other pod. This is accomplished by deploying a pod networking solution to the cluster. A pod network is an internal virtual network that spans across all the nodes in the cluster to which all the pods connect to. The service can’t join the pod network because the service is not an actual thing. It doesn’t have any interface or an actively listening process. It is a virtual component that lives only in the k8s memory. Kube-proxy is a process that runs on each node in the k8s cluster. Its job is to look for new services and every time a new service is created, it creates the appropriate rule on each node to forward traffic to those services to the backend pods. One way it does this is by using the iptables rules. It creates an iptable rule on each node in the cluster to forward traffic heading to the IP of the service to the IP of the actual pod. 
// Namespace Namespace is a way to divide cluster resources between multiple users (or applications) within the same k8s cluster. The resources within a namespace can refer to each other simply by their names
For pod to access resource of another namespace, use “servicename.namespace.svc.cluster.local”
When a service is created, a DNS entry is added automatically in this format.
cluster.local is the default domain name of the k8s cluster. svc is the subdomain for the service
To create a resource always in a particular namespace, add “namespace:[namespace]” in the metadata of yaml file of that resource.
Namespace can be created using a yaml file.
“Kubectl create namespace dev” To limit resources in a namespace, create a resourceQuota 
// Scheduling The easiest way to schedule a pod is to simply set the nodeName field to the name of the node in pod specs file while creating the pod.
Another way to assign a node to an existing pod is to create a binding object and send a POST request to the pod’s binding API.

Taints and tolerations are used to set restrictions on what pods can be scheduled on a node. Suppose node 1 has dedicated resources for a particular use case or application. So, we would like only that pods that belong to those application to be placed on node 1. First, we prevent all the pods from being placed on the node by placing a taint on the node. By default, pods have no tolerations. We then add a toleration to the pod that we want to be placed on node 1.
“kubectl taint node node-name [key]=[value]:taint-effect” There are three taint effects -> NoSchedule (pods won’t be scheduled), PreferNoSchedule (system will try to avoid placing the pods but not guarenteed), NoExecute (new pods won’t be scheduled and existing pods will be evicted) Toleration of a pod goes in spec of pod.yaml life tolerations:
	- key: “[key]”
	   operator: “Equal”
	   value: “[value]”
	   effect: “[taint-effect]” 
NodeSelector is used to restrict pod to be scheduled on a node with a particular label only. For advanced options like or/not, we use node affinity
“Kubectl label nodes node-1 size=large” The primary purpose of Node affinity feature is to ensure pods are hosted on particular nodes.  DaemonSet ensures that one copy of a pod runs on every worker node in the k8s cluster. It can be used for kube-proxy. Another use case is for networking. Networking solutions like weave net requires an agent to be deployed on each node in the cluster.  Static pods are directly created by Kubelet without any intervention from the kube-apiserver. kubeadm tool uses static pods to set up the k8s cluster. It runs control plane components as pods in the kube-system namespace.   “kubectl top node” is the command to view cluster performance.  // Networking  If we connect two systems (computers) A and B to a switch, then the switch creates a network containing two systems. To connect to a switch, we need an interface on each host, physical or virtual, depending on the host.  “ip link” command is used to see (list/modify) the interfaces for the host. 
“ip addr” command is used to see IP address assigned to those interfaces.
“ip addr add <IP_address>/<subnet_mask> dev <interface_name>” command is used in linux systems to add an IP address to a network interface.  <IP_address> is the IP address you want to assign to the interface. “ip route” command is used to view the routing table. 
“ip route add” command is used to add entries to the routing table. 
“cat /proc/sys/net/ipv4/ip_forward” command is used to check if IP forwarding is enabled on a host. Also change in the /etc/sysctl.conf file to persist changes across reboots.
 “ip addr add <192.168.1.10 (IP of the network)>/<24 (subnet_mask)> dev <interface>” command we use to then assign the systems with IP addresses on same network. 
Changes made using these commands are only valid till restart. If we want to persist these changes, we must set them in the /etc/network/interfaces file.

// Router A router helps connect two networks together.
If a router is used to connect 2 separate networks together, then it gets 2 IPs assigned, one on each network.  
When one system in 1st network ties to communicate with other system in 2nd network, how does it know where to send the packet. The router is just another device in a network. That’s where gateway/route comes into the picture.
“ip route add <2nd-network-swithc-ip> via <router-ip-1>” this command adds a route so  that we can reach 2nd network via router-ip-1 route.
“Ip route” command is used to see the existing routing configurations on the system  // DNS
/etc/hosts file contains information about Domain names and their IP addresses. Translating hostnames to IP address this way is known as name resolution.
Earlier all hosts had list of these domain names and ip addresses. Later, all these entries were moved to a single server who managed it centrally. We call it DNS (Domain Name Server) Every host has a DNS resolution configuration file at /etc/resolv.conf We add an entry into it specifying the address of the DNS server. “nameserver <IP_address_of_DNS>” 
nslookup ignores the /etc/hosts file

// Network Namespace
To create an internal bridge network, we add a new interface to the host using “ip link add” command with the type set to bridge
“ip netns add <net-ns-name>” command creates a network namespace
“ip netns” command lists all the network namespaces
“ip netns exec <net-ns-name> ip link” / “ip -n <net-ns-name> link” lists all the interfaces of the network namespace
 
We can connect two network namespaces using a virtual ethernet pair, or a virtual cable with two interfaces on either ends. 
“ip link add <interface-1> type veth peer name <interface-2>” command is used to create a virtual cable
“ip link set <interface-name> nets <net-ns-name>” command is used to attach interface to the appropriate namespace. We do the same for 2nd interface.
“ip -n <net-ns-name> link del <interface-1>” command is used to delete the virtual cable. When we delete one end, the other end gets deleted automatically. 
Then, we can assign IP addresses to each of these namespaces. “ip -n <net-ns-name> addr add <ip-address> dev <interface-name>“ command will assign IP address to the interface in the network namespace
“ip -n <net-ns-name> link set <interface-name> up” command is used to bring up the interface

The above was for connecting 2 network namespaces. For more than 2 net namespaces, we create a virtual network using a virtual switch. To create a virtual switch in a host, there are multiple solutions available such as Linux bridge.  “ip link add <interface-name> type bridge” command is used to create an internal bridge network. 
“ip link set dev <interface-name> up” command is used to turn it up
For host, it is just another interface. For network namespaces, this interface is like a switch to which it can connect to. It is an interface for the host but a switch for the network namespaces.
The next step is to connect the namespaces to this virtual network switch. We create cables to add namespaces to the bridge.
“ip link add <interface1(veth-red)> type veth peer name <interface-2(veth-red-br)>” command is used to create a virtual cable
“ip link set <interface1(veth-red)> netns <net-ns-name(red)>” command is used to attach one interface of this to the red namespace 
“ip link set <interface2(veth-red-br)> master <vitual-switch-interface(v-net-0)>” command is used to attach other end to the bridge network
We then set up IP address for these links and then set them up.
We can then assign IP address to the virtual-switch-interface in order to access any of the network namespace from the host. 
NAT rule is used to forward the traffic coming to one port to another port

// Docker networking
When docker is installed on the host, it creates an internal private network called bridge by default. Docker calls the network by the name bridge but on the host the network is created by the name docker0 (it is present as an interface & acts as switch for the network namespaces of docker).  // Ingress
Ingress helps your users access your application using a single externally accessible URL that you can configure to route to different services within your cluster based on the URL path, at the same time, implement SSL security as well. 
For ingress, we need ingress-controller and ingress-resource
Ingress-Controller: Deployment, Service, ConfigMap & ServiceAccount Ingress-Resource: Create a yaml file of type ingress
We have rules at the top for each domain name 
// Yaml file of an ingress resource apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: example-ingress
spec:
  rules:
  -host: example.com
    http:
      paths:
      - path: /app1
        pathType: Prefix
        backend:
          serviceName: app1-service
          servicePort: 80
      - path: /app2
        pathType: Prefix
        backend:
          serviceName: app2-service
          servicePort: 80

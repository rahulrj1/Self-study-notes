Kubernetes for Absolute Beginner Notes

Volume is a kind of permanent storage in docker that persists even when container is deleted
Node -> A machine (physical or virtual) on which Kubernetes is installed
Cluster -> A set of nodes grouped together
API Server -> act as the frontend for k8s. Users, Management devices, CLI, all talk to API server to interact with the k8s cluster
etcd -> A distributed reliable key value store used by k8s to store all data used to manage the cluster
Scheduler -> responsible for distributing work or container across multiple nodes
Controller -> brain behind orchestration. They are responsible for noticing and responding when nodes, containers or end point goes down.
Container Runtime -> Underlying software that is used to run containers (Docker in our case)
Kubelet -> agent that runs on each node in the cluster. Responsible for making sure that the containers are running on the nodes as expected.  Kubectl -> this tool is used to deploy and manage applications on a k8s cluster
“kubectl run” command is used to deploy an application on the cluster. It deploys a docker container by creating a pod.
“kubectl cluster-info” command is used to view information about the cluster
“kubectl get nodes” command is used to list all the nodes part of the cluster

Pod -> a single instance of an application. It is inside a node. The smallest object that you can create in k8s. To scale up, you create new pods and to scale down, you delete existing. You don’t add additional containers to an existing pod to scale your application. Pods don’t have 2 identical containers in it. But one pod can have a helper container with a main container. They share the same network namespace (can communicate via localhost) and they can easily share the same storage space as well.

“kubectl run nginx1 --image nginx” => nginx1 is the name provided to the resource being created. --image nginx specifies the docker image to use for creating the containerised app. The command creates a pod. The image is downloaded from the docker hub repository
“kubectl get pods” => list of all the pods

Deployment: A Deployment manages a set of identical Pods, ensuring that the desired state (specified by a user) is maintained. Deployments are responsible for scaling, updating, and managing the lifecycle of Pods.

Yaml file
Key value => 		fruit: apple
Array => 			fruits:	
		   			- 	apple
		   			- 	mango
Dictionary/map => banana:
						calories: 105
						fat: 0.4 g
List/array are ordered whereas dictionary are unordered


A k8s definition file always contains 4 top level fields: apiVersion, kind, metadata (name & label [dictionary]), spec

Example of YAML file for creating a pod is apiVersion: v1
kind: Pod
metadata:
  name: my-pod
  labels:
    app: my-app
spec:
  containers:
  - name: my-container
    image: nginx:latest
    ports:
    - containerPort: 80

Replication Controller is a controller that creates & maintains n replicas of pods. Its spec will look like  spec:
	template:
		// metadata of pod
		// spec of pod
	replicas: 3

ReplicaSet is very similar to Replication Controller
Yaml for replicaSet will have apiVersion: apps/v1 and rest everything same as replication controller. Also, it has selector in spec

Spec:
	template:
	replicas: 3
	selector: 
		matchLabel: 
			// same as label of pod

To change the number of replicas, we can use either of the 2 commands:
“kubectl scale —replicas=6 -f [replicaset-file].yml”
“kubectl scale —replicas=6 replicaset [replicaset-name]”

Deployment.yaml file has everything same as replicaset.yaml except kind.
Deployment => (creates) ReplicaSet => (creates) pods

When you create a deployment, it triggers a rollout. A new rollout creates a new deployment revision.
“kubectl rollout status [deployment-name]” to see the status of rollout
“kubectl rollout history [deployment-name]”

There are 2 deployment strategies: recreate (destroy and recreate at once) and rolling update (one by one destroy and recreate)

When you upgrade your deployment, it creates a new replicaset
“kubectl rollout undo [deployment/deployment-name]” to undo the rollout

Unlike in the docker world where an IP address is always assigned to a container, in k8s IP address is assigned to a pod.

K8s services enable communication between various components within and outside of the application. It helps us connect applications together with other applications or users. 
NodePort service => One use case is to listen to a port on the node and forward that request to a port on the pod running the web application. This type of service is NodePort service.
Cluster IP service => the service creates a virtual IP inside the cluster to enable communication between different services.
LoadBalancer service => It provisions a load balancer for our application.

There are 3 IPs in the service 1. Target Port -> port of the pod
2.Port -> port of the service object
3.Node Port -> port of the node

For service,
Spec:	
	type: NodePort
	ports:
		- targetPort: 80
		   port: 80
		   nodePort: 30008
	selector: 
		// labels of pod

If there are 3 pods with the same label, then the service balances the incoming request onto all the 3 pods.
If there are 3 nodes with 1 pod each, then we can access our app by using the ip of any node/NodePort. The service maps nodePort of each node to the targetPort of desired pods. 
If we have 4 nodes in the cluster: 2 for voting app pods & 2 for result app pods. Each (Voting/Result) Pods will be accessible on IPs of all the 4 nodes.

ClusterIP service: In k8s, clusterIP is a type of service that exposes an internal service on a cluster-internal IP. This means that the service is accessible only within the k8s cluster. It is used for creating a common service for a group of pods, say backend so that frontend pods can communicate to the backend pods using service name.

Link is a command line option which can be used to link two containers together  eg, “docker run -d --name=vote -p 5000:80 —link [container-name]:[alias-host-name] voting-app”  kubeadmin tool can be used to bootstrap a Kubernetes cluster. It helps us set up a Multi-node cluster using k8s best practices

Steps to create a Kubernetes cluster: 
1. We need to have multiple systems or VMs provisioned. We need a couples of machines which can be physical or virtual machines
2. Install container runtime on the nodes
3. Install kubeadm on all the nodes
4. Initialise the master server
5. Prior to joining cluster nodes to the cluster, we must ensure that network prerequisites are met. Normal network connectivity is not enough for this. K8s requires a special networking solution between the master and the worker nodes called the POD Network
6. Join worker nodes to the master node
7. Deploy our application onto the Kubernetes environment  

In k8s, a namespace is a way to divide cluster resources between multiple users, multiple teams or to partition resources in some other way. 

DaemonSet is used to create pods such that one copy of pod exists in each node.
 A ConfigMap is an API object used to store non-confidential data in key-value pairs. Scope of ConfigMap is only upto the namespace. Environment variables defined in the yaml file of a pod has scope only upto that pod

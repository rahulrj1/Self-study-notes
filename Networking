Networking Notes

There are 4 distinct networking problems:
1. Container-to-container: via localhost and containerPort
2. Pod-to-pod: Every pod has a unique IP address which is reachable from all the other pods in the k8s cluster.
3. Pod-to-service
4. External-to-service 
Network plugin is configured to assign IP addresses to Pods
Kube-apiserver is configured to assign IP addresses to Services
Kubelet or cloud controller manager is configured to assign IP addresses to Nodes. 

Cluster networking types:
1. IPv4 only
2. IPv6 only
3. IPv4/IPv6 Dual stack

CNI stands for Container Network Interface. The primary role of CNI is to establish network connectivity between containers and pods running on different nodes within a Kubernetes cluster. Real world analogy: Imagine a city with various neighbourhoods and buildings. The road infrastructure represents the underlay network. And the postal service represents the CNI. Different postal services may operate in the same city, just as different CNI plugins may be used within the same Kubernetes cluster. 
Multus is a Kubernetes CNI plugin that enables the support for multiple network interfaces on pods. By default, k8s supports a single network interface per pod.
A network interface in the context of k8s refers to a connection point on a node or a pod through which it communicates with other entities on a network. It’s like a gateway or a doorway that allows information to flow in and out of a container or a server within a k8s cluster.  Real world analogy: In k8s, a pod typically has one network interface which is analogous to some house which is having only a single door.  // Underlay and overlay networks Underlay Network: The underlay network refers to the physical network infrastructure that connects the nodes(servers) within a data centre. It is responsible for providing basic connectivity between the physical hardware, such as switches, routers and servers. Real world analogy: Think of the underlay network as the roads, highways, and the streets in a city. 
Overlay Network: The overlay network is a virtual network layer built on top of the underlay network. It provides logical connectivity between containers, pods, or VMs running on different nodes within a k8s cluster. If we use a digital map to navigate the streets in city then this digital map is analogous to the overlay network.
In summary, the underlay network provides the physical infrastructure for data transmission within a data-center or cloud environment, while the overlay network creates a logical network layer that enables communication between distributed entities like containers or pods. Together, they ensure efficient and reliable communication in complex distributed systems like Kubernetes clusters.

// Traditional Networking vs SDN
In traditional networking, data is transmitted via routers. Router has two components: Control plane & Data forwarding plane. Control plane does the computational task for routing (eg, to decide data should be forwarded to which router) & Data forwarding plane does the transfer of data packets.
SDN simply removes the control plane task of the routers. Routing decisions are now taken care by the SDN controller in the SDN.

// Physical function refers to a specific capability or task performed by a physical hardware component within a system.  NAT stands for Network Address Translation. It is used to translate internal IP address into external IP address when pods try to communicate with resources outside the cluster.    Switch is the device we use for computers in the same environment to communicate with each other. Switch use cables. If we want to do it using wireless technology, we can use access point device. Suppose that if seven computers are connected to a switch in a home/office then they all constitute a LAN (Local Area Network). A switch have many ports and they are called as LAN ports.  Router is a device required for a computer to connect to the internet. The main task of a router is to enable computers to connect to the internet. In our previous LAN setup of 7 computers and a switch, we connect the switch to a router via a cable. We now need a connection between the router and the internet. A special cable comes to our houses/offices which is given to us by ISP (Internet Service Provider) and we connect to the internet by using this cable.
The router is a door to access the internet.  We can also think of router as a device we use to communicate with a computer in the other part of the world (or in a different LAN). We can think of internet as a structure that connects all the LAN all over the world.  Home router = Switch + Router Each router must have a special table called as routing table. This table tells us which route the packet should choose. Each router has a special processor and the information in the routing table is created by this special processor using many different algorithms.  WAN (Wide Area Network) -> combination of LANs
Suppose there are two offices of a company in two different parts of the world. And we want the two LANs of each office to communicate with each other. We can use a WAN in this situation. The two LANs can also communicate via Internet but Internet is public network and for security reasons, the company should use WAN. The most popular and cost effective method of setting up WAN is by using VPN. VPN tunnelling provides privacy, anonymity and security to us by creating a special network connection over a public network. We use switches to create LAN while we use routers to create WAN.  Trick for OSI Layer: All People Should Try New Dominos Pizza  DNS -> Domain Name System. It resolves Domain names to IP addresses. When you type yahoo.com in your web browser and if your web browser or operating system can’t find the IP address in its own cache memory, it will send the query to the next level, to resolver server. The resolver server is basically your ISP. Then the resolver on receiving the query will check its own cache memory to find an IP address for yahoo.com and if it can’t find it will send the query to the next level which is the root server. The root servers are the top or root of the DNS hierarchy. 13 sets of this root servers are strategically placed around the world operated by 12 different organisations. Each set has their own unique IP address. When the root server receives a query for the IP address of yahoo.com, it doesn’t know what the IP address is but the root server know where to send the resolver to help it find the IP address. Root server will direct the resolver to TLD (top level domain) server for the .com domain. The TLD server stores the address information for a top level domain, such as .com, .net, .org and so on. This particular TLD manages the .com domain which yahoo.com is a part of. The TLD server also doesn’t know the IP address. It redirects the request to Authoritative Name Server. It will respond with the IP address of the yahoo.com to the resolver server which then sends the IP address to your computer.  Web Browser -> Resolver Server (ISP) -> Root Server -> TLD (.com) -> Authoritative Name Server (It knows the IP address) -> Resolver Server (ISP) -> Web Browser

Load Balancer: A Load balancer is a device or software that distributes network or application traffic across multiple servers. It can be at level 4 (transport) or level 7 (application) of OSI model. 
Server Load Balancer: SLB refers to a load balancer designed to distribute incoming network traffic across multiple servers at the server level. It operates at layer 4 of the OSI model. LB operates at network level whereas SLB operates at server level.   Global Server Load Balancer: GSLB is an advanced form of load balancing that operates across multiple locations or data centres. GSLB can distribute traffic across multiple geographically dispersed locations.  Requirements a global load balancer:
1. Ability to load balance over LoadBalancer Services and Ingresses
2. Sophisticated load balancing policies
3. Support for health checks and ability to remove unhealthy endpoints from the load balanced pool
There are two architectural approaches for designing a global load balancer: DNS-based and Anycast-based.  Ingress is an API object that exposes HTTP routes from outside the cluster to services within the cluster.

Data Link Layer is subdivided into further two layers. Upper half is LLC (Logical Link Control) and lower half is MAC
Data	Application
Data	Presentation
Data	Session
Segment 	Transport
Packet	Network
Frame	DataLink
Bit	Physical
OSI Model was never implemented but TCP/IP model was implemented

TCP/IP Model
Application Layer 
Transport Layer
Internet Layer
Network Interface Layer  MAC address has six bytes. The first three bytes are OUIs and are assigned by IEEE. MAC address is associated with a specific NIC card. We can’t change it but we can spoof it. 
Ethernet is a family of standards that applies to physical and logical aspects of the LAN. It is a network protocol that controls how data is transmitted over a LAN  Network Interface Card -> If you have a device that you want to be able to communicate on a network, you need to have at least one NIC

Hub & Switch both are central connecting devices. Hub is a dumb device and Switch is a smart device. Hub will broadcast every message to all the connected devices. Switch memorises the MAC address of every device connected to it via using a MAC address table, also called a CAM (Content Addressable Memory) table. Switches have ASIC (Application Specific Integrated Circuitry) which helps to do this.

WAP (Wireless Access Point) -> It is a bridge that extends the wired network to the wireless network. WAP isn’t a router. Routers break up a broadcast domain.

Modem (Modulators/Demodulators) modulate one signal to another (eg. analog to digital)  Firewall is a fundamental IT security device. Firewall can be network-based (hardware) or host-based (software). It protects our network from malicious activity on the internet.  DHCP stands for Dynamic Host Configuration Protocol and we utilise a DHCP server to automatically assign IP address configurations to devices on a network. It uses UDP (User Datagram Protocol).  IP address configurations mean: IP address, Subnet mask, Default gateway, DNS server

Application layer serves as gateway for user applications to access network services and resources. It defines the protocols and interfaces used by applications to communicate over the network
Presentation layer takes care of data representation & encryption.
Session layer is responsible for setting up, managing and then tearing down sessions between network devices. 

CSMA (Carrier Sense Multiple Access) and Token ring are two specific types of network access methodologies. In token ring, only device with the token can send the data  ARP (Address Resolution Protocol): Its purpose is to resolve addresses. Specifically, IP address to MAC address.  
IP is a protocol that defines routing and logical addressing of packets that allows data to travel WANs and the internet. It specifies the formatting of packets and the logical addressing schema. It is connectionless. 
ICMP is an error control companion protocol to the IP. It is used to generate error messages to the source IP address when network issues prevent delivery of a packet. Ping is one of the command line utilities that leverages the ICMP protocol. 
Protocols are rules governing how machines exchange data. A protocol runs as a process or service in an operating system. Protocols always have a port (logical port)number associated to them.  Ports are logical constructs that bind a unique port number to a protocol service or process Sockets are a combination of an IP address and a port number.   NTP (Network Time Protocol) is a protocol that automatically synchronises a system’s time with a network time server. It uses TCP.   Telnet is a protocol used to connect to remote host.  SSH is a cryptographic protocol that’s used to securely connect to a remote host. It encrypts data with PKI (public key infrastructure) 
Telnet and SSH allow us to get a terminal connection to a remote host.  RDP (Remote Desktop Protocol) is a protocol that allows us to get a GUI to remote system.   SMTP (Simple Mail Transfer Protocol) is a protocol that is used to send an email from an email client to an email server
POP3 (Post Office Protocol) is a protocol that is used to retrieve emails from an email server
IMAP (Internet Message Access Protocol)
With POP3, emails aren’t synced in multiple devices whereas in IMAP, emails are synced in multiple devices.  HTTP is a protocol that provides browsing services for the World Wide Web (web browsers and web servers). SSL (Secure Socket Layer)/TLS (Transport Layer Security) is used to encrypt HTTP content.

An IPv4 has 2 portions: a network portion and a host portion
 NAT (Network Address Translation) allows us to have private IP addresses on our internal network and at our NAT device, we have a single (or a handful of) public IP address. So, they allow us to map private IP addresses to public IP addresses  Routing are of two types: static and dynamic There are three types of dynamic routing protocols: distance-vector, link-state & hybrid Interior Gateway Protocol (IGP) are routing protocols within a single autonomous system Exterior Gateway Protocol (EGP) are used to route traffic from one AS to another. BGP is used as EGP by almost all ISP.  Routing table entries: directly connected routes (network that are directly connected to the router), remote network routes, default routes  The standard for wireless networking is IEEE 802.11 Any device that is connected to a wireless network, it’s going to have a wireless network interface card and at least one antenna that creates radio waves

Docker for absolute beginners notes

Containers are completely isolated environments. They can have their own processes, services, network interfaces, mounts just like VMs, except they all share the same OS kernel.
Docker utilises LXC containers.

OS consists of two things: an OS kernel (responsible for interacting with the underlying hardware) and a set of softwares. Docker can run any flavour of OS on top of it as long as they all are based on the same kernel.
The main purpose of docker is to package and containerise apps and to ship them and to run them anywhere any times 

A container only lives as long as the process inside it is alive.

Basic Docker commands:
“docker run” command is used to run a container from an image. “docker run nginx” command will run an instance of nginx application on the docker host if it already exists. If the image is not present on the host it will go out to docker host and pull that image down. This is only done the first time. “docker ps” command list all the running containers and some basic information about them.
“docker ps -a” command list all the running as well as previously stopped or exited containers.
“docker stop” command will stop the running container but we must provide either the container Id or the container name.
“docker rm” command is used to remove a stopped command permanently.
“docker images” command will list all the available images on our host.
“docker rmi” command is used to remove images from our host. You must stop and delete all the dependent containers to be able to delete an image.
“docker pull” command is used to only pull the image and not run the container
“docker exec [container-name] [command]” command is used to execute a command on a running docker container
"docker run -d” will run the docker container in the background mode
“docker attach” command is used to attach back to the background running container 
“docker run -p [host-port]:[container-port]”

By default, the docker container doesn’t listen to a standard input. It doesn’t have a terminal to read input from. If you would like to provide your input, you must map the standard input of your host to the docker container using the “-I” parameter. It stands for interactive mode. “-t” stands for pseudo terminal. “-it” option is used overall for input as well as prompt (terminal as well as interactive mode)  Every docker container gets an IP assigned by default 

// volume 
For data persistence, we can map a directory outside the container on the docker host to a directory inside the container
“docker run -v [docker-host-dir]:[container-dir] mysql” 
“docker inspect [container-name]” -> Details of a container in docker format

Steps to create your own image:
1. Create a Dockerfile
2. Run the “docker build . -f Dockerfile -t [image-name (/rahul/my-custom-image)]” command
This will create an image locally on your system.
To make it available to the public docker hub registry, 
3. “docker push [image-name]” command

// Docker image creation
Dockerfile is a tex file written in a specific format that docker can understand. It’s an [instruction, argument] format. 

Sample Dockerfile is:

FROM Ubuntu 			// Start from a base image or another image
RUN apt-get update && apt-get -y install python
RUN pip install flask flask-mysql
COPY . /opt/source-code		// copy files from the local system onto the docker image ENTRYPOINT FLASK_APP=/opt/source-code/app.py flask run  			
// entry point allows us to specify a command that will be run when the image is run as a container

Docker builds the images in a layered architecture. All the layers are cached by the docker. Only the layers after the updated layer needs to be rebuild

// environment variable
“docker run -e [env_var_name]=[env_var_value] [container_name]”   Suppose we have an ubuntu-sleeper image whose Dockerfile looks like:
FROM Ubuntu
ENTRYPOINT [“sleep”]
CMD [“5”]

If we run the command “docker run ubuntu-sleeper 10” then 10 will get appended to the ENTRYPOINT command so the command that will be run when the container starts is sleep 10. In case of CMD command, the command line parameters will get replaced entirely whereas in case of ENTRYPOINT command, the command line parameters will get appended.   // docker-compose file
If we need to run a complex application using multiple services, a better way than running various containers one-by-one is to use docker-compose
Then we can run “docker-compose up” command to bring up the entire application stack
This is only applicable to running containers on a single docker host

Link is a command line option which can be used to link two containers together  eg, “docker run -d --name=vote -p 5000:80 —link [container-name]:[alias-host-name] voting-app”

Our app structure: Voting-app -> redis -> worker -> Postgres -> result app

To run the above, we use the following commands:
docker run -d --name=redis redis
docker run -d --name=db postgres:9.4
docker run -d --name=vote -p 5000:80 --link redis:redis voting-app docker run -d --name=result -p 5001:80 --link db:db result-app
docker run -d --name=worker --link db:db --link redis:redis worker 
Our docker-compose.yaml will look like:
redis:
	image: redis
db: 	image: postgres:9.4 vote: 	image: voting-app
	ports:	
		- 5000:80
	links:
		- redis result:
	image: result-app
	ports:
		- 5001:80
	links:
		- db
worker:
	image: worker
	links: 
		- redis
		- db  Run the “docker-compose up” command to bring up the entire application stack

If we would like to instruct docker-compose to run a docker build instead of trying to pull an image, we could replace “image:voting-app” with “build ./vote” where ./vote is the location of the directory which contains the application code and a Dockerfile with instructions to build the docker image.

For docker-compose file version 2 and above:
version: 2
services:
	// here are the container info

Another difference is with networking. In version 1, docker-compose attaches all the containers it runs to default bridge network and then use link to enable communication between the containers.  With version 2, docker-compose automatically creates a dedicated bridge network for this application and then attaches all the containers to that new network. Container are able communicate using each other’s service name. So, we don’t need to use links.
There is also a depends on feature in version 2 and above. This is used to create a startup order.

If we would like to separate the user generated traffic from the app’s internal traffic, then we can create two networks: a front-end network & a back-end network. We can then connect the voting and result app to the frontend network & all the apps to the backend network.

docker-compose.yaml version: 2
services:
	redis:
		//
		networks:
			- back-end
	vote:
		//
		networks:
			- front-end
			- back-end
	//
networks:
	front-end:
	back-end:


// Docker registry
If the containers were the rain, then they would rain from the docker registry which are the clouds. It is a central repository of all docker images.

// Private docker registry
When you have applications built in-house that shouldn’t be made available to the public, hosting an internal private registry may be a good solution.  When you are running your application on premise and don’t have a private registry, how do you deploy your own private registry within your organisation:	Docker registry is itself another application and is available as a docker image. Run the docker registry image as a container. Use the “docker image tag” command to tag the image with a private registry URL in it. “Docker push” command pushes the image to the local private registry.  docker engine refers to a docker host with docker installed on it. Installing docker on linux host means installing three different components: 
1. Docker Daemon - The background process that manages docker objects such as images, containers, volumes and networks.
2. REST API server - The API interface that programs can use to talk to the Daemon and provide instructions.
3. Docker CLI - The command line interface. It used the REST API to interact with the Docker Daemon.
Docker CLI need not necessarily be on the same host. 
Docker uses namespaces to isolate workspaces. Process IDs, Network, InterProcess communication, mount & Unix timesharing systems are created in their own namespace, thereby providing isolation between containers.  Suppose our Dockerfile consists of 5 steps then during the “docker build” command, docker will create 5 image layers which will be read only. Then, when we run “docker run” command, docker builds a 6th new writable layer on top of the image layer. The writable layer is used to store data created by the container such as log files by the applications, any temporary files generated by the container or just any file modified by the user on that container. The life of this layer is only as long as the container is alive.  The same image layer is shared by all the containers created using this image.  // Docker volume
“docker volume create data_volume” command is used to create data_volume directory inside the /var/lib/docker/volumes directory
“docker run -v data_volume:/var/lib/mysql mysql” command is used to run the mysql container along with docker mounting the volume data_volume to the container. This is called volume mounting. 
“docker run -v [full_path_to_folder]:/var/lib/mysql mysql” -> This is called bind mounting.  --mount option is used nowadays instead of -v
"docker run --mount type=bind, source=[full_path_to_folder], target=/var/lib/mysql mysql”  Docker uses storage driver to enable layered architecture.

// Docker networking
Docker automatically creates three networks: 
1. Bridge -> the default network a container gets attached to. A private internal network created by docker on the host. All containers get attached to this network by default and they get an internal IP address. Containers can access each other using this internal IP. To access any of these containers from the outside world, map the ports of these containers to the port on the docker host.
2. Host -> Another way to access the containers externally is to associate the container to the host network. This takes out any isolation between the docker host and the docker container. Ports are now common to all containers in the host network.
3. None -> With the none network, the containers are not attached to any network and doesn’t have any access to the external network or other containers. They run in an isolated network.
 By default, docker creates only one internal bridge network. We could create our own internal network using the command “docker network create --driver bridge --subnet 182.18.0.0/16 [custom-network-name]”  Docker has a built-in DNS server that helps the containers to resolve each other using the container name. Docker uses network namespaces that creates a separate namespace for each container. It then uses virtual ethernet pairs to connect containers together.

DDD (Domain Driven Design)

DDD is a set of principles and practices for designing and developing software systems that match the complex needs of the business. It emphasises domain,  ubiquitous language,  entities (objects that have a distinct ID, mutable),  value objects (objects that are defined by their attributes, immutable),  aggregate,  repositories (It is a mechanism for abstracting data access and persistence logic from the domain model),  services (operations that do not naturally belong to entities or value objects),  bounded context (It defines explicit boundaries within which a specific term has a specific meaning),
context mapping (It is the process of defining the relationships/interactions between bounded contexts)
 

Some architectures that follow DDD are: Hexagonal Architecture, Layered Architecture, EDA (Event-Driven Architecture), CQRS (Command Query Responsibility Segregation), Micro-services Architecture, Onion Architecture, Service-Oriented Architecture (SOA), Pipeline Architecture and Screaming Architecture. 

Hexagonal Architecture: 
	Core domain -> should have models & services
	Ports -> interfaces for primary interactions
	Adapters -> Connect ports to external systems (databases, REST API, 3rd party apps,… )

DBMS

Disk Structure: Disk is divided into tracks (concentric circles) and sectors. Blocks are the intersection of tracks and sectors. Every block is of 512 bytes. Each byte has offset like 0, 1, 2, …, 511.
We can read info at any byte by block address (track, sector) and offset.
Any data is brought from disc to RAM (main memory), then it is processed and then sent back at the disc byte. 
Organising the data on the main memory is Data Structure & organising the data efficiently on the disc is DBMS.

Database table -> first level index -> high level index
We want self-managed high level indexing => B tree and B+ tree

BST (1 key and 2 children) -> M-way search tree (M-1 keys and M children) -> B tree

B trees are M-way search trees with rules
In B trees, we have a record pointer from every node whereas In B+ trees, we have a record pointer only from leaf nodes. Every key will have its copy in the leaf node also. Leaf nodes are connected like a linked list. 

Indexing: Indexing in databases is a technique used to speed up the retrieval of data by creating a data structure that improves the efficiency of search operations.

————————————————————————————————————————————————————————————————————————————————

Components of a Database server
Network Layer
Frontend of database -> [Tokeniser, Parser, Optimiser]
Execution engine consists of query executor, cache manager, utilities services
[Transaction manager, Lock manager, Recovery manager]
Concurrency manager
[Shard manager, Cluster manager, Replication manager]
Storage Engine -> [Disk storage manager, Buffer manager, Index manager]
OS Interaction Layer

Transaction :- A transaction is a collection of queries that treated as one unit of work. Transaction lifespan is from BEGIN to COMMIT, there is ROLLBACK also sometimes.
Dirty read -> when a transaction reads some value which was changed by other transaction but was not committed yet. It could be rolled back in future.
Non-repeatable read -> when two consecutive reads gives different values (change is committed by another transaction in between)
Phantom read -> when a new row adds and that changes our read value

Postgres uses snapshot method for isolation.  

Generally, OS writes to its own cache instead of writing to the database. It then batches the changes together and write to database. But if the system faces crash in between, the data is lost. Fsync command is used for OS to force it to write to database.

Consistency -> Consistency in data & Consistency in reads
Tuple v/s row -> Tuple is what user sees, the tuple is the physical instance of row in the page. The same row can have have 10 tuples: 1 active and 7 for older transactions (MVCC) and 2 dead tuples

Searching string by expression is slower than searching string by value bcoz we can’t use indexing on expression.
We can also put non-key column in the index sometimes for faster reads as per our requirements.
Index Only Scan -> when the query don’t have to go back to the table.

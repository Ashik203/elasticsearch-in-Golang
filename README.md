WebSite links:

1. https://github.com/elastic/elasticsearch
2. https://www.knowi.com/blog/what-is-elastic-search/
3. https://www.elastic.co/guide/en/elasticsearch/reference/current/elasticsearch-intro.html
    
    ## What is ElasticSearch?
    
    Elasticsearch is a distributed, open-source search and analytics engine built on Apache Lucene and developed in Java. 
    Elasticsearch allows you to store, search, and analyze huge volumes of data quickly and in near real-time and give back answers in milliseconds. 
    
    It’s able to achieve fast search responses because instead of searching the text directly, it searches an index.
    
    **At its core, you can think of Elasticsearch as a server that can process JSON requests and give you back JSON data.** 
    
    A vector database is a type of database optimized for storing and searching high-dimensional data, like vectors (arrays of numbers).
    
    ## How it works?
    
    Elasticsearch uses inverted indices, a data structure that maps words to their document locations, for an efficient search.
    
    ## Documents
    
    Documents are the basic unit of information.
    
    You can think of a document like a row in a relational database, representing a given entity — the thing you’re searching for. 
    
    A document can represent an encyclopedia article or log entries from a web server.   
    
    ## Indices/Index
    
    An index is a collection of documents that have similar characteristics.
    
    You can think of the index as being similar to a database in a relational database schema.
    
    ## Inverted index
    
    It is a data structure that stores a mapping from content, such as words or numbers, to its locations in a document or a set of documents. Basically, it is a hashmap-like data structure that directs you from a word to a document.
    
    ## Cluster
    
    An Elasticsearch cluster is a group of one or more node instances that are connected together.
    
    ## Node
    
    A node is a single server that is a part of a cluster. 
    
    3 types of nodes:
    
    1. master node: controls cluster,create/ delete index,add/remove node
    2. data node: store/execute data
    3. client node: forward cluster request
    
    ## Shards
    
    Slice of index
    
    ## Replicas
    
    a replica shard is a copy of a primary shard.
    
    Replicas provide redundant copies of your data to protect against hardware failure and increase capacity to serve read requests
    
    ## The Elastic stack(ELK, E-ElasticSearch L-Logstash K-Kibana)
    
    Elastic Stack, a set of open-source tools for data ingestion, enrichment, storage, analysis, and visualization.
    
    ## Kibana
    
    Data visualization and management tool provides real time histogram.
    
    Drawback: every visualization works for one index.
    
    ## Logstash
    
    Aggregate and process data then send it to elasticsearch.
    
    ## Beats
    
    single purpose data shipping agent that sends data from system to elasticsearch.
    
    For example, Filebeat can sit on your server, monitor log files as they come in, parses them, and import into Elasticsearch in near-real-time.
    
    ## User cases
    
    1. Application search
    2. Website search
    3. Enterprise search
    4. Logging and log analytics
    5. Security and business analytics
    
    ## Reason to use elasticsearch
    
    1. Powerful API
    2. Great search engine
    3. Open source
    4. Restful
    5. Near real-time search
    6. Free
    7. Search everything
    8. Easy to get started
    9. Analytics
    10. Distributed
    
    ## Summary
    
    Elasticsearch is at its core a search engine, whose underlying architecture and components makes it fast and scalable, sitting at the heart of an ecosystem of complementary tools that together can be used for many uses cases including search, analytics, and data processing and storage.
Summary:

This folder contains the project files, notebooks and images from my final year project.

The project aim was to ingest and analyse server log files.

For proof of concept (POC) I used R to quickly prototype and analyse the data. 

In the second phase, I opted to use python because it would be easier to take my project to production.

The final project is a pipeline with two main branches. The first branch, data is ingested using Logstash and sent to Elasticsearch and displayed on dashboards using Kibana. The second branch, the data from logstash is sent to Amazon S3 and ingested into Apache Spark hosted on Databricks Community Edtion. It is there that the batch and streaming Spark applications are developed. Using databricks in-built dashboards, results from the batch and streaming spark applications are viewed.

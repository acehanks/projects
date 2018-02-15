# Airflow

This is a simple airflow project illustrating some of the capabilities of the software.

The DAG of interest is trip_dag. It runs three python scripts on an interval of 40 minutes every hour.  Also included is a dummy task and 
obtaining the data via a bash command.

The python scripts return a welcome message to screen "Hello World", retrieve the current Bitcoin price (USD) from [coindesk](https://www.coindesk.com/) api and gets the current USD/MYR and USD/SGD forex rates from [cuurencylayer](https://currencylayer.com/) api.

Airflow can be easily extended to handle more complex workflows/data pipelines.

![picture](https://github.com/dannylwe/projects/blob/master/airflow_example/screen_airflow.png) 

To get started with airflow;

```sh
export AIRFLOW_HOME=~/airflow

pip install apache-airflow

# initialize the database
airflow initdb

# start the web server on port 8080, can be adjusted as neccessary
airflow webserver -p 8080
```

The official documentation can be found [here](https://airflow.apache.org/index.html).

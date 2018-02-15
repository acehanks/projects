from datetime import datetime, timedelta
from airflow import DAG
from airflow.operators.dummy_operator import DummyOperator
from airflow.operators.python_operator import PythonOperator
from airflow.operators.bash_operator import BashOperator

default_args = {
    'owner': 'airflow',
    'depends_on_past': False,
    'email': ['airflow@example.com'],
    'email_on_failure': False,
    'email_on_retry': False,
    'retries': 2,
    'retry_delay': timedelta(minutes=5),
    'schedule_interval': '40 0-23 * * *'
}

dag = DAG('trip_dag', catchup=False, start_date=datetime(2018, 2, 15), default_args=default_args)

date_operator= BashOperator(task_id='Print_date', bash_command='date', dag=dag)
greeting_operator= BashOperator(task_id= 'Greeting', bash_command='python3 /home/osboxes/helloWorld.py', dag=dag)
t3= BashOperator(task_id= 'currency', bash_command='python3 /home/osboxes/currency.py', dag=dag)
t4= BashOperator(task_id= 'crypto_usd', bash_command='python3 /home/osboxes/crypto.py', dag=dag)
t5= BashOperator(task_id= 'pause', bash_command='sleep 5', dag=dag)
dummy_operator = DummyOperator(task_id='dummy_task', retries=3, dag=dag)


date_operator >> greeting_operator >> t3 >> t5 >> t4 >> dummy_operator


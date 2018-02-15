from datetime import datetime
from airflow import DAG
from airflow.operators import PythonOperator, BashOperator, DummyOperator

def print_hello():
    return 'Hello world!'

dag = DAG('tut1', description='Simple tutorial DAG',
          schedule_interval='40 0-23 * * *',
          start_date=datetime(2018, 2, 15), catchup=False)

dummy_operator = DummyOperator(task_id='dummy_task', retries=3, dag=dag)

hello_operator = PythonOperator(task_id='hello_task', python_callable=print_hello, dag=dag)

t4= BashOperator(task_id= 'crypto_usd', bash_command='python3 /home/osboxes/crypto.py', dag=dag)

date_operator= BashOperator(task_id='Print_date', bash_command='date', dag=dag)

dummy_operator >> hello_operator >> t4 >> date_operator

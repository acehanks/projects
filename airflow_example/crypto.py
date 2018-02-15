import requests
import json

url= "https://api.coindesk.com/v1/bpi/currentprice.json"

def crypto():
	"""
	Bitcoin price on Coindesk, updated aproximately every 1 mintute according to the documentation
	"""
	resp= requests.get(url)
	conn= json.loads(resp.text)
	#print(conn)
	print ("\nBitcoin price is ${} on {} ".format(conn['bpi']['USD']['rate'], conn['time']['updateduk']) )
	
if __name__ == '__main__':
	crypto()


import requests

#http://apilayer.net/api/live?access_key=f9f65847ef1f170f5566767861fdf62a&currencies=USD,AUD,CAD,PLN,MXN&format=1

api = "f9f65847ef1f170f5566767861fdf62a"
uri = "http://apilayer.net/api/live?"
json_format= "&format=1"
source = "&source=USD"

def forex():
	"""
	Getting forex rates from API
	doc files @ https://currencylayer.com/documentation
	limit= 1000 calls per month

	"""

	resp= requests.get(uri + "access_key" + "=" + api + source + "&currencies=EUR,AUD,MYR,SGD,GBP,EUR")
	conn= resp.json()
	usd_myr= conn['quotes']['USDMYR']
	usd_sgd= conn['quotes']['USDSGD']
	print(conn)	
	print("\n**-----------------------------------------------------------**\n")
	print("USD/MYR rate is {}\n".format(usd_myr))
	print("**-----------------------------------------------------------**")
	print("USD/SGD rate is {}\n".format(usd_sgd))

	return conn, usd_sgd, usd_myr

if __name__=='__main__':
	forex()

	



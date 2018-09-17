
import re
password = "dAkjd123$"

def checker(string):
    if len(string) < 6: 
        print("too few chars")
    elif len(string) > 12:
        print("too many chars")
    if len(string) >=6 and len(string) <=12:
        print("length ok")

def s_checker():
    if search_num is None:
        print("no numeric")
    if search_word is None:
        print("no alpha numeric")
    if search_symb is None:
        print("no special character")
    if (search_num and search_word and search_symb) is not None:
        print("passed numeric, alpha-numeric and non numeric test")
        


search_num= re.search(r'[\d]+', password) #numeric
#print(search_num)
search_word= re.search(r'[\w]+', password) #alpha numeric
#print(search_word)
search_symb= re.search(r'[\W]+', password) #non-alpha numeric
#print(search_symb)

checker(password)
s_checker()


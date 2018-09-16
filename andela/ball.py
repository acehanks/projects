import random, time, sys

def not_valid():
	print("invalid input")
	return magic()

def start():
	return ready()

def ready():
	ans = raw_input("Do you want to continue? y/n \n")
	
	if (ans != "y" and ans != "n" and ans != "YES" and ans!= "NO"
		and ans != "yes" and ans != "no"):
	 	
		start()

	if ans.lower().startswith("y"):

		magic()

	elif ans.lower().startswith("n"):
		print("Goodbye")
		time.sleep(1)
		exit()


def magic():
	while True:
		user = str(raw_input("Enter a question?\n"))
					
		print("Loading...\n")

		time.sleep(3)

		responses= ["Choose your fate", "In the end, there can only be one", "choose wisely young grasshopper", 
					"trying is never a bad thing", "even I am stuck on this one", "wait for it.......LEGENDARY!! Hope that helps",
					"tomorrow is a new day with many opportunities", "coldplay. Listen to some coldplay", "no zero sum game. Look it up",
					"that is for another day, the aliens are COMING!", "honestly, I don't know. you are on your own on this one."]

		print(random.choice(responses))
		time.sleep(1)

		ready()

		return True

if __name__ == '__main__':
	magic()
	
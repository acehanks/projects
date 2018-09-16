import unittest


def num_list(nums):
	not_three= []

	for i in set(nums):
		if nums.count(i) < 3:
			not_three.append(i)

		pass

	print("number with count less than three in list is: ")
	return not_three

def test():
	assert num_list([5, 5, 5, 3, 4, 4, 4, 6, 6, 6, 6]) == [3]
def test_1():
	assert num_list([6, 6, 6, 9]) == [9]
"""
run test // pip install pytest
py.test three.py
"""


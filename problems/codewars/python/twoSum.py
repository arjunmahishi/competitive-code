# Problem link: https://www.codewars.com/kata/52c31f8e6605bcc646000082/train/python

# this is the shitty O(n^2) way to do it
def two_sum(numbers, target):
    for i in range(len(numbers)):
        for j in range(i, len(numbers)):
            if i!=j and numbers[i]+numbers[j] == target: return (i,j)

# two_sum([2, 2, 3], 4) 
#   -> (0,1)
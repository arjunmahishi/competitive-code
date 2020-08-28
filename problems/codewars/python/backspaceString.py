# Problem link: https://www.codewars.com/kata/5727bb0fe81185ae62000ae3/train/python

def clean_string(s):
    stack = []
    for char in s:
        if char != "#":
            stack.append(char)
        elif len(stack) > 0: 
            stack.pop()
    return ''.join(stack)

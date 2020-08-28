# Problem link: https://www.codewars.com/kata/5a2c22271f7f709eaa0005d3/train/python

def solve(s):
    if isPala(s):
        return "OK"
    
    for i in range(len(s)):
        if isPala(skipChar(s, i)):
            return "remove one"
    return "not possible"

def isPala(s):
    return s == s[::-1]

def skipChar(s, index):
    return s[:index]+s[index+1:]

# Problem link: https://www.codewars.com/kata/5ca22e6b86eed5002812061e/train/python

gFib = [0, 1]

def nearest_fibonacci(number):
    if number > gFib[-1]:
        extendFib(number)
        
    for i in range(len(gFib)):
        if gFib[i] >= number:
            break
                        
    if abs(gFib[i-1] - number) <= abs(gFib[i] - number):
        return gFib[i-1]
    
    return gFib[i]

def extendFib(n):
    while gFib[-1] < n:
        gFib.append(gFib[-1] + gFib[-2])
    return gFib
    

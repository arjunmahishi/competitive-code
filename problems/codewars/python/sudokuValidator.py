# Problem link: https://www.codewars.com/kata/53db96041f1a7d32dc0004d2/train/python

def done_or_not(board): #board[i][j]
    regionSum = {}

    for i in range(9):
        sum = 0
        for j in range(9):
            sum += board[i][j]
            regionCode = getRegionCode(i,j)
            if regionCode in regionSum:
                regionSum[regionCode] += board[i][j]
            else:
                regionSum[regionCode] = board[i][j]
        if sum != 45:
            return "Try again!"

    if not all([e == 45 for e in regionSum.values()]):
        return "Try again!"
    return "Finished!"

def getRegionCode(i, j):
    return  str(int(i/3)) + str(int(j/3))
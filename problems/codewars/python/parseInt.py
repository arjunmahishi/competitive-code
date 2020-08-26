# Problem link: https://www.codewars.com/kata/525c7c5ab6aecef16e0001a5/train/python

## Work in progress

wordMap = {
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9,
    'ten': 10,
    'eleven': 11,
    'twelve': 12,
    'thirteen': 13,
    'fourteen': 14,
    'fifteen': 15,
    'sixteen': 16,
    'seventeen': 17,
    'eighteen': 18,
    'nineteen': 19,
    'twenty': 20,
    'thirty': 30,
    'fourty': 40,
    'fifty': 50,
    'sixty': 60,
    'seventy': 70,
    'eighty': 80,
    'ninty': 90,
    'hundred': 100,
}


def parse_int(string):
    string = string.split(" ")[::-1]
    for word in string:
        if word.lower() in wordMap.keys():
            print (wordMap[word.lower()])
    return # number


parse_int("two hundred and three")

import sys
import math

n=int(input())
arr=[int(input()) for i in range(5)]

m = min(arr)
print(5 + math.ceil(n/m)-1)

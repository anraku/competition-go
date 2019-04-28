import math
a, b, t = map(int, input().split())
if a > t:
    print(0)
else:
    print (b * math.floor(t / a))

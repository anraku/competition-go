from functools import reduce
import copy

def gcd( x, y ):
    while y != 0:
        x, y = y, x%y
    return x

n = int(input())
a = list(map(int, input().split()))

result = 0
l = [0]*n
r = [0]*n
for i in range(n-1):
    l[i+1] = gcd(l[i],a[i])

for i in reversed(range(1, n)):
    r[i-1] = gcd(r[i], a[i])

ans = 0
for i in range(n):
    lr = gcd(l[i], r[i])
    ans = max(ans, lr)

print(ans)

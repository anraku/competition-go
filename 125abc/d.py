n = int(input())
a = list(map(int, input().split()))

abs_list = []
count = 0
for i in range(n):
   abs_list.append(abs(a[i]))
   if a[i] < 0:
       count += 1

if count%2 == 1:
    print(sum(abs_list) - min(abs_list)*2)
else:
    print (sum(abs_list))

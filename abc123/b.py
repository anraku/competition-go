arr=[int(input()) for i in range(5)]

div=list(map(lambda x: x%10, arr))
index=0
min_v=9
for i, v in enumerate(div):
    if v!=0 and v < min_v:
        index = i
        min_v = v

del div[index]

interval = sum(list(map(lambda x: 10-x if x!=0 else 0, div)))
print(sum(arr)+interval)

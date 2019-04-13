x,y = map(int,input().split()) # 1 2
sum = 0
for i in range(2):
    if x < y:
        sum += y
        y -= 1
    else:
        sum += x
        x -= 1

print(sum)

# li = input().split() # 文字列の配列を受け取る
# print(li)
# li = list(map(int, input().split())) # integerの配列を受け取る
# print(li)

# 複数行を取得する
# 3
# foo
# bar
# hoge
# n=int(input())
# string_list=[input() for i in range(n)]
# print(string_list)

# data = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
# print(data)
#
# flat = [flatten for inner in data for flatten in inner]
# print(flat)

# a = [2, 4, 6, 8]
# b = [3, 6, 9]

# print(list(set(a) | set(b)))  # 和集合
# print(list(set(a) & set(b)))  # 積集合
#
# print(set(a))
# print(set(b))

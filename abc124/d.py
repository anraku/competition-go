# def count_zero(s, index):
#     count = 0
#     i = index
#     while s[i] != '1':
#         count += 1
#         i += 1
#     return count

n,k = map(int,input().split())
s = input()
result = 0
continue_count = 0

def count_max1(s):
    max1 = 0
    count = 0
    for c in s:
        if c == 1:
            count += 1
        else:
            if count > max1:
                max1 = count
                count = 0
    if count > max1:
        max1 = count
        count = 0
    return max1

            
def stand1(s, index):
    s_list = list(s)
    index = 0
    for i in range(index, len(s_list)):
        if s_list[i] == '0':
            s_list[i] = '1'
        else:
            index = i
            break

    return "".join(s_list), i

def next_zeros(s, index):
    for i in range(index, len(s)):
        if s[i] == '0':
            return i
    return -1

def stand_str(i, s, now):
    global continue_count
    global result
    if i > k:
        return
    cm = count_max1(s)
    if continue_count < cm:
        print (i)
        result = i
        continue_count = cm
    zero_index = next_zeros(s, now)
    if zero_index == -1:
        return
    s, now_index = stand1(s, zero_index)
    stand_str(i+1, s, now_index)

stand_str(0, s, 0)
print(result)



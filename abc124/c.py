s = input() 
answer1 = '10'*50000
answer2 = '01'*50000 

count1 = 0
for i in range(len(s)):
    if s[i] != answer1[i]:
        count1 += 1

count2 = 0
for i in range(len(s)):
    if s[i] != answer2[i]:
        count2 += 1

print(count1) if count1 < count2 else print(count2)

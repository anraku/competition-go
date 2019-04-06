import sys
 
arr=[int(input()) for i in range(5)]
k=int(input())
 
for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if abs(arr[i] - arr[i+j+1]) > k:
            print(":(")
            sys.exit()
 
print("Yay!")

# なんかうまい計算が思いつかない
a, b, t = map(int, input().split())
count = 0
for i in range(1, t+1):
    if i % a == 0:
        count += b

print(count)

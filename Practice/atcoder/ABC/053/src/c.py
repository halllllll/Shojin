# 最初上をむいてる数はなんでもいいっぽいので6と5をこの順番で繰り返し貪欲するのがいい
# 6まで -> 1回 11まで -> 2回 17まで -> 3回 22まで -> 4回  28まで -> 5回...
x = int(input())
if x <= 6:
    print(1)
elif x <= 11:
    print(2)
else:
    print((x // 11) * 2 + (2 if (x % 11) > 6 else 1 if (x % 11) > 0 else 0))
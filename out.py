all_ = 2000000-1
count_true = 0
count_false = 0

with open("test.txt", "r") as f:
    for i in range(0,all_):
        if f.readline() == "True\n":
            count_true = count_true + 1
        else :
            count_false = count_false + 1

a = count_true/all_
b = count_false/all_

print("true",a)
print("false",b)

import random

def gen():

    rooms = [0,1,2]
    guess = random.randint(0,2)

    guess_man = random.randint(0,2)

    not_to_guess_ = random.randint(0,2)
    while 1 :
        not_to_guess_ = random.randint(0,2)
        if not_to_guess_ == guess:
            continue
        elif not_to_guess_ == guess_man:
            continue
        else:
            break

    rooms.pop(not_to_guess_)

    with open("test.txt", "a") as f:
        f.write(str(guess == guess_man)+"\n")

for i in range(0,2000000):
    gen()
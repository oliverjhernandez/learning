# students_heights = input("Heights: ").split()
#
# default = 0
# for i in range(0, len(students_heights)):
#     students_heights[i] = int(students_heights[i])
#     default = default + students_heights[i]
#
#
# print(f"Average: {default / len(students_heights)}")
#


# scores = input("Scores: ").split()
#
# highest = 0
# for s in range(0, len(scores)):
#     scores[s] = int(scores[s])
#     highest = scores[s] if scores[s] > highest else highest
#
# print(f"Highest: {highest}")
#


# even = input("Even: ").split()
#
# total = 0
# for e in range(0, len(even)):
#     even[e] = int(even[e])
#     total = total + even[e] if even[e] % 2 == 0 else total
#
# print(f"Total: {total}")
#

import random

print("Welcome to the PyPassword Generator!")
length = int(input("How long would you like your password? "))
nr_numbers = int(input("How many numbers? "))
nr_letters = int(input("How many letters? "))
nr_symbols = int(input("How many symbols? "))

if nr_numbers + nr_letters + nr_symbols != length:
    print("Error. Numbers and Letters must add up to Length")
    exit(1)

letters = [
    "a",
    "b",
    "c",
    "d",
    "e",
    "f",
    "g",
    "h",
    "i",
    "j",
    "k",
    "l",
    "m",
    "n",
    "o",
    "p",
    "q",
    "r",
    "s",
    "t",
    "u",
    "v",
    "w",
    "x",
    "y",
    "z",
    "A",
    "B",
    "C",
    "D",
    "E",
    "F",
    "G",
    "H",
    "I",
    "J",
    "K",
    "L",
    "M",
    "N",
    "O",
    "P",
    "Q",
    "R",
    "S",
    "T",
    "U",
    "V",
    "W",
    "X",
    "Y",
    "Z",
]
numbers = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]
symbols = ["!", "#", "$", "%", "&", "(", ")", "*", "+"]

password_list = []

for char in range(1, nr_letters + 1):
    password_list.append(random.choice(letters))

for char in range(1, nr_symbols + 1):
    password_list += random.choice(symbols)

for char in range(1, nr_numbers + 1):
    password_list += random.choice(numbers)

random.shuffle(password_list)

passwd = "".join(password_list)
print(f"Your password: {passwd}")

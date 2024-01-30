# grades = {"Harry": 89, "Ron": 78, "Hermione": 99, "Draco": 74, "Neville": 62}
# student_grades = {}
#
#
# def evaluate_grade(grade):
#     if grade >= 91:
#         return "Outstanding"
#     elif grade >= 81:
#         return "Exceeds expectations"
#     elif grade >= 71:
#         return "Acceptable"
#     else:
#         return "Fail"
#
#
# for student in grades:
#     student_grades[student] = evaluate_grade(grades[student])
#
# print(student_grades)

import os


def clear_screen():
    os.system("cls" if os.name == "nt" else "clear")


print("Welcome to the Secret Auction Program")
bidders = {}


continue_bid = "yes"

while continue_bid == "yes":
    name = input("Whats your name? ")
    bid = int(input("Whats your bid? $"))
    bidders[name] = bid
    continue_bid = input("Are there any other bidders? yes/no ")
    clear_screen()

bid_winner = ""
bid_amount = 0

for k in bidders:
    if bidders[k] > bid_amount:
        bid_winner = k
        bid_amount = bidders[k]


print(f"The winner is {bid_winner} {bid_amount}")

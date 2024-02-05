from art import logo, vs
from data import data
import random
import os


def clear_screen():
    os.system("cls" if os.name == "nt" else "clear")


def evaluate_result(person_a, person_b):
    if person_a["follower_count"] > person_b["follower_count"]:
        return "a"
    elif person_a["follower_count"] < person_b["follower_count"]:
        return "b"
    else:
        return "draw"


def play(person_a, person_b):
    print(
        f"Compare A: {person_a['name']}, a {person_a['description']}, from {person_a['country']}"
    )
    print(vs)
    print(
        f"Compare B: {person_b['name']}, a {person_b['description']}, from {person_b['country']}"
    )


def get_random_account():
    return random.choice(data)


def game():
    person_a = get_random_account()
    is_over = False
    score = 0
    clear_screen()
    print(logo)

    while not is_over:
        person_b = get_random_account()
        while person_a == person_b:
            person_b = get_random_account()

        play(person_a, person_b)

        response = input("Who has more followers? a/b: ")

        winner = evaluate_result(person_a, person_b)

        if response == winner:
            clear_screen()
            print(logo)
            print("That's right!")
            score += 1
            person_a = person_b
        else:
            print(f"Sorry, you're wrong. Your score was {score}")
            is_over = True


game()

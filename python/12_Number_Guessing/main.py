import random


def set_difficulty(difficulty):
    if difficulty == "easy":
        return 10
    elif difficulty == "hard":
        return 5
    else:
        return 0


def play_game(pick, attempts):
    result = 0
    for _ in range(attempts):
        guess = int(input("Make a guess: "))
        if guess == pick:
            result = 1
            break
        elif guess > pick:
            print("Too high!")
            result = 0
        else:
            print("Too low!")
            result = 0
    if result == 1:
        print("That's it! You win!")
    else:
        print("Oh no, you lost!")
    return result


def game():
    print("Welcome to the Number Guessing game!!")
    print("I'm thinking of a number between 1 and 100.")
    response = input("Choose a difficulty. Type 'easy' or 'hard'.")
    pick = random.choice(range(0, 100))
    print(f"Pick: {pick}")
    attempts = set_difficulty(response)
    if attempts == 0:
        print("Wrong option. Try again.")
        exit()
    play_game(pick, attempts)


game()

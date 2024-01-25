# Step 1

import random
import hangman_art
from hangman_words import word_list
import os


def clear_screen():
    os.system("cls" if os.name == "nt" else "clear")


display = []
lives = len(hangman_art.stages) - 1

selected_word = random.choice(word_list)

for _ in selected_word:
    display.append("_")

while "_" in display:
    print(hangman_art.logo)
    print(selected_word)
    print(display)
    print(hangman_art.stages[lives])

    guess = input("Guess a letter: ").lower()
    clear_screen()
    if guess in display:
        print(f"You already tried this letter: {guess}")

    if guess in selected_word:
        for k, v in enumerate(selected_word):
            if guess == v:
                display[k] = guess
    else:
        print(f"Nice try, {guess} is not in the word")
        lives -= 1
        if lives == 0:
            end_of_game = True
            print("You lose!")
            exit()

        end_of_game = True
    if "_" not in display:
        print("You win!")

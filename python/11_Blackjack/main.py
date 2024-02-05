############### Blackjack Project #####################

# Difficulty Normal 😎: Use all Hints below to complete the project.
# Difficulty Hard 🤔: Use only Hints 1, 2, 3 to complete the project.
# Difficulty Extra Hard 😭: Only use Hints 1 & 2 to complete the project.
# Difficulty Expert 🤯: Only use Hint 1 to complete the project.

############### Our Blackjack House Rules #####################

## The deck is unlimited in size.
## There are no jokers.
## The Jack/Queen/King all count as 10.
## The the Ace can count as 11 or 1.
## Use the following list as the deck of cards:
## cards = [11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10]
## The cards in the list have equal probability of being drawn.
## Cards are not removed from the deck as they are drawn.
## The computer is the dealer.

##################### Hints #####################

# Hint 1: Go to this website and try out the Blackjack game:
#   https://games.washingtonpost.com/games/blackjack/
# Then try out the completed Blackjack project here:
#   http://blackjack-final.appbrewery.repl.run

# Hint 2: Read this breakdown of program requirements:
#   http://listmoz.com/view/6h34DJpvJBFVRlZfJvxF
# Then try to create your own flowchart for the program.

# [] Deal both user and computer a starting hand of 2 random card values.
# [] Detect when computer or user has a blackjack. (Ace + 10 value card).
# [] If computer gets blackjack, then the user loses (even if the user also has a blackjack). If the user gets a blackjack, then they win (unless the computer also has a blackjack).
# [] Calculate the user's and computer's scores based on their card values.
# [] If an ace is drawn, count it as 11. But if the total goes over 21, count the ace as 1 instead.
# [] Reveal computer's first card to the user.
# [] Game ends immediately when user score goes over 21 or if the user or computer gets a blackjack.
# [] Ask the user if they want to get another card.
# [] Once the user is done and no longer wants to draw any more cards, let the computer play. The computer should keep drawing cards unless their score goes over 16.
# [] Compare user and computer scores and see if it's a win, loss, or draw.
# [] Print out the player's and computer's final hand and their scores at the end of the game.
# [] After the game ends, ask the user if they'd like to play again. Clear the console for a fresh start.

# Hint 3: Download and read this flow chart I've created:
#   https://drive.google.com/uc?export=download&id=1rDkiHCrhaf9eX7u7yjM1qwSuyEk-rPnt

import random
import os
from art import logo


def clear_screen():
    os.system("cls" if os.name == "nt" else "clear")


def print_logo():
    print(logo)


def deal_card():
    """Returns a random card from the deck"""
    cards = [11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10]
    return random.choice(cards)


def add_values(x, y):
    return x + y


def calculate_score(cards):
    """Evaluate if there's blackjack in a deck of cards"""
    if len(cards) == 2 and sum(cards) == 21:
        return 0
    elif 11 in cards and sum(cards) > 21:
        index = cards.index(11)
        cards[index] = 1
    return sum(cards)


def compare(user_score, computer_score):
    if user_score == computer_score:
        return "Draw 😄"
    elif computer_score == 0:
        return "Lose, opponent has blackjack 😢"
    elif user_score == 0:
        return "Win with blackjack 🤣"
    elif user_score > 21:
        return "You went over 21! You lose! 💀"
    elif computer_score > 21:
        return "Opponent went over 21, you win! 💃"
    elif user_score > computer_score:
        return "You win! 😄"
    else:
        return "You lose! 🙃"


def play_game():
    user_cards = [deal_card() for _ in range(2)]
    computer_cards = [deal_card() for _ in range(2)]
    is_game_over = False

    while not is_game_over:
        user_score = calculate_score(user_cards)
        computer_score = calculate_score(computer_cards)
        print(f"Your cards: {user_cards}, current score: {user_score}")
        print(f"Dealer first card: {computer_cards[0]}")

        if user_score == 0 or computer_score == 0 or user_score > 21:
            is_game_over = True
        else:
            user_should_deal = input("Type 'y' to draw another card: ")
            if user_should_deal == "y":
                user_cards.append(deal_card())
            else:
                is_game_over = True

    while computer_score != 0 and computer_score < 17:
        computer_cards.append(deal_card())
        computer_score = calculate_score(computer_cards)

    print(f"Your final hand is: {user_cards}, final score: {user_score}")
    print(f"Dealer final hand is: {computer_cards}, final score: {computer_score}")
    print(compare(user_score, computer_score))


while input("Do you want to play a new game of BlackJack? y/n") == "y":
    clear_screen()
    print_logo()
    play_game()

import random


player = int(input("What do you choose? "))

plays = [
    """
    _______
---'   ____)
      (_____)
      (_____)
      (____)
---.__(___)
""",
    """
     _______
---'    ____)____
           ______)
          _______)
         _______)
---.__________)
""",
    """
    _______
---'   ____)____
          ______)
       __________)
      (____)
---.__(___)
""",
]

computer = random.randrange(0, 2)


if player >= 3 or player < 0:
    print("Error. Invalid input")
    exit(1)
else:
    print(f"You chose {player}")
    print(plays[player])

    print(f"I chose {computer}")
    print(plays[computer])


def main():
    if (
        (player == 0 and computer == 1)
        or (player == 1 and computer == 2)
        or (player == 2 and computer == 0)
    ):
        print("You lose")
    elif player == computer:
        print("It's a draw")
    else:
        print("You win")


if __name__ == "__main__":
    main()

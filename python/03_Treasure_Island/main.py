print("Welcom to treasure island")
print("Your mission is to find the treasure")
crossroad = input(
    "You're at a crossroad. Where do you want to go? Type 'left' or 'right' "
)

if crossroad == "left":
    lake = input(
        "You find yourself in front of a lake. Do you want to 'swim' or 'wait'? "
    )
    if lake == "wait":
        door = input(
            "You have found three doors. 'red' 'blue' and 'yellow', which one do you want to go through? "
        )
        if door == "yellow":
            print("Congrats! you have found the treasure! ")
        else:
            print("Game over")
    else:
        print("Game over")
else:
    print("Game over")

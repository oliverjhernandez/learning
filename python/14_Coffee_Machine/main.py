from time import sleep
from data import MENU, resources


def check_quantities(request):
    result = {"water": False, "coffee": False, "milk": False}

    if resources["water"] >= MENU[request]["ingredients"].get("water", 0):
        result["water"] = True
    else:
        return False, "Not enough water"

    if resources["coffee"] >= MENU[request]["ingredients"].get("coffee", 0):
        result["coffee"] = True
    else:
        return False, "Not enough coffee"

    if resources["milk"] > MENU[request]["ingredients"].get("milk", 0):
        result["milk"] = True
    else:
        return False, "Not enough milk"

    return True, result


def process_payment(request):
    quarters = float(input("How many quarters? "))
    dimes = float(input("How many dimes? "))
    nickles = float(input("How many nickles? "))
    pennies = float(input("How many pennies? "))

    quarters_to_dollars = quarters * 0.25
    dimes_to_dollars = dimes * 0.10
    nickles_to_dollars = nickles * 0.05
    pennies_to_dollars = pennies * 0.01

    total = (
        quarters_to_dollars + dimes_to_dollars + nickles_to_dollars + pennies_to_dollars
    )

    if total >= MENU[request]["cost"]:
        return True, total - MENU[request]["cost"]

    return False, total


def logic(request):
    enough, message = check_quantities(request)

    if not enough:
        print(f"Sorry, {message}")
        return
    print("We're good. Let's go on.")

    enough, message = process_payment(request)

    if not enough:
        print(f"Sorry, {message} is not enough")
    else:
        resources["money"] = MENU[request]["cost"]
        if message > 0:
            print(f"Here's {round(message)} dollars in change.")
        print("Preparing your drink, wait a minute...")
        resources["water"] = MENU[request]["ingredients"].get("water", 0)
        resources["coffee"] = MENU[request]["ingredients"].get("coffee", 0)
        resources["milk"] = MENU[request]["ingredients"].get("milk", 0)
        sleep(5)
        print(f"Here's your {request}, enjoy!")


def main():
    options = ["espresso", "latte", "cappuccino"]
    hidden_options = ["off", "report"]

    while True:
        request = input("What would you like? espresso/latte/cappuccino : ")

        if request in hidden_options:
            if request == "off":
                print("Shutting down...")
                sleep(2)
                exit(0)
            if request == "report":
                print(f"Report {resources}")
        elif request in options:
            logic(request)
        else:
            print("Wrong option. Please, try again.")


if __name__ == "__main__":
    main()

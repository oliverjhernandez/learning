from time import sleep
from Menu import Menu, MenuItem
from Coffee_Maker import CoffeeMaker
from Money_Machine import MoneyMachine


def main():
    menu = Menu()
    money = MoneyMachine()
    coffee = CoffeeMaker()
    hidden_commands = ["off", "report"]
    while True:
        request = input("What would you like? espresso/latte/cappuccino ")
        if request in hidden_commands:
            if request == "off":
                print("Shutting down...")
                sleep(2)
                exit()
            elif request == "report":
                coffee.report()
                money.report()
        else:
            order = menu.find_drink(request)
            if coffee.is_resource_sufficient(order) and money.make_payment(order.cost):
                coffee.make_coffee(order)

if __name__ == "__main__":
    main()

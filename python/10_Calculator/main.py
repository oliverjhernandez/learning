from art import logo
import os


def clear_screen():
    os.system("cls" if os.name == "nt" else "clear")


def add(a, b):
    return a + b


def substract(a, b):
    return a - b


def multiply(a, b):
    return a * b


def divide(a, b):
    return a / b


operations = {"+": add, "-": substract, "*": multiply, "/": divide}


def calculator():
    clear_screen()
    print(logo)
    num1 = float(input("What's the first number? "))
    should_continue = True

    while should_continue:
        num2 = float(input("What's the next number? "))
        for o in operations:
            print(o)
        symbol = input("What operation do you want to perform? ")

        function = operations[symbol]
        answer = function(num1, num2)
        print(f"{num1} {symbol} {num2} = {answer}")

        stack_up = input(f"Should we continue with {answer}? yes/no ")
        if stack_up == "yes":
            num1 = answer
        else:
            should_continue = False
            calculator()


calculator()

print("Welcome to the Tip Calculator")
bill = float(input("What was the total bill? "))
people = float(input("How many people to split the bill? "))
tip = float(input("What percentage tip would you like to give? 10, 12 or 15? "))

tip = tip / 100
total = (bill * tip) + bill
per_people = total / people

print(f"Each person should pay: {round(per_people, 2)}")

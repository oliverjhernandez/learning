from random import randint
from typing import Dict
from turtle import Turtle, Screen


screen = Screen()
screen.screensize(500, 400)
screen.setup(width=500, height=400)

bet = screen.textinput("Turtle Race", "Who will win the race?")

race_is_on = True
colors = ["red", "blue", "green", "purple", "black", "pink"]
turtles = Dict[str, Turtle]
turtles = {}
initial_pos = (-230, -150)
starting_pos_y = initial_pos[1]


# Turtle init
for i in colors:
    new_turtle = Turtle("turtle")
    new_turtle.color(i)
    new_turtle.penup()
    new_turtle.pos()
    turtles[i] = new_turtle
    starting_pos_y += 50
    turtles[i].goto(initial_pos[0], starting_pos_y)

while race_is_on:
    for k, v in turtles.items():
        v.forward(randint(0, 10))
        if v.pos()[0] >= 200:
            if bet == k:
                print(f"The winning turtle is {k}. You won!. This is great!")
                race_is_on = False
            else:
                print(f"The winning turtle is {k}. You lost!")
                race_is_on = False


screen.exitonclick()

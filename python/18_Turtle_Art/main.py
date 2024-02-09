from turtle import Turtle, Screen, colormode
from random import choice, randint
import colorgram


trl = Turtle()

trl.shape("turtle")
trl.pensize(10)
trl.speed("fastest")
colormode(255)

screen = Screen()


def dashed_line():
    for _ in range(5):
        trl.forward(10)
        trl.penup()
        trl.forward(10)
        trl.pendown()


def angle(sides):
    return 360 / sides


def multi_figures():
    sides = 4
    while sides <= 10:
        for _ in range(0, sides):
            trl.forward(100)
            trl.left(angle(sides))
        sides += 1


def random_direction():
    directions = [0, 90, 180, 270]
    return choice(directions)


def random_color():
    r = randint(0, 255)
    g = randint(0, 255)
    b = randint(0, 255)
    return (r, g, b)


def random_walk():
    while True:
        trl.color(random_color())
        trl.forward(20)
        trl.setheading(random_direction())


def spirograph():
    trl.pensize(1)
    trl.speed("fastest")
    jump = 0
    while jump <= 360:
        trl.circle(100)
        trl.setheading(jump)
        trl.color(random_color())
        jump += 10


def spot():
    trl.pendown()
    trl.circle(1)
    trl.penup()


def initial_position():
    x_total = screen.window_width()
    y_total = screen.window_height()
    x = (x_total / 2) * -1
    y = (y_total / 2) * -1
    return (x, y)


def rgb_to_tuple(rgb):
    return (rgb[0], rgb[1], rgb[2])


def hirst_spots():
    trl.hideturtle()
    colors = colorgram.extract("./spots.jpg", 20)
    initial = initial_position()
    line = 1
    x = initial[0]
    y = initial[1]
    while line <= 10:
        spots = 1
        x = initial[0]
        while spots <= 10:
            trl.color(choice(colors).rgb)
            trl.penup()
            print(f"({x}, {y})")
            trl.setpos(x, y)
            spot()
            x += 100
            spots += 1
        line += 1
        y += 100


hirst_spots()
# spirograph()
# multi_figures()
# random_walk()

screen.exitonclick()

from turtle import Turtle, Screen

tim = Turtle()
screen = Screen()


def move_forward():
    tim.forward(10)


def move_backwads():
    tim.backward(10)


def rotate_left():
    tim.left(10)


def rotate_right():
    tim.right(10)


def clear():
    tim.reset()


screen.listen()
screen.onkey(move_forward, "w")
screen.onkey(rotate_left, "a")
screen.onkey(rotate_right, "d")
screen.onkey(move_backwads, "s")
screen.onkey(clear, "c")

screen.exitonclick()

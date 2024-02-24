from turtle import Screen
from Paddle import Paddle

screen = Screen()
screen.setup(width=800, height=600)
screen.bgcolor("black")
screen.title("PONG!")

player1_position = [(380, 20), (380, 0), (380, -20), (380, -40)]
player2_position = [(-380, 20), (-380, 0), (-380, -20), (-380, -40)]

r_pad = Paddle(screen, player1_position)
l_pad = Paddle(screen, player2_position)

screen.listen()
screen.onkey(r_pad.up, "Up")
screen.onkey(r_pad.down, "Down")
screen.onkey(l_pad.up, "w")
screen.onkey(l_pad.down, "s")

game_is_over = False

# while not game_is_over:


screen.exitonclick()

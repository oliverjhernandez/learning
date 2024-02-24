from turtle import Turtle
from typing import List

UP = 90
DOWN = 270


class Paddle(Turtle):
    """Creates a new Paddle"""

    def __init__(self, screen, position: List[tuple]):
        self.screen = screen
        self.position = position
        self.segments = []
        self.paddle = self.initialize_paddle()

    def initialize_paddle(self):
        self.screen.tracer(0)
        for pos in self.position:
            self.new_segment(pos)
        self.screen.update()
        self.screen.tracer(1)

    def new_segment(self, pos):
        sqr = Turtle("square")
        sqr.color("white")
        sqr.penup()
        sqr.setheading(UP)
        sqr.goto(pos[0], pos[1])
        self.segments.append(sqr)

    def up(self):
        self.screen.tracer(0)
        for sq in self.segments:
            sq.sety(sq.ycor() + 20)
        self.screen.update()
        self.screen.tracer(1)

    def down(self):
        self.screen.tracer(0)
        for sq in self.segments:
            sq.sety(sq.ycor() - 20)
        self.screen.update()
        self.screen.tracer(1)

from turtle import Turtle

INITIAL_POSITIONS = [(0, 0), (-20, 0), (-40, 0)]
MOVE_DISTANCE = 20

UP = 90
DOWN = 270
LEFT = 180
RIGHT = 0


class Snake:
    """Creates an snake object wich can move"""

    def __init__(self):
        self.segments = []
        self.snake = self.create_snake()
        self.head = self.segments[0]

    def create_snake(self):
        for pos in INITIAL_POSITIONS:
            self.new_segment(pos)

    def move(self):
        for seg in range(len(self.segments) - 1, 0, -1):
            new_x = self.segments[seg - 1].xcor()
            new_y = self.segments[seg - 1].ycor()
            self.segments[seg].goto(new_x, new_y)
        self.head.forward(MOVE_DISTANCE)

    def up(self):
        if self.head.heading() != DOWN:
            self.head.setheading(UP)

    def down(self):
        if self.head.heading() != UP:
            self.head.setheading(DOWN)

    def right(self):
        if self.head.heading() != LEFT:
            self.head.setheading(RIGHT)

    def left(self):
        if self.head.heading() != RIGHT:
            self.head.setheading(LEFT)

    def new_segment(self, pos):
        sqr = Turtle("square")
        sqr.color("white")
        sqr.penup()
        sqr.goto(pos[0], pos[1])
        self.segments.append(sqr)

    def extend(self):
        self.new_segment(self.segments[-1].position())

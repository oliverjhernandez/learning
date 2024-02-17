from turtle import Turtle


class Scoreboard(Turtle):
    """Board to display current score"""

    def __init__(self):
        super().__init__()
        self.score = 0
        self.penup()
        self.goto(0, 280)
        self.color("white")
        self.penup()
        self.update_scoreboard(self.score)

    def increaseScore(self):
        self.clear()
        self.score += 1
        self.update_scoreboard(self.score)

    def update_scoreboard(self, value):
        self.write(arg=f"Score: {value}", align="center", font=("Arial", 15, "normal"))

    def game_over(self):
        self.goto(0, 0)
        self.write(arg="GAME OVER", align="center", font=("Arial", 15, "normal"))

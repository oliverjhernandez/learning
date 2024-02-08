from typing import DefaultDict


class QuizBrain:
    def __init__(self, question_list):
        self.list = question_list
        self.question_number = 0
        self.score = 0

    def manage_response(self, response):
        return response.lower() in [
            "true",
            "1",
            "t",
            "y",
            "yes",
            "yeah",
            "yup",
            "certainly",
            "uh-huh",
        ]

    def still_has_question(self):
        return self.question_number < len(self.list)

    def next_question(self):
        next_question = self.list[self.question_number]
        self.question_number += 1
        response = input(
            f"Q.{self.question_number}: {next_question.text} (True/False)?: "
        )
        treated_response = self.manage_response(response)
        treated_answer = self.manage_response(next_question.answer)
        self.check_answer(treated_answer, treated_response)

    def check_answer(self, answer, response):
        if response == answer:
            self.score += 1
            print("You're correct!. Nice.")
            print("Get ready for the next question.")
        else:
            print("Nooooo")
        print("")
        print(f"Your score is: {self.score}/{self.question_number}")

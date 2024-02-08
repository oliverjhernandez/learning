from question_model import Question
from quiz_brain import QuizBrain
from data import question_data

question_bank = []

for item in question_data:
    text = item["text"]
    ans = item["answer"]
    question = Question(text, ans)
    question_bank.append(question)

quiz = QuizBrain(question_bank)
while quiz.still_has_question():
    reposne = quiz.next_question()

print("You've completed the quiz.")
print(f"Your score is {quiz.score} / {quiz.question_number}")

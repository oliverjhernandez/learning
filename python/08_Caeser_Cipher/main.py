import string
from art import logo


alphabet_list = list(string.ascii_letters)


def caesar_cypher(message, shift, operation):
    message_list = [char for char in message]
    encrypted_message_list = []
    coded_char_position = 0

    for i in message_list:
        if i in alphabet_list:
            position = alphabet_list.index(i)
            if operation == "encode":
                coded_char_position = position + shift
            elif operation == "decode":
                coded_char_position = position - shift
            encrypted_message_list.append(alphabet_list[coded_char_position])
        else:
            encrypted_message_list.append(i)
    print(f"Your message: {"".join(encrypted_message_list)}")


def main():
    state = "yes"
    print(logo)
    while state == "yes":
        operation = input("Type 'encode' or 'decode': ")
        message = input("Type your message: ")
        shift = int(input("Type the shift number: "))
        shift = shift % 26
        caesar_cypher(message, shift, operation)
        state = input("Do you want to do it again? ")


if __name__ == "__main__":
    main()

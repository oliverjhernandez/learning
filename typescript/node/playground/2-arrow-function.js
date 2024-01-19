
// const square = function(number) {
//     return number * number
// }

// const square = (number) => {
//     return number * number;
// }

// const square = (number) => number * number

// console.log(square(8));


const event = {
    name: "Birthday Party",
    guests: ['Andrew', 'Mike', 'Jane', 'Julie'],
    printGuestList() {
        console.log('Guest List for ' + this.name)
        this.guests.forEach((guest) => {
            console.log(guest + ' is attending ' + this.name)
        })
    }
}

event.printGuestList();

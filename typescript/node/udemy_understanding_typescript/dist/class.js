"use strict";
class Department {
    constructor(name, id) {
        this.name = name;
        this.id = id;
        this.getName = () => {
            return this.name;
        };
        this.addNewEmployee = (name) => {
            this.employees.push(name);
            return this.employees;
        };
        this.employees = [];
    }
}
class ITDepartment extends Department {
    constructor(id, admins) {
        super(id, "IT");
        this.admins = admins;
    }
}
const lpg = new ITDepartment("LPG", ["corina", "oliver"]);
lpg.addNewEmployee("Corina");
lpg.addNewEmployee("Tatiana");
const newList = lpg.addNewEmployee("Oliver");
console.log(newList);

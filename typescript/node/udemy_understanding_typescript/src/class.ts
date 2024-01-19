abstract class Department {
  protected employees: string[];

  constructor(
    private name: string,
    protected id: string,
  ) {
    this.employees = [];
  }

  static createEmployee(name: string) {
    return { name: name };
  }

  abstract describe(this: Department): void;

  getName() {
    return this.name;
  }

  addEmployee(name: string) {
    this.employees.push(name);
    return this.employees;
  }

  printEmployees() {
    console.log(this.employees);
  }
}

export class ITDepartment extends Department {
  admins: string[];
  constructor(id: string, admins: string[]) {
    super(id, "IT");
    this.admins = admins;
  }

  describe(): void {
    console.log(`Department: ${this.id}`);
  }
}

export class Accounting extends Department {
  private lastReport: string;
  private static instance: Accounting;

  get mostRecentReport() {
    if (this.lastReport) {
      return this.lastReport;
    }
    throw new Error("No report found");
  }

  set mostRecentReport(value: string) {
    if (!value) {
      throw new Error("No value input");
    }
    this.addReport(value);
  }

  private constructor(
    id: string,
    private reports: string[],
  ) {
    super(id, "AC");
    this.lastReport = reports[0];
  }

  static getInstance(id: string, reports: string[]) {
    if (this.instance) {
      return this.instance;
    }
    this.instance = new Accounting(id, reports);
    return this.instance;
  }

  describe() {
    console.log(`Department: ${this.id}`);
  }

  addEmployee(name: string) {
    if (name === "Max") {
      return [];
    }
    this.employees.push(name);
    return this.employees;
  }

  addReport(text: string) {
    this.reports.push(text);
  }

  printReports() {
    console.log(this.reports);
  }
}

// const lpg = new ITDepartment("d1", ["corina", "oliver"]);
// lpg.addEmployee("Corina");
// lpg.addEmployee("Tatiana");
// const newList = lpg.addEmployee("Oliver");
// console.log(lpg);
// console.log(newList);

// const acc = new Accounting("d3", ["Report1", "Report2", "Report3"]);
// const acc = Accounting.getInstance("d5", ["reports4"]);
// const acc2 = Accounting.getInstance("d3", ["report5"]);

// acc.mostRecentReport = "Hello";
// acc.printReports();

// acc.addEmployee("Max");
// acc.addEmployee("Test");
// acc.addEmployee("Oliver");
// acc.printEmployees();

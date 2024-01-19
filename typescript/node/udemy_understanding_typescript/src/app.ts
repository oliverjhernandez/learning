interface GreetFn {
  (text: string): string;
}

interface Named {
  readonly name: string;
}

interface Greets extends Named {
  greet: GreetFn;
}

interface Ages extends Named {
  readonly age: number;
}

interface Smiles extends Named {
  smile: boolean;
}

class Person implements Greets, Ages, Smiles {
  constructor(
    public name: string,
    public age: number,
    public smile: boolean,
  ) {}

  greet(text: string): string {
    console.log(text);
    return text;
  }
}

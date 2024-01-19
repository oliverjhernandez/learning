// Generalization: translates simply as reusing code, in this case
// by using inheritance so the Dog class gets all attributes and
// methods from the Animal class.

public abstract class Animal {
  protected int numberOfLegs;
  protected int numberOfTails;
  private String name;

  public Animal(String petName, int legs, int tails) {
    this.name = petName;
    this.numberOfLegs = legs;
    this.numberOfTails = tails;
  }

  public void walk() {}

  public void run() {}

  public void eat() {}
}

public class Dog extends Animal {
  public Dog(string name, int legs, int tails) {
    super(name, legs, tails);
  }

  public void playFetch() {}
}

public class Cat extends Animal {
  public Cat(string name, int legs, int tails) {
    super(name, legs, tails);
  }

  public void playWithYarn() {}
}

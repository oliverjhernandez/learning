// Decomposition: A relation between two objects can be
// weak or strong, in this first case, a PetStore has a
// relationship with pets but one doesnt stop existing
// when the other does.
// In the second example, the relation is strong so one
// object can stop existing if the other does.

public class PetStore {
  private ArrayList<Pet> pets;

  public PetStore() {
    pets = new ArrayList<Pet>();
  }

  public void addNewPet(Pet pet) {
    push(pets, pet);
  }

  public String[] getPets() {
    return pets;
  }
}

public class Employee {
  private Salary salary;

  public Salary getSalary() {}

  public void setSalary(Salary salary) {}
}

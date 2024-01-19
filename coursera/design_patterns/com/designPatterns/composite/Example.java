// // Overall type
// public interface IStructure {
//   public void enter();
//
//   public void exit();
//
//   public void location();
//
//   public String getName();
// }
//
// // Composite class
// public class Housing implements IStructure {
//
//   private ArrayList<IStructure> structures;
//   private String address;
//
//   public Housing(String address) {
//     this.structures = new ArrayList<IStructure>();
//     this.address = address;
//   }
//
//   public String getName() {
//     return this.address;
//   }
//
//   public int addStructure(IStructure component) {
//     this.structures.add(component);
//     return this.structures.size() - 1;
//   }
//
//   public IStructure getStructures(int componentNumber) {
//     return this.structures.get(componentNumber);
//   }
//
//   public void location() {
//     System.out.println("You are currently in " + this.getName() + ". It has ");
//     for (IStructure struct : this.structures) System.out.println(struct.getName());
//   }
//
//   public void enter() {}
//
//   public void exit() {}
// }
//
// // Leaf class
// public abstract class Room implements IStructure {
//   public String name;
//
//   public void enter() {
//     System.out.println("You have entered the " + this.name);
//   }
//
//   public void exit() {
//     System.out.println("You have exited the " + this.name);
//   }
//
//   public void location() {
//     System.out.println("You are currently in " + this.name);
//   }
//
//   public String getName() {
//     return this.name;
//   }
// }
//
// public class Main {
//
//   public static void main() {
//     Housing building = new Housing("123 Street");
//     Housing floor1 = new Housing("123 Street -- First Floot");
//     int firstFloor = building.addStructure(floor1);
//
//     Room washRoom1m = new Room("1F Men's Washroom");
//     Room washRoom1w = new Room("!f Women's Washroom");
//     Room common1 = new Room("1F common area");
//
//     int firstMens = floor1.addStructure(washRoom1m);
//     int firstWomans = floor1.addStructure(washRoom1w);
//     int firstCommon = floor1.addStructure(common1);
//
//     building.enter();
//     Housing currentFloor = building.getStructures(firstFloor);
//     currentFloor.enter();
//
//     Room currentRoom = building.getStructures(firstMens);
//     currentRoom.enter();
//
//     currentRoom = building.getStructures(common1);
//     currentRoom.enter();
//   }
// }

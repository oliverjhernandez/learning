// public class Subject {
//   private ArrayList<Observer> observers = new ArrayList<Observer>();
//
//   public void registerObserver(Observer observer) {
//     observers.add(observer);
//   }
//
//   public void unregisterObserver(Observer observer) {
//     observers.remove(observer);
//   }
//
//   public void notify() {
//     for (Observer o : observers) {
//       o.update();
//     }
//   }
// }
//
// public class Blog extends Subject {
//   private String state;
//
//   private String getState() {
//     return state;
//   }
//
//   // some other blog responsabilities
//
// }
//
// public class Auction extends Subject {
//
//   private String state;
//
//   private String getState() {
//     return state;
//   }
// }
//
// public interface IObserver {
//   public void update();
// }
//
// public class Auctioneer1 implements IObserver {
//
//   public void update() {
//     // make a bid
//
//   }
// }
//
package com.designPatterns.observerPattern;

public interface Observer {
  public void update(String status);
}

// Follower.java

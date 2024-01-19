package com.designPatterns.statePattern;

final class State {
  private State() {}

  // all potential vending machine states as singletons

  public static final State Idle = new State();
  public static final State HasOneDollar = new State();
  public static final State OutOfStock = new State();
}

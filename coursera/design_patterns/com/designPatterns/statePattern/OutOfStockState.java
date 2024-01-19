package com.designPatterns.statePattern;

public class OutOfStockState implements IState {
  public void insertDollar(VendingMachine vendingMachine) {
    System.out.println("Dollar inserted");
    vendingMachine.setState(vendingMachine.getHasOneDollarState());
  }

  public void ejectMoney(VendingMachine vendingMachine) {
    System.out.println("No money to return");
  }

  public void dispense(VendingMachine vendingMachine) {
    System.out.println("Payment required");
  }
}

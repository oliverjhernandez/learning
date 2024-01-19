package com.designPatterns.adapter;

public class CoffeeTouchscreenAdapter implements CoffeeMachineInterface {
  public void chooseFirstSelection() {
    OldCoffeeMachine old = new OldCoffeeMachine();
    old.selectA();
  }

  public void chooseSecondSelection() {
    OldCoffeeMachine old = new OldCoffeeMachine();
    old.selectB();
  }
}

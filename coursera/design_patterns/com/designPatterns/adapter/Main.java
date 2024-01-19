package com.designPatterns.adapter;

public class Main {
  public static void main(String selection) {
    CoffeeTouchscreenAdapter adapter = new CoffeeTouchscreenAdapter();

    if (selection == "A") {
      adapter.chooseFirstSelection();
    } else if (selection == "B") {
      adapter.chooseSecondSelection();
    }
  }
}

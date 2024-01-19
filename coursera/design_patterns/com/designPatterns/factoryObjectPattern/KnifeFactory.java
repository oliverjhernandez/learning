package com.designPatterns.factoryObjectPattern;

public class KnifeFactory {
  public Knife createKnife(String knifeType, int length) {
    Knife knife = null;

    // create knife object
    if (knifeType.equals("steak")) {
      knife = new SteakKnife(knifeType, length);
    } else if (knifeType.equals("chefs")) {
      knife = new ChefKnife(knifeType, length);
    }

    return knife;
  }
}

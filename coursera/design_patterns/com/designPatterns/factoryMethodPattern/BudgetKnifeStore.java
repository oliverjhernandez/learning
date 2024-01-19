package com.designPatterns.factoryMethodPattern;

public class BudgetKnifeStore extends KnifeStore {

  Knife createKnife(String knifeType) {
    Knife knife;
    // create knife object
    if (knifeType.equals("steak")) {
      knife = new SteakKnife("flat", 3);
    } else if (knifeType.equals("chefs")) {
      knife = new ChefKnife("curved", 6);
    } else {
      return new SteakKnife("saw", 4);
    }
    return knife;
  }
}

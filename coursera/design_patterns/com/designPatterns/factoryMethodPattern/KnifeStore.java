package com.designPatterns.factoryMethodPattern;

public abstract class KnifeStore {

  public Knife OrderKnife(String knifeType) {
    Knife knife;
    knife = this.createKnife(knifeType);

    // prepare knife
    knife.sharpen();
    knife.polish();
    knife.store();

    return knife;
  }

  abstract Knife createKnife(String knifeType);
}

package com.designPatterns.factoryObjectPattern;

public class KnifeStore {
  private KnifeFactory factory;

  public KnifeStore(KnifeFactory factory) {
    this.factory = factory;
  }

  public Knife OrderKnife(String knifeType, int length) {
    Knife knife = factory.createKnife(knifeType, length);

    // prepare knife
    knife.sharpen();
    knife.polish();
    knife.store();

    return knife;
  }
}

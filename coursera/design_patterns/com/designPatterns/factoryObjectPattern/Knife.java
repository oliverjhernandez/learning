package com.designPatterns.factoryObjectPattern;

public abstract class Knife {
  private String handle;
  private int length;

  public Knife(String handle, int length) {
    this.handle = handle;
    this.length = length;
  }

  public String getHandle() {
    return this.handle;
  }

  public int getLength() {
    return this.length;
  }

  public void sharpen() {}

  public void polish() {}

  public void store() {}
}

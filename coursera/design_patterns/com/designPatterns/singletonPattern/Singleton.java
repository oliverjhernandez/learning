package com.designPatterns.singletonPattern;

// Singleton: In java a private constructor wont
// allow it to be called from outside the class
// that wat we can control how the instance is
// created and when

public class Singleton {
  private static Singleton uniqueInstance = null;

  private Singleton() {}

  public static Singleton getInstance() {
    if (uniqueInstance == null) {
      uniqueInstance = new Singleton();
    }
    return uniqueInstance;
  }
}

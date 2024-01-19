package com.designPatterns.templatePattern;

public class SelfDrivingCar extends SelfDrivingVehicle {

  protected void drive() {
    System.out.println("Im driving now...");
  }

  protected void accelerate() {
    System.out.println("Look how fast i can gooo.....");
  }

  protected void useBreak() {
    System.out.println("No way, i'm not crashing");
  }

  protected void steer() {
    System.out.println("Whoooooo");
  }
}

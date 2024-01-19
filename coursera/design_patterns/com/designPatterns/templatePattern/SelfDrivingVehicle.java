package com.designPatterns.templatePattern;

public abstract class SelfDrivingVehicle {
  public void driveToDestination() {
    accelerate();
    drive();
    useBreak();
    steer();
    reachDestination();
  }

  private void reachDestination() {
    System.out.println("Arriving destination");
  }

  abstract void drive();

  abstract void accelerate();

  abstract void useBreak();

  abstract void steer();
}

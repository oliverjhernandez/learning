package com.designPatterns.composite;

public interface IComponent {
  public void setPlaybackSpeed(float speed);

  public void play();

  public String getName();
}

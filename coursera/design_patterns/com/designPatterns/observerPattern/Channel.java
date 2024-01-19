package com.designPatterns.observerPattern;

import java.util.ArrayList;

public class Channel implements Subject {

  private ArrayList<Observer> observers;
  private String channelName;
  private String status;

  public Channel(String channelName) {
    this.channelName = channelName;
    this.status = "someStatus";
    this.observers = new ArrayList<Observer>();
  }

  public String getChannelName() {
    return this.channelName;
  }

  public String getStatus() {
    return this.status;
  }

  public void setStatus(String status) {
    this.status = status;
  }

  public void notifyObservers() {
    for (Observer o : observers) {
      System.out.println("Boom!" + o + ", you're notified!");
    }
  }

  public void registerObserver(Observer observer) {
    this.observers.add(observer);
  }

  public void removeObserver(Observer observer) {
    int elementToRemove = observers.indexOf(observer);

    this.observers.remove(elementToRemove);
  }
}

package com.designPatterns.decoratorPattern;

public abstract class WebPageDecorator {

  protected WebPage page;

  public WebPageDecorator(WebPage webpage) {
    this.page = webpage;
  }

  public void display() {
    this.page.display();
  }
}

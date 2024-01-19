package com.designPatterns.decoratorPattern;

public class AuthenticatedWebPage extends WebPageDecorator {
  public AuthenticatedWebPage(WebPage webpage) {
    super(webpage);
  }

  public void authenticatedUser() {
    System.out.println("Authenticating user");
  }

  public void display() {
    super.display();
    this.authenticatedUser();
  }
}

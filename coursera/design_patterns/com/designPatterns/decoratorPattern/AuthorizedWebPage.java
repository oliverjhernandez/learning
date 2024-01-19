package com.designPatterns.decoratorPattern;

public class AuthorizedWebPage extends WebPageDecorator {
  public AuthorizedWebPage(WebPage webpage) {
    super(webpage);
  }

  public void authorizedUser() {
    System.out.println("Authorizing user");
  }

  public void display() {
    super.display();
    this.authorizedUser();
  }
}

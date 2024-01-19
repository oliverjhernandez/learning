package com.designPatterns.decoratorPattern;

public class BasicWebPage implements WebPage {
  private String html;
  private String styleSheet;
  private String scripts;

  public BasicWebPage(String html, String styleSheet, String scripts) {
    this.html = html;
    this.styleSheet = styleSheet;
    this.scripts = scripts;
  }

  public String getHtml() {
    return this.html;
  }

  public String getStylesSheet() {
    return this.styleSheet;
  }

  public String getScripts() {
    return this.scripts;
  }

  public void display() {
    /* Renders the HTML of the stylesheet and any
    embedded scripts */
    System.out.println("Base web page");
  }
}

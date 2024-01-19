package com.designPatterns.proxyPatter;

import java.util.ArrayList;

public class OrderFulfillment implements IOrder {
  private ArrayList<Warehouse> warehouses;

  /* Constructor and other attributes go here */

  public void fulfillOrder(Order order) {

    /* For each item in a customer order, check each warehouse
    to see if it is in stock. If it is then create a new Order
    for that warehouse. Else check the next warehouse. Send all the Orders
    to the warehouse(s) after you finish iterating over all the items in the original Order. */

  }
}

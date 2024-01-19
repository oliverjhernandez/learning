package com.designPatterns.proxyPatter;

import java.util.Hashtable;

public class Warehouse implements IOrder {
  private Hashtable<String, Integer> stock;
  private String address;

  /*  Constructore and other methods here */

  public void fulfillOrder(IOrder order) {
    for (Item item : order.itemList) this.stock.replace(item.sku, stock.get(item) - 1);

    /* Process the order for shipment and delivery */

  }

  public int currentInventory(Item item) {
    if (stock.containsKey(item.sku)) {
      return stock.get(item.sku).intValue();
      return 0;
    }
  }
}

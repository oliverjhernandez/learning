package com.designPatterns.composite;

import java.util.ArrayList;

public class Playlist implements IComponent {
  public String playlistName;
  public ArrayList<IComponent> playlist = new ArrayList<IComponent>();

  public Playlist(String playlistName) {
    this.playlistName = playlistName;
  }

  public void play() {
    for (IComponent song : playlist) {
      song.play();
    }
  }

  public void setPlaybackSpeed(float speed) {
    for (IComponent component : playlist) {
      component.setPlaybackSpeed(speed);
    }
  }

  public String getName() {
    return this.playlistName;
  }

  public void add(IComponent component) {
    // Add song to playlist
  }

  public void remove(IComponent component) {
    // Add song to playlist
  }
}

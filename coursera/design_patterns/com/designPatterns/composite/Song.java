package com.designPatterns.composite;

public class Song implements IComponent {
  public String songName;
  public String artist;
  public float speed = 1; // Default playback speed

  public Song(String songName, String artist) {
    this.songName = songName;
    this.artist = artist;
  }

  public void play() {
    // Play song
  }

  public void setPlaybackSpeed(float speed) {
    // Set playback speed
  }

  public String getName() {
    return this.songName;
  }

  public String getArtist() {
    return this.artist;
  }
}

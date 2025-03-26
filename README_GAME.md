# CyberBASIC Game Programming Features

This document describes the game programming features available in CyberBASIC.

## Graphics

CyberBASIC provides several functions for drawing graphics:

- `CLEARSCREEN(r, g, b)`: Clears the screen with the specified RGB color.
- `DRAWLINE(x1, y1, x2, y2, r, g, b)`: Draws a line between two points.
- `DRAWCIRCLE(x, y, radius, r, g, b)`: Draws a circle.
- `DRAWRECT(x, y, width, height, r, g, b)`: Draws a rectangle.
- `DRAWTEXT(text, x, y, size, r, g, b)`: Draws text on the screen.
- `LOADIMAGE(filename)`: Loads an image and returns its ID.
- `DRAWIMAGE(imageID, x, y)`: Draws an image at the specified coordinates.
- `LOADSPRITE(filename)`: Loads a sprite and returns its ID.
- `DRAWSPRITE(spriteID, x, y)`: Draws a sprite at the specified coordinates.
- `SETSPRITEFRAME(spriteID, frameIndex)`: Sets the current frame of a sprite.
- `SAVECANVAS(filename)`: Saves the current canvas to a file.

## Audio

CyberBASIC provides functions for playing audio:

- `LOADSOUND(filename)`: Loads a sound file and returns its ID.
- `PLAYSOUND(soundID)`: Plays a sound by ID.
- `STOPSOUND(soundID)`: Stops a sound by ID.
- `SETSOUNDVOLUME(soundID, volume)`: Sets the volume for a specific sound.
- `SETSOUNDLOOPING(soundID, looping)`: Sets whether a sound should loop.
- `ISSOUNDPLAYING(soundID)`: Checks if a sound is currently playing.
- `LOADMUSIC(filename)`: Loads a music file and returns its ID.
- `PLAYMUSIC(musicID)`: Plays music by ID.
- `STOPMUSIC()`: Stops the currently playing music.
- `SETMUSICVOLUME(musicID, volume)`: Sets the volume for a specific music track.
- `ISMUSICPLAYING()`: Checks if any music is currently playing.
- `SETMASTERVOLUME(volume)`: Sets the master volume for all audio.
- `SETSOUNDCATEGORYVOLUME(volume)`: Sets the volume for all sound effects.
- `SETMUSICCATEGORYVOLUME(volume)`: Sets the volume for all music.
- `MUTE()`: Mutes all audio.
- `UNMUTE()`: Unmutes all audio.
- `ISMUTED()`: Checks if audio is currently muted.

## Input Handling

CyberBASIC provides functions for handling keyboard and mouse input:

- `KEYPRESSED(key)`: Returns true if the key was pressed this frame.
- `KEYDOWN(key)`: Returns true if the key is currently held down.
- `KEYRELEASED(key)`: Returns true if the key was released this frame.
- `INKEY()`: Returns the last key pressed as a string.
- `MOUSEX()`: Returns the current mouse X coordinate.
- `MOUSEY()`: Returns the current mouse Y coordinate.
- `MOUSEPOSITION()`: Returns the current mouse position as (x, y).
- `MOUSEBUTTONPRESSED(button)`: Returns true if the button was pressed this frame.
- `MOUSEBUTTONDOWN(button)`: Returns true if the button is currently held.
- `MOUSEBUTTONRELEASED(button)`: Returns true if the button was released this frame.
- `MOUSEWHEEL()`: Returns the mouse wheel delta for this frame.
- `SETMOUSEPOSITION(x, y)`: Sets the mouse position.

## Game Loop

CyberBASIC provides functions for managing a game loop:

- `STARTGAMELOOP(updateFunc, renderFunc)`: Starts the game loop with provided update and render functions.
- `STOPGAMELOOP()`: Stops the game loop.
- `SETTARGETFPS(fps)`: Sets the target frames per second.
- `GETFPS()`: Returns the current frames per second.
- `GETELAPSEDTIME()`: Returns the total elapsed time since the game started.
- `GETDELTATIME()`: Returns the time elapsed since the last frame.
- `GETFRAMECOUNT()`: Returns the total number of frames processed.
- `ISGAMERUNNING()`: Returns whether the game loop is currently running.

## Screen Information

CyberBASIC provides functions for getting screen information:

- `GETSCREENWIDTH()`: Returns the width of the screen.
- `GETSCREENHEIGHT()`: Returns the height of the screen.
- `GETSCREENCENTER()`: Returns the center of the screen as (x, y).

## Utilities

CyberBASIC provides utility functions:

- `SLEEP(milliseconds)`: Pauses execution for the specified number of milliseconds.

## Example

Here's a simple example of a game in CyberBASIC:

```basic
# Game state variables
LET player_x = 100
LET player_y = 100
LET player_speed = 200
LET player_sprite = 0
LET score = 0

# Initialize game
FUNCTION init()
  # Set target frame rate
  SETTARGETFPS(60)
  
  # Load the player sprite
  player_sprite = LOADSPRITE("player.png")
  
  # Initialize player at center of screen
  player_x = GETSCREENWIDTH() / 2
  player_y = GETSCREENHEIGHT() / 2
  
  RETURN TRUE
END FUNCTION

# Update game state
FUNCTION update(delta_time)
  # Handle player input
  IF KEYDOWN("up") OR KEYDOWN("w") THEN
    player_y = player_y - player_speed * delta_time
  END IF
  
  IF KEYDOWN("down") OR KEYDOWN("s") THEN
    player_y = player_y + player_speed * delta_time
  END IF
  
  IF KEYDOWN("left") OR KEYDOWN("a") THEN
    player_x = player_x - player_speed * delta_time
  END IF
  
  IF KEYDOWN("right") OR KEYDOWN("d") THEN
    player_x = player_x + player_speed * delta_time
  END IF
  
  # Keep player within screen bounds
  IF player_x < 0 THEN
    player_x = 0
  END IF
  
  IF player_x > GETSCREENWIDTH() THEN
    player_x = GETSCREENWIDTH()
  END IF
  
  IF player_y < 0 THEN
    player_y = 0
  END IF
  
  IF player_y > GETSCREENHEIGHT() THEN
    player_y = GETSCREENHEIGHT()
  END IF
  
  # End game if escape key is pressed
  IF KEYPRESSED("escape") THEN
    STOPGAMELOOP()
  END IF
  
  RETURN
END FUNCTION

# Render game
FUNCTION render()
  # Clear screen with blue background
  CLEARSCREEN(0, 0, 0.8)
  
  # Draw player sprite
  DRAWSPRITE(player_sprite, player_x, player_y)
  
  # Draw score
  DRAWTEXT("Score: " + STR(score), 20, 20, 24, 1, 1, 1)
  
  RETURN
END FUNCTION

# Main program
FUNCTION main()
  # Initialize game
  LET success = init()
  
  IF NOT success THEN
    PRINT "Failed to initialize game!"
    RETURN 1
  END IF
  
  # Start game loop with update and render functions
  STARTGAMELOOP(update, render)
  
  # Wait for game to end
  WHILE ISGAMERUNNING()
    SLEEP(1)
  END WHILE
  
  PRINT "Game ended with score: " + STR(score)
  RETURN 0
END FUNCTION

# Run the main function
main() 
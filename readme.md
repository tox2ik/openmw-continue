
# OpenMW continue

This utility starts openMW and loads the latest save

By default it will look for the game in several paths and use the latest one.

Same for the save files, it picks the one that was created last, according to filesysem metadata.


**Usage continue.exe**

      -c string
            Name of character in OpenMW/Saves/«Name» (default "{none}")
      -f string
            Path to a save file (default "{none}")
      -n	Don't launch the game (norun)

## Assumptions


the game is installed in one of:

    C:/Program Files/OpenMW*/openmw.exe
    C:/Games/OpenMW*/openmw.exe


the saves are in:

    %Userprofile%\Documents\My Games\OpenMW\Saves


## Download

- https://nexus.whateve.com

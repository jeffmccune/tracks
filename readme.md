# Music Tracks

## Instructions

1. Add the album to record to a new playlist in music.app
2. Select the playlist in the left sidebar
3. Execute `go run ./cmd > playlist.txt`
4. Rename playlist.txt to match the album name.
5. Set the mac sound output to BlackHole 2ch.
6. Turn the volume all the way up in music.app.
7. Turn the volume all the way up in the mac system (already set to BlackHol 2-chan)
7. Open Audacity
8. Start recording and start playing at the same time.

Good enough to use one hand to press play on the keyboard and the other to
press record with the mouse.

Once recorded:

1. Save the project.
2. Import the labels (File -> Import -> Labels...)
3. Select the playlist.txt file created with the Go app.
4. Export the audio (File -> Export Audio...)
5. File Name: Example.mp3 (Numbering and Labels will be added)
6. Folder: Select it.
7. Format: MP3 Files.
8. Sample Rate: 44100 Hz
9. Bit Rate Mode: Preset
10. Quality: Standard, 170-210 kbps
11. Export Range: Multiple Files
12. Split files based on: Labels
13. Name files: Numbering before Label/Track Name

## Example

For Audacity recording from Blackhole with Music.app playing into Blackhole.

```
❯ go run ./cmd
0.000000        0.000000        Prologue:  Beauty and the Beast
146.906006      146.906006      Belle
456.039001      456.039001      Belle Reprise
520.972000      520.972000      Gaston
741.078003      741.078003      Gaston (Reprise)
864.864006      864.864006      Be Our Guest
1089.597008     1089.597008     Something There
1228.583000     1228.583000     The Mob Song
1438.889000     1438.889000     Beauty and the Beast
1605.195000     1605.195000     To The Fair
1723.100998     1723.100998     West Wing
1988.100998     1988.100998     The Beast Lets Belle Go
2130.100998     2130.100998     Battle On The Tower
2458.926987     2458.926987     Transformation
2808.299973     2808.299973     Beauty and the Beast (Duet)
```

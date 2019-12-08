# Submarine
a tool for playing with MOLEK-SYNTEZ solution files.

Submarine can view Molek-Syntez's .solution files, and import and export them with JSON.

This uses my [molek-cheatez](https://github.com/gamma-delta/molek-cheatez) library.

## Syntax

You can use `submarine help` for help, or `submarine help [command]` for help on a specific command.

```
# View your amphetamine solution.
$ submarine view amphetamine-1.solution
    Amphetamine: PRODUCTION
    9 parts
    Solved. Cycles: 479; Modules: 6; Symbols: 7
    Input Ammonia @ (-3, 1) Rot 0 #2
    ---
    Input Benzene @ (0, 1) Rot 0 #0
    ---
    Input Propene @ (-1, 0) Rot -5 #1
    ---
    Emitter #0 @ (-7, 1) Rot 0 #24 [~>-<....................]
    ---
    Emitter #1 @ (-2, -5) Rot -5 #24 [.~......................]
    ---
    Emitter #2 @ (0, -7) Rot -5 #24 [-.#.....................]
    ---
    Emitter #3 @ (7, 0) Rot -3 #24 [........................]
    ---
    Emitter #4 @ (7, -7) Rot -5 #24 [........................]
    ---
    Emitter #5 @ (3, 4) Rot -5 #24 [........................]

# Copy a file to JSON
$ submarine copy chloroform-1.solution chloroform-cheating.json

# Edit the JSON to your heart's content
$ nano/vim/whatever chloroform-cheating.json

# ...and put it back
$ submarine copy chloroform-cheating.json chloroform-2.solution
```

## A) SWITCHING USERS/USERS IN GENERAL

The IRC server will have the ability to store "games"; objects containing variable information about what characters currently exist in said game, what room they're in, etc.
This information will never be stored on the hard drive, and will only be in memory (text messages will be logged, though).

The game object will hold "owners"; each owner will have a password that they will be required to enter upon joining the game. If somebody tries to join with a user that doesn't exist, they'll be allowed to watch the game, but not talk in it. They'll instead be asked to open the respective "audience" chat where they can commentate over what's happening in a certain room.
Each owner will have several characters they can control, and they can switch between these characters via the user switcher.

## B) ATTACKING

As mentioned in A, the server will hold games in memory. These games will of course have the attacks that players create. Interacting with that data will be quite strict. In terms of attacks, only a player's attacks will ever be open for the client to read, and they'll only be able to write to the server to tell it to send an attack.
On that note, attacks will be numbered, and the client will send the command "/ATTACK <attackID>"; no extra args will be allowed in the interest of making sure the server handles the data 100% of the time. As the server goes through the attack, it will send messages to everyone in the room, and the client will hide these messages and display special effects depending on what this message was. The server will not be concerned with whether an effect happens successfully or if it takes too long, it will simply keep sending commands and advance the turn when it finishes sending them, and those commands will be cached and executed in order on the client; the exception is if the special effect is a program that needs to be launched, in which case the server will wait until the client returns that it's ready.
On that note, the server will send an authentication token - a SHA256 that's salted using the recipient's password - when it sends a microgame to launch. The program will be expected to send either a quit signal or a final score, and the score must be sent back with a decrypted copy of that auth token. So this prevents people from hijacking somebody else's microgame, but it doesn't prevent the player from doing such. todo: find out how to prevent the player from doing such. 

## C) INVENTORY

Items will literally just be files. They'll be 15 or so bytes in size, and the first two bytes of the file will correspond to any of the (max: 65535) files that the game specifies. The remaining bytes will be any of the unique (but different from ones used in battle) authentication tokens that the server holds. To use an item, the player will drag and drop one of these "item files" from a file manager program that'll be forked and modified to have a "UI-less" mode, and the entire file will be sent to the server. If the file is successfully validated, the server will remove that authentication token and apply the effects of the item in question to the player.
Hangman Game
Description
Hangman is a word guessing game in which a player has to guess a secret word letter by letter before running out of guesses. This is a digital version of the classic game, developed in Golang.

How to Play
The game selects a random word from a list of words.
The player is shown how many letters are in the word, represented by underscores.
The player guesses a letter.
If the letter is in the word, it is revealed in its corresponding position.
If the letter is not in the word, the player loses a life.
The game continues until the player either guesses the word or loses all their lives.
Installation
Install Golang on your computer if it is not already installed.
Clone the repository to your local machine.
Navigate to the directory containing the main.go file.
Run the command go run main.go to start the game.
Game Options
Difficulty: The game offers different levels of difficulty, which can be selected by the player at the start of the game. The difficulty levels are easy, medium, and hard, and they affect the number of lives the player has and the length of the secret word.
Word List: The game uses a list of words to randomly select the secret word. The list can be modified to include different words or languages.
Author
Babar Zaman

Credits
The list of words used in the game was obtained from ["JAZZ", "ZIGZAG", "ZILCH", "ZIPPER", "ZOMBIE", "ZODIAC", "FLUFF"].
The Golang programming language was used to develop this game.
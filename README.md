# tremcat

Streams random cursed cats straight into your terminal via chafa.

## Requirements

- Chafa
- Go 1.25+

## Run server

    go run .

## Get a cat

    curl -s "http://<host>:8080/?w=$(tput cols)&h=$(tput lines)"

## Credits

Images sourced from r/cursedcats for non-commercial / meme purposes.

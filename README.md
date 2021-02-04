# Limelight

Inspired by a club I used to go to back in the good old days,
this tool is a quick interface to Clubhouse.

In particular, I want to be able to use Clubhouse UI for 
high level planning but I want to be able to use a CLI
that let's me stay plugged in and focus on getting things
done like todo.sh.

## Set Up

You need to get the program either from releases or using go get.

You need to get an API key from Clubhouse.

You can then specify that with `--clubhouse-token xyz` at the command line,
or put it in your `~/.limelight.yaml` file.  See `limelight.yaml.example` in
the project root for the format.

## Running

Running simple search on stories.

`go run limelight.go clubhouse stories`

You can also supply a query in `--clubhouse-query`.  If it works here: 
`https://app.clubhouse.io/<your-org>/search` it should work with limelight.

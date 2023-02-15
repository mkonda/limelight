# Limelight

Inspired by [a club](https://en.wikipedia.org/wiki/The_Limelight#New_York_City_location)
I used to go to back in the good old days, this tool is a quick interface to Shortcut,
which was formerly known as Clubhouse.

In particular, I want to be able to use Shortcut UI for
high level planning but I want to be able to use a CLI
that let's me stay plugged in and focus on getting things
done like todo.sh.

## Set Up

You need to get the program either from releases or using go get.

Suppose you are developing with it, you might want to do this:

```sh
git clone github.com/mkonda/limelight.git
cd limelight
go build
go install
```

You need to get an API key from Shortcut in your settings [here](https://app.shortcut.com/settings/account/api-tokens).

You can then specify that with `--shortcut-token xyz` at the command line,
or put it in your `~/.limelight.yaml` file.  See `limelight.yaml.example` in
the project root for the format.

## Running

If you are running while tweaking, you'll want to run `go run limelight.go shortcut <action>`.  If you did the go install, you can also do `limelight shortcut stories`.

Running simple search on stories.

`limelight shortcut stories`

You can also supply a query in `--stories-query`.  If it works here:
`https://app.shortcut.com/<your-org>/search` it should work with limelight.

Some examples with tailored queries to `limelight shortcut stories`:

* `--stories-query "epic:\"Epic Name\""` Pull stories from an Epic
* `--stories-query "state:500000027 -is:archived"` Kanban state open and not archved
* `--stories-query "due:2021-01-01..2021-02-05 owner:mattkonda -is:archived"` Owned by Matt, not archived, due in a date range.
* `--stories-query "due:2021-01-01..2021-02-05 owner:mattkonda -is:archived label:PriorityA"` with label!

## Knowing Your States

Some of the specific states can be really handy to manipulate by hand.  But we need to know the ids for this.

`limelight shortcut states`

```txt
Workflow: Feature Requests
Workflow State:         Name: Backlog   ID: 500000016
Workflow State:         Name: To be Prioritized ID: 500000015
Workflow State:         Name: Low Impact        ID: 500000014
Workflow State:         Name: High Impact       ID: 500000013
Workflow State:         Name: Out of Scope      ID: 500000017

Workflow: Engineering
Workflow State:         Name: Unscheduled       ID: 500000008
Workflow State:         Name: Ready for Development     ID: 500000007
Workflow State:         Name: In Development    ID: 500000006
Workflow State:         Name: Ready for Review  ID: 500000010
Workflow State:         Name: Ready for Deploy  ID: 500000009
Workflow State:         Name: Completed ID: 500000011

Workflow: Simple Kanban
Workflow State:         Name: Unstarted ID: 500000027
Workflow State:         Name: Started   ID: 500000026
Workflow State:         Name: Done      ID: 500000028

Workflow: Consulting
Workflow State:         Name: Unstarted ID: 500000252
Workflow State:         Name: Started   ID: 500000251
Workflow State:         Name: Done      ID: 500000253
```

## Updating Stories

The following illustrate the two ways that we can currently update stories.

* `limelight shortcut story 1316 --story-label PriorityA` - Add a label to an existing story.
* `limelight shortcut story 885 --story-state 500000028` - Change a story stage.

For now, this is enough to get me to GTD™ from the CLI with Clubhouse as the back end.

## References

* [Shortcut API](https://developer.shortcut.com/api/rest/v3)

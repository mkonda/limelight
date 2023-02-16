# Limelight

>Limelight gives you Shortcut stories in your shell.

Inspired by [a club](https://en.wikipedia.org/wiki/The_Limelight#New_York_City_location)
I used to go to back in the good old days that was inside an
old church (see below), this tool is a quick interface to Shortcut,
which was formerly known as Clubhouse.  Basically, I want work to be as fun as dancing at Limelight so I wanted stories in my shell.  (You should go look at pictures, this place
was wild!)  OK, I can't really explain the name, I just liked it when it came up.

![LimelightPhoto](https://upload.wikimedia.org/wikipedia/commons/thumb/8/84/Avalon_NYC_2007_006.jpg/640px-Avalon_NYC_2007_006.jpg)

## Anyway, why?

I want to be able to use Shortcut UI for
high level planning but I want to be able to use a CLI
that let's me stay plugged in at my shell prompt and
focus on getting the next thing done like todo.sh.

[Here](PriorityABC.md) is a simple writeup of how I use all of
this, including how I prioritize and use the search in a structured
way to help keep track of things.

## Set Up

To run `limelight` you can either:

* Download it from the releases page
* If you have go (1.20+) installed and use it, you can `go get github.com/mkonda/limelight`

You will need to configure it using either command line flags
or by setting up a `~/.limelight.yaml` file with your API key
and query defaults.  See `limelight.yaml.example` file in the
project root for the format, or you can create a file with contents
like this and specify your own info:

```yaml
shortcut-token: abc-123
stories-query: "due:*..tomorrow owner:you -is:archived"
next-query: "owner:you"
```

### Getting a Shortcut API Key

You need to get an API key from Shortcut in your settings [here](https://app.shortcut.com/settings/account/api-tokens).

You can then specify that with `--shortcut-token xyz` at the command
line, or put it in your `~/.limelight.yaml` file.  

### Developing and Changing

Suppose you are developing with it, you might want to do this:

```sh
git clone github.com/mkonda/limelight.git
cd limelight
go build
go install
```

## Running

If you are running while tweaking, you'll want to run `go run limelight.go shortcut <action>`.  If you did the go install, you can also do `limelight shortcut <action>`.

Running simple search on stories.

`limelight shortcut stories`

This defaults to using the `story-query` in your `~/.limelight.yaml`.

You can also supply a query in `--stories-query`.  

If it works here:  `https://app.shortcut.com/<your-org>/search` it should work with limelight.

### Examples

Some examples with tailored queries to `limelight shortcut stories`:

* `--stories-query "epic:\"Epic Name\""` Pull stories from an Epic
* `--stories-query "state:500000027 -is:archived"` Kanban state open and not archved
* `--stories-query "due:2021-01-01..2021-02-05 owner:mattkonda -is:archived"` Owned by Matt, not archived, due in a date range.
* `--stories-query "due:2021-01-01..2021-02-05 owner:mattkonda -is:archived label:PriorityA"` with label!

## Knowing Your States

Some of the specific states can be really handy to manipulate by hand.
Eg. the query `--stories-query "state:50000027 -is:archived"` is looking for
a specific state that exists in our workflow.  So far, it seems like the
easiest way to do this is to use the actual state ID in the query.

To do that, we need to know the *ids*.

`limelight shortcut states`

```txt
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
...
```

## Updating Stories

The following illustrate the two ways that we can currently update stories.

* `limelight shortcut story 1316 --story-label PriorityA` - Add a label to an existing story.
* `limelight shortcut story 885 --story-state 500000028` - Change a story stage.

For now, this is enough to get me to GTD™ from the CLI with Clubhouse as the back end.

## Advanced Configuration

There are some neat things you can do that may not be immediately
obvious that were a foundational part of why I wanted `limelight`.

### Supporting Multiple Projects and Configs

You can have multiple config files and use them specifically, eg:

`go run limelight.go --config ~/.other_company_limelight.yaml shortcut stories`

I use this to be able to get at my different shortcut tasks without
having to log in and switch tenants.  You just have to set up the correct
API Key in each `.yaml` file and then reference it and everything will
"just work".

I use the default `~/.limelight.yaml` for my Jemurai work and then use
different `.yaml` files for each partner we work with that uses Shortcut.

### Default Queries

Since every Shortcut project may have different standards for how they
set up Stories, you can have a default query that correctly slices the
work to get you what you want.

For example, in the example YAML file, we specified a `stories-query`
parameter.  

```yaml
shortcut-token: abc-123
stories-query: "due:*..tomorrow owner:you -is:archived"
next-query: "owner:you"
```

This could be any valid Shortcut search term.  So you
could use labels, or dates or specific states - again depending on
what the project is actively using to make sure you have a default
query that applies well for that tenant.

You can also specify a `--stories-query` from the CLI as shown above,
but I find it helpful to have a useful default for this.

## References

* [Shortcut API](https://developer.shortcut.com/api/rest/v3)

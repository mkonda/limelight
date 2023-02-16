# Prioritizing with Limelight and Shortcut

This document captures part of how I like to start the
week and organize my work day to day.  (@mkonda)

By following a few simple conventions, I can make it
easy to get the benefit of a full featured agile project
management system like Shortcut and a simple command line
interface.

It turns out the conventions are pretty important though,
and while I think both Shortcut and `limelight` can handle
lots of different setups, I thought sharing mine might
make it easier to use.

## An Anecdote - The Why

In 1998, when I was working in Perl at Pfizer (where we were embedding the interpreter into running Apache processes to
make it fast), someone (I just remember "Bennie") that had
helped program the NYSE gave me some advise around project planning that I will never forget.

>Start by categorizing everything into Priority A, B and C.
>Then do the A's.

At the time, I was mystified.  I couldn't imagine not doing
the B's or C's.  It seemed like throwing away work.  I was
probably also overwhelmed by all of the things I had to do.

Anyway, I have found that such a simple system can really
help bubble things up and make sure we're doing what matters.

## My Setup

I like to start the week by looking at the things I really
must do.  Those get a lable of PriorityA in Shortcut, which
I associate with Red in the UI.

This includes looking at anything open, anything in the
backlog, anything that comes in through channels that I
might have otherwise missed.

Sometimes things move from PriorityA to PriorityB after
a few days.

*Note:  I found that using just "A", "B" and "C" caused
issues with searching - so a more robust term made it
easier to ensure I really got what I wanted.*

So then in Shortcut, I look at open tasks across workflows
in a planning session.  Anything I know is a priority this
week gets a PriorityA, PriorityB or PriorityC label.

In my head, I know only the A's and B's are even going to
get looked at, but honestly I am mostly focused on the A's.

Then, with `limelight` I set up my `~/.limelight.yaml` to provide appropriate searches.

1. I want the next thing I work on to always be a "PriorityA" item.  So the Shortcut search I use is:  `"owner:mattkonda label:PriorityA -is:archived -is:done"`.  That way, when I do:
`limelight shortcut next` it always grabs a PriorityA item.

1. When I look at stories, I can use the story query to show
me the PriorityA and PriorityB items as follows:  `"(label:PriorityA or label:PriorityB) due:*..tomorrow owner:mattkonda -is:archived -is:done"`.  This allows me to
see everything I set up at the beginning of the week as higher
priority.

1. Then, I can use the update capability to add PriorityA to
any previously PriorityB story I want to elevate and work on.
Eg. `limelight shortcut story 3874 --story-label PriorityA`.

1. Implied is that I have:

    1. Review stories in A/B/C each day to adjust.
    1. Review all stories weekly to put into A/B/C/None buckets.

## Workaround

Ultimately, this process also helps me to workaround some of the nuances
of Shortcut Workflows.  I often like the workflows when planning in advance,
but I don't like them so much when I just want to know what to do next!

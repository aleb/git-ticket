# Searching bugs

You can search bugs using a micro query language for both filtering and sorting. A query could look like this:

```
status:open sort:edit
```

A few tips:

- queries are case insensitive.
- you can combine as many qualifiers as you want.
- you can use double quotes for multi-word search terms. For example, `author:"René Descartes"` searches for bugs opened by René Descartes, whereas `author:René Descartes` will throw an error since full-text search is not yet supported.
- instead of a complete ID, you can use any prefix length. For example `participant=9ed1a`.


## Filtering

### Filtering by status

You can filter bugs based on their status.

| Qualifier       | Example                             |
| ---             | ---                                 |
| `status:open`   | `status:open` matches open bugs     |
| `status:closed` | `status:closed` matches closed bugs |

### Filtering by author

You can filter based on the person who opened the bug.

| Qualifier      | Example                                                                          |
| ---            | ---                                                                              |
| `author:QUERY` | `author:descartes` matches bugs opened by `René Descartes` or `Robert Descartes` |
|                | `author:"rené descartes"` matches bugs opened by `René Descartes`                |

### Filtering by participant

You can filter based on the person who participated in any activity related to the bug (Opened bug or added a comment).

| Qualifier           | Example                                                                                            |
| ---                 | ---                                                                                                |
| `participant:QUERY` | `participant:descartes` matches bugs opened or commented by `René Descartes` or `Robert Descartes` |
|                     | `participant:"rené descartes"` matches bugs opened or commented by `René Descartes`                |

### Filtering by actor

You can filter based on the person who interacted with the bug.

| Qualifier     | Example                                                                         |
| ---           | ---                                                                             |
| `actor:QUERY` | `actor:descartes` matches bugs edited by `René Descartes` or `Robert Descartes` |
|               | `actor:"rené descartes"` matches bugs edited by `René Descartes`                |
| `

**NOTE**: interaction with bugs include: opening the bug, adding comments, adding/removing labels etc...

### Filtering by label

You can filter based on the bug's label.

| Qualifier     | Example                                                                   |
| ---           | ---                                                                       |
| `label:LABEL` | `label:prod` matches bugs with the label `prod`                           |
|               | `label:"Good first issue"` matches bugs with the label `Good first issue` |

### Filtering by title

You can filter based on the bug's title.

| Qualifier     | Example                                                                        |
| ---           | ---                                                                            |
| `title:TITLE` | `title:Critical` matches bugs with a title containing `Critical`               |
|               | `title:"Typo in string"` matches bugs with a title containing `Typo in string` |


### Filtering by missing feature

You can filter bugs based on the absence of something.

| Qualifier  | Example                                |
| ---        | ---                                    |
| `no:label` | `no:label` matches bugs with no labels |

## Sorting

You can sort results by adding a `sort:` qualifier to your query. “Descending” means most recent time or largest ID first, whereas “Ascending” means oldest time or smallest ID first.

Note: to deal with differently-set clocks on distributed computers, `git-ticket` uses a logical clock internally rather than timestamps to order bug changes over time. That means that the timestamps recorded might not match the returned ordering. More on that in [the documentation](model.md#you-cant-rely-on-the-time-provided-by-other-people-their-clock-might-by-off-for-anything-other-than-just-display)

### Sort by Id

| Qualifier                  | Example                                              |
| ---                        | ---                                                  |
| `sort:id-desc`             | `sort:id-desc` will sort bugs by their descending Ids |
| `sort:id` or `sort:id-asc` | `sort:id` will sort bugs by their ascending Ids       |

### Sort by Creation time

You can sort bugs by their creation time.

| Qualifier                               | Example                                                            |
| ---                                     | ---                                                                |
| `sort:creation` or `sort:creation-desc` | `sort:creation` will sort bugs by their descending creation time    |
| `sort:creation-asc`                     | `sort:creation-asc` will sort bugs by their ascending creation time |

### Sort by Edit time

You can sort bugs by their edit time.

| Qualifier                       | Example                                                            |
| ---                             | ---                                                                |
| `sort:edit` or `sort:edit-desc` | `sort:edit` will sort bugs by their descending last edition time    |
| `sort:edit-asc`                 | `sort:edit-asc` will sort bugs by their ascending last edition time |

# Log

Fast, composable structured and unstructured logging for Go

## Usage

## Configuration

The logger can be controlled using environment variables. The following is a list of all configuration options.

| Name                | Values                                                                              | Default                                              | Description                                                                                                                                                   |
|---------------------|-------------------------------------------------------------------------------------|------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `LOG_COLOR`         | `on`, `off`                                                                         | If outputting to TTY, `on`, otherwise `off`          | Colorize terminal output based on message level                                                                                                               |
| `LOG_DEVICE`        | `stdout`, `stderr`                                                                  | `stderr`                                             | Sets the device to which log messages will be written                                                                                                         |
| `LOG_LEVEL`         | `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `_none`, `_all`, `_min`, `_max` | `_all`                                               | [Level filtering](#level-filtering)                                                                                                                           |
| `LOG_OUTPUT_FORMAT` | `json`, `text`                                                                      | `text`                                               | Configure the format of the messages. By default the log is intended for human consumption, setting this to `json` will make it easier to digest for machines |
| `LOG_TAGS_OUTPUT`   | `on`, `off`                                                                         | `on`                                                 | Controls if the text output should include tags in the message`                                                                                               |
| `LOG_TAGS`          | Example: `aTag,+anotherTag,-otherTag,_untagged`, `_all`                             | No tag filtering applied (every message is included) | [Tag filtering](#tag-filtering)                                                                                                                               |

### Filtering Output

Log output can be filtered through [level filtering](#level-filtering) and [tag filtering](#tag-filtering).
The filters are applied in sequence, that is, if a message has been excluded from the output from the level filter, it won't even be processed by tag filtering. Tag filtering will be performed only if the message passed through the level filtering and is going to be included into the output up to this point.

#### Level Filtering

Levels are incremental. Each level is _greater_ than the previous one, according to the following sequence:

- `trace`
- `debug`
- `info`
- `warn`
- `error`
- `fatal`

`LOG_LEVEL` configuration helps filtering which messages will be written.
Setting `LOG_LEVEL` to a specific level will ensure that any message with a level _lower_ than the selected one, will be omitted.

If the level is set to `info`, all messages of `info` level or above will be written.

`_none` will ensure no messages are going to be witten, `_all` and `_min` are equivalent to `trace`, while `_max` is equivalent to `fatal`

#### Tag Filtering

Tag filtering allows finer-grained control of messages included or excluded from the log output.

Setting `LOG_TAGS` to `_all` will ensure no message is excluded from the output due to tag filtering. When `LOG_TAGS` is set to `_all`.

If `LOG_TAGS` is not set to `_all`, it should contain a list of tags separated by comma. If the message has at least one of the tags listed in `LOG_TAGS`, it will be included in the output.
A tag can be prefixed with a _modifier_, which is either `+` or `-`.

A tag prefixed with `+` **must be present on the message**. If multiple tags are prefixed with `+` they must all be present on the message. That means that a message must have every tag prefixed with a `+` and **at least one unprefixed tag** (if any present in `LOG_TAGS`).

If a tag is prefixed with `-`, it must be absent from the list of tags of the message.

Every untagged message (a message with no tags) is excluded from the output unless the keyword `_untagged` is present in the tag list.

##### Examples

###### `foo,tag1,_untagged,+bar,+baz`

A message to be included in the output must satisfy **all the following conditions**:

- Have at least one of: `foo` or `tag1`
- Have **both**: `bar` and `baz`

Alternatively the message could be untagged and it will be included in the output

###### `foo,tag1,+bar,-baz`


A message to be included in the output must satisfy **all the following conditions**:

- Have at least one of: `foo` or `tag1`
- Have: `bar`
- **Not have**: `baz`

## TODO

- `LOG_FILTERS` env variable (`_all`, `_none`, `level,tags`)

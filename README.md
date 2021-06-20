# Log

Fast, composable structured and unstructured logging for Go

## Usage

## Configuration

| Name                | Type      | Values                                             | Default | Description |
|---------------------|-----------|----------------------------------------------------|---------|-------------|
| `LOG_COLOR`         | Boolean   | `on`, `off`                                        |         |             |
| `LOG_DEVICE`        | Selection | `stdout`, `stderr`                                 |         |             |
| `LOG_FORMATTERS`    | Boolean   | `on`, `off`                                        |         |             |
| `LOG_LEVEL`         | Selection | `trace`, `debug`, `info`, `warn`, `error`, `fatal` |         |             |
| `LOG_OUTPUT_FORMAT` | Selection | `json`, `text`                                     |         |             |
| `LOG_TAGS_OUTPUT`   | Boolean   | `on`, `off`                                        |         |             |
| `LOG_TAGS`          | List      | [Tag filtering](#tag-filtering)                    |         |             |

### Level Filtering

### Tag Filtering

## TODO

- `LOG_FILTERS` env variable (`_all`, `_none`, `level,tags`)

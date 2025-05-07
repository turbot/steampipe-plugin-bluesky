# Bluesky Plugin for Steampipe

Use SQL to query posts, users and followers from Bluesky.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/bluesky)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/bluesky/tables)
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-bluesky/issues)

## Quick Start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install bluesky
```

Run a query:

```sql
select
  uri,
  text,
  created_at,
  like_count
from
  bluesky_search_recent
where
  query = 'steampipe'
limit 20;
```

## Developing

Prerequisites

- [Go](https://golang.org/doc/install) 1.24 or later
- [Steampipe](https://steampipe.io/downloads) 1.0 or later

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-bluesky.git
cd steampipe-plugin-bluesky
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/bluesky.spc
```

Try it!
```
steampipe query
> .inspect bluesky
```
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Bluesky Plugin](https://github.com/turbot/steampipe-plugin-bluesky/labels/help%20wanted)
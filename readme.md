# DiscordAutoRoleGranter
## 概要
サーバにはいった段階で、自動的にロールを付与する為のプログラム。また、同時にあるロールを付与すると、前者のロールを削除する機能付き。

## 目次
<!-- TOC -->

- [DiscordAutoRoleGranter](#discordautorolegranter)
    - [概要](#概要)
    - [目次](#目次)
    - [事前準備](#事前準備)
    - [設定](#設定)

<!-- /TOC -->

## 事前準備

- Discord上にManageRoleを持ったアプリを作成。
  - BotTokenを取得

- ビルド

```sh
$ go build 
```

## 設定
環境変数を設定

```
DISCORD_TOKEN=***********
DISCORD_ROLE_ID=***************
DISCORD_TRIGGER_ROLES=*****************, *******************
DISCORD_GUILD_ID=******************
```
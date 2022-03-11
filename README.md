# rehearsal
これは指定されたプロセスの入出力を接続し，効率よくテストツールを作成するためのツールです．

## Issue作成（各リポジトリごとにうまい具合に設定してください）
| Project \ Issue | improve | bugfix |
| :-- | :-: | :-: |
| [issue and improve][@issue_and_improve] | [Make Issue][improve@issue_and_improve] | [Make Issue][bugfix@issue_and_improve] |

## ビルド方法
dockerを用いてビルドする．
| コマンド | 内容 |
| :-: | :-- |
| `docker compose up version` | コンテナ内のGolangバージョンを確認する |
| `docker compose up tidy` | `go mod tidy` 相当を実行する |
| `docker compose up build` | `cmd/rehearsal` のビルドを実行し，`bin/rehearsal`として出力する |

## テンプレートとして使用したらはじめにすること
1. **readme.mdの修正，やらないと許さない**
    ```sh
    # 置換対策のためにスペースを追加しています
    # コマンドで使用するときはとしてください
    $ sed -i -e "s%streamwest-1629 / repo_template%<username> / <reponame>%g" readme.md
    ```

2. **リポジトリの設定を変更**
    1. Projectsの作成
        [make projects](https://github.com/StreamWest-1629/repo_template/projects/new) から`issue and improve`プロジェクトを作成します．Templateは `Automated kanban` が良いかと思います（楽なので）．
    2. [Settings](https://github.com/StreamWest-1629/repo_template/projects/settings) の `Pull Requests` から `Automatically delete head branches` の項目にチェックをつけてください．

3. **dockerfileの修正**
    - ビルド・デプロイ用にdocker (compose)を用いることを前提に組まれている．
    - そのため，`build.dockerfile` や `docker-compose.yml` をうまい具合に書き直してほしい．
        （Future Feature: 新しい言語を触るときはその都度このリポジトリのどこかにプリセットとしておいておきたい）
    > ちなみに，デフォルトではGolangを修正しやすいようにわざわざ `build.dockerfile` に記述している．

## ブランチ運用ルール（`Git-flow`ベースな感じ）

> **大前提として，およそ Issue > (fixing) > Pull request > (review) > mergeの順を守ってください**

- Issueに紐づいたブランチを作るときはIssueのDevelopmentからブランチを切る
    - Issueとブランチを紐づけて管理しやすくするため
- コミット時には必ず先頭に `#[Issue番号]:` を付ける
    - Issueから追跡できるようにするため
- ブランチ名に関しては下に示すものに加え，プロダクトに応じて `release`, `staging`, `nightly` などのブランチを用意しておく．

### ブランチ命名規則とその運用
| ブランチ名 | 運用方法 | チェックアウト/マージ先 |
| :-: | :-- | :-: |
| `main` | **開発用ブランチ．**<br/>ここで直接作業を行わないでください． | （最上位格） |
| `[Issue番号]-improve/[Issue説明]` | **機能実装系のIssueを裁く際のブランチ．**<br/>IssueのDevelopmentから作ると左記のようなブランチ名になるはず． | `main` |
| `[Issue番号]-bugfix/[Issue説明]` | **バグ修正系のIssueを裁く際のブランチ．**<br/>IssueのDevelopmentから作ると左記のようなブランチ名になるはず． | `main` |
| `[Issueベース系ブランチ名]/hotfix` | `main` ブランチ以外で修正を行う必要があるときに使用するブランチ．<br/>`123-bugfix/緊急の修正/hotfix` みたいな使い方をする． | 任意のブランチ |

<!-- Issue and Improve project -->
[@issue_and_improve]:https://github.com/streamwest-1629/rehearsal/projects/1 
<!-- Make Issue in Issue and Improve project -->
[improve@issue_and_improve]:https://github.com/streamwest-1629/rehearsal/issues/new?labels=enhancement&template=improve.md&title=improve%2F%3C%E6%A9%9F%E8%83%BD%E3%81%AE%E7%B0%A1%E5%8D%98%E3%81%AA%E8%AA%AC%E6%98%8E%3E&projects=streamwest-1629/rehearsal/1
<!-- Make Issue in Issue and Improve project -->
[bugfix@issue_and_improve]:https://github.com/streamwest-1629/rehearsal/issues/new?labels=bug&template=bugfix.md&title=bugfix%2F%3C%E5%95%8F%E9%A1%8C%E3%81%AE%E7%B0%A1%E5%8D%98%E3%81%AA%E8%AA%AC%E6%98%8E%3E&projects=streamwest-1629/rehearsal/1

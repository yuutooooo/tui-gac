# tui-app-gac

## 本ツールの概要
### 対象os
- mac os
### 何をするものなのか？
- gitコマンドを簡単に実行できるようにするツール
- `git add`,`git commit`, `git push`を自動的に行うことができるようなツール
- branchごとのissue番号を保持してそれをコミットメッセージに含めることができるツール

### 作成した背景
- 長期インターンで働いている会社ではコミットメッセージに`#issue番号`という文字列を先頭につけてコミットしなければならないという規則があり、それを毎回コミットの際に確認しにgithubのissueページを確認しに行くのがめんどくさいと感じた。
- コミットメッセージに統一性がなく、一目でどのような修正を行ったのかがわかりにくいと感じることが多かったので`#issue番号 FIX_OVERVIEW commit_message`の形式でコミットメッセージが必ず作成されるように実装した。

## 実行環境の構築
1. 本リポジトリをクローン`git clone https://github.com/mimurayutoo/tui-gac.git`を実行
2. クローンしたプロジェクトをビルド`go build -o gac`(gacの部分に関しては自身がこのツールを使用する際に入力したいコマンドを入れることでカスタムできる。)
3. 管理者権限でビルドしたファイルを`/usr/local/bin/`に移動させる。`sudo mv gac(自分で設定した名前) /usr/local/bin/`(パソコンのパスワードを求められることがある。)
4. macの`~/.config/`直下に`/gac`というディレクトリを作成して、そこに`branchIssueNum.json`という名前のファイルを作成する。

以上の手順を踏むことでコマンドに対して本ツールの起動画面が確認できる。

## 操作方法
操作方法に関してはツールのターミナルに表示してあるのでそれを確認しながら入力を行ってほしい。

https://github.com/user-attachments/assets/9cf0b920-435c-4f01-bfdf-f99501ca15c6

本ツールを使用してpushした際のコミット履歴のイメージ

<img width="1216" alt="Image" src="https://github.com/user-attachments/assets/cdc69605-90b2-4b76-b4a4-7a17df60841d" />

## 注意
- 本ツールは個人の趣味で作成したものなので、業務などでリモートのリポジトリなどにpushする際には十分に気をつけて行ってほしい。
- このツールを使用して何かトラブルが発生した際には作成者は責任も負えないのでご了承ください。
- エラーハンドリングなどの処理が自身でも甘いと感じている。コミットを取り消すのか、ステージングを取り消すのかなどの実装が甘いのでそこは目を瞑ってほしい。

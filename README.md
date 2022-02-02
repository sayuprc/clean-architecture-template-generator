# クリーンアーキテクチャひな形生成ツール

入力値を基にEntityやUseCase等のひな型を生成します。

## 使い方

### 設定ファイル生成

生成されるファイル

- .cagt.json

```shell
$ cagt init
```

### Entity生成

生成されるファイル

- Entity
- RepositoryInterface*
- Repository*
- InMemoryRepository*

* コマンド実行中にリポジトリを生成するを選択した場合

```shell
$ cagt entity -d User -e UserId
```

### UseCase生成

生成されるファイル

- UseCaseInterface
- UseCaseInteractor
- MockUseCaseInteractor
- InputDataStruct
- OutputDataStruct

#### コマンド例

```shell
$ cagt usecase -d User -u Create
```
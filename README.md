# nirvana-cms
user and menu management api for nirvana-cms.

## usage
1. firstly, your need [golang](https://golang.org/) environment.
    ```shell
    # clone nirvana-cms to your go src directory
    git clone https://github.com/narrowizard/nirvana-cms.git
    go get
    go build
    ```
1. config db info in `config/db.cfg`, we use gorm and mysql in this situation.
    ```shell
    # run nirvana-cms, it will do database definition automatically.
    ./nirvana-cms
    # import init data to db.
    mysql –u root –p your-db-name < sql/initdata.sql
    ```
1. Then config and run [nirvana-cms-auth](https://github.com/narrowizard/nirvana-cms-auth.git) and [react-cms](https://github.com/narrowizard/react-cms.git)
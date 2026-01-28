# RSS Aggregator learning project

### Install
This system runs on Go and POSTGRES
You will need both installed to run the program
Install Gator by CLI with 'go install https://github.com/devinbgreen/gator

### Configure
Configure your system in  ~/.gatorconfig.json 
It should contain:
```{"db_url":"postgres://postgres:postgres2@localhost:5432/gator?sslmode=disable","current_user_name":"yourName"}```

## Use
Commands are: login, register, users, agg, addfeed, feeds, follow, following, unfollow
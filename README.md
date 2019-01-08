# mathehappen-content-crawler

Program that pulls data from [mathehappen.de](https://mathehappen.de), transforms it, and pushes it to the [content server](https://github.com/schul-cloud/schulcloud-content).

# Compile and run
```
go get -u github.com/schul-cloud/mathehappen-content-crawler        # download, compile, and install the program

export MH_USER=MY_USERNAME                                          # must export mathehappen username
expert MH_PASSWORD=MY_PASSWORD                                      # and password
mathehappen-content-crawler                                         # then run the program
```

Options that are available through environment variables are
 - `BASIC_AUTH_USER` is the username used to authenticate at the target API (content server)
 - `BASIC_AUTH_PASSWORD` is the password matching the username
 - `TARGET_URL` is the content servers URL
 - `MH_USER` is the username used to authenticate at the mathehappen API
 - `MH_PASSWORD` is the password matching the mathehappen username

`MH_USER` and `MH_PASSWORD` have no default values and must be set externally. All other options should have correct default values so the do not need to be set explicitly.